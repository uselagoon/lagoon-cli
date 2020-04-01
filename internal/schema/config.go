package schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicode"

	"github.com/amazeeio/lagoon-cli/pkg/api"
	"sigs.k8s.io/yaml"
)

// Config represents a collection of Lagoon platform configuration.
// Fields for comprising structs are dictated by the add* Lagoon APIs.
type Config struct {
	// API objects
	Projects      []ProjectConfig        `json:"projects,omitempty"`
	Groups        []GroupConfig          `json:"groups,omitempty"`
	BillingGroups []AddBillingGroupInput `json:"billingGroups,omitempty"`
	Users         []User                 `json:"users,omitempty"`
	Notifications *NotificationsConfig   `json:"notifications,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler interface to control how lagoon
// config files are unmarshaled.
func (c *Config) UnmarshalJSON(data []byte) error {
	// alias Config to create an unmarshal target and avoid infinite recursion
	// during json.Unmarshal()
	type UnmarshalConfig Config
	var uc UnmarshalConfig
	if err := json.Unmarshal(data, &uc); err != nil {
		return err
	}
	// post-process the unmarshaled object to Lagoon API requirements
	sshKeyType := map[api.SSHKeyType]api.SSHKeyType{
		"ssh-rsa":     api.SSHRsa,
		"ssh-ed25519": api.SSHEd25519,
	}
	for _, user := range uc.Users {
		for j, sshKey := range user.SSHKeys {
			if val, ok := sshKeyType[sshKey.KeyType]; ok {
				user.SSHKeys[j].KeyType = val
			}
		}
	}
	environmentDeployType := map[api.DeployType]api.DeployType{
		"branch":      api.Branch,
		"pullrequest": api.PullRequest,
		"promote":     api.Promote,
	}
	environmentType := map[api.EnvType]api.EnvType{
		"production":  api.ProductionEnv,
		"development": api.DevelopmentEnv,
	}
	envVarScope := map[api.EnvVariableScope]api.EnvVariableScope{
		"build":                       api.BuildVar,
		"runtime":                     api.RuntimeVar,
		"global":                      api.GlobalVar,
		"internal_container_registry": api.InternalContainerRegistryVar,
		"container_registry":          api.ContainerRegistryVar,
	}
	for _, project := range uc.Projects {
		for j, ev := range project.EnvVariables {
			if val, ok := envVarScope[ev.Scope]; ok {
				project.EnvVariables[j].Scope = val
			}
		}
		for j, env := range project.Environments {
			if val, ok := environmentDeployType[env.DeployType]; ok {
				project.Environments[j].DeployType = val
			}
			if val, ok := environmentType[env.EnvironmentType]; ok {
				project.Environments[j].EnvironmentType = val
			}
			for k, ev := range env.EnvVariables {
				if val, ok := envVarScope[ev.Scope]; ok {
					project.Environments[j].EnvVariables[k].Scope = val
				}
			}
		}
	}
	*c = Config(uc)
	return nil
}

// ProjectsToConfig translates an array of Projects to Lagoon Config data.
// It assumes a list of unique Projects.
func ProjectsToConfig(
	projects []Project, exclude map[string]bool) ([]byte, error) {
	config := Config{Notifications: &NotificationsConfig{}}

	// avoid adding objects multiple times
	users := map[string]bool{}
	groups := map[string]bool{}
	slackNotifications := map[string]bool{}
	rocketChatNotifications := map[string]bool{}
	emailNotifications := map[string]bool{}
	microsoftTeamsNotifications := map[string]bool{}

	for _, project := range projects {
		projectConfig :=
			ProjectConfig{Project: project, Notifications: &ProjectNotifications{}}
		for _, group := range project.Groups.Groups {
			// project group users are appended to the project directly because this
			// group is automatically created in Lagoon.
			if fmt.Sprintf("project-%s", project.Name) == group.Name {
				for _, member := range group.Members {
					// skip default users, these are created by Lagoon automatically
					if fmt.Sprintf("default-user@%s", project.Name) == member.User.Email {
						continue // next member
					}
					// add the user directly to the project
					projectConfig.Users = append(projectConfig.Users,
						UserRoleConfig{
							Email: member.User.Email,
							Role:  member.Role,
						})
					// skip creating the user if already done
					if users[member.User.Email] {
						continue // next member
					}
					users[member.User.Email] = true
					config.Users = append(config.Users, member.User)
				}
				continue // next group
			}

			projectConfig.Groups =
				append(projectConfig.Groups, group.Name)
			// skip creating the group if already done
			if groups[group.Name] {
				continue // next group
			}
			groups[group.Name] = true
			newGroup := GroupConfig{
				AddGroupInput: group.AddGroupInput,
			}
			for _, member := range group.Members {
				// add the user to the group
				newGroup.Users = append(newGroup.Users,
					UserRoleConfig{
						Email: member.User.Email,
						Role:  member.Role,
					})
				// skip creating the user if already done
				if users[member.User.Email] {
					continue // next member
				}
				users[member.User.Email] = true
				config.Users = append(config.Users, member.User)
			}
			config.Groups = append(config.Groups, newGroup)
		}
		// add billing groups
		for _, billingGroup := range project.Groups.BillingGroups {
			projectConfig.BillingGroups =
				append(projectConfig.BillingGroups, billingGroup.Name)
			// skip creating the group if already done
			if groups[billingGroup.Name] {
				continue // next group
			}
			groups[billingGroup.Name] = true
			config.BillingGroups =
				append(config.BillingGroups, billingGroup.AddBillingGroupInput)
		}
		// add notifications
		for _, n := range project.Notifications.Slack {
			projectConfig.Notifications.Slack =
				append(projectConfig.Notifications.Slack, n.Name)
			// skip creating the notification if already done
			if slackNotifications[n.Name] {
				continue // next notification
			}
			slackNotifications[n.Name] = true
			config.Notifications.Slack = append(config.Notifications.Slack, n)
		}
		for _, n := range project.Notifications.RocketChat {
			projectConfig.Notifications.RocketChat =
				append(projectConfig.Notifications.RocketChat, n.Name)
			// skip creating the notification if already done
			if rocketChatNotifications[n.Name] {
				continue // next notification
			}
			rocketChatNotifications[n.Name] = true
			config.Notifications.RocketChat =
				append(config.Notifications.RocketChat, n)
		}
		for _, n := range project.Notifications.Email {
			projectConfig.Notifications.Email =
				append(projectConfig.Notifications.Email, n.Name)
			// skip creating the notification if already done
			if emailNotifications[n.Name] {
				continue // next notification
			}
			emailNotifications[n.Name] = true
			config.Notifications.Email =
				append(config.Notifications.Email, n)
		}
		for _, n := range project.Notifications.MicrosoftTeams {
			projectConfig.Notifications.MicrosoftTeams =
				append(projectConfig.Notifications.MicrosoftTeams, n.Name)
			// skip creating the notification if already done
			if microsoftTeamsNotifications[n.Name] {
				continue // next notification
			}
			microsoftTeamsNotifications[n.Name] = true
			config.Notifications.MicrosoftTeams =
				append(config.Notifications.MicrosoftTeams, n)
		}
		minimiseProjectConfig(&projectConfig, exclude)
		config.Projects = append(config.Projects, projectConfig)
	}

	minimiseConfig(&config, exclude)

	// logic copied from yaml.Marshal() to be able to clean the JSON to ensure
	// valid UTF-8 before converting to YAML because otherwise the YAML library
	// throws errors.
	j, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal to JSON: %v", err)
	}
	// strip any non-printable runes
	j = bytes.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, j)
	// convert to YAML
	return yaml.JSONToYAML(j)
}

// minimiseProjectConfig zeroes default or empty values in a Project to
// minimise the config file.
func minimiseProjectConfig(p *ProjectConfig, exclude map[string]bool) {
	// clear bits we don't want to serialise at all
	p.Project.Notifications = nil
	// omit IDs from config
	for i := range p.Environments {
		p.Environments[i].ID = 0
	}
	p.ID = 0
	// clear exclusions
	if exclude["project-users"] {
		p.Users = nil
	}
	if exclude["project-private-keys"] {
		p.PrivateKey = ""
	}
	// this could be part of exclude, but haven't seen a use for it yet
	p.OpenshiftID = nil
	// clear empty notifications
	if p.Notifications != nil &&
		p.Notifications.Slack == nil &&
		p.Notifications.RocketChat == nil &&
		p.Notifications.Email == nil &&
		p.Notifications.MicrosoftTeams == nil {
		p.Notifications = nil
	}

	// don't set options if they're already set to default values
	defaults := projectDefaults()
	if p.ActiveSystemsDeploy == defaults.ActiveSystemsDeploy {
		p.ActiveSystemsDeploy = ""
	}
	if p.ActiveSystemsPromote == defaults.ActiveSystemsPromote {
		p.ActiveSystemsPromote = ""
	}
	if p.ActiveSystemsRemove == defaults.ActiveSystemsRemove {
		p.ActiveSystemsRemove = ""
	}
	if p.ActiveSystemsTask == defaults.ActiveSystemsTask {
		p.ActiveSystemsTask = ""
	}
	if p.PullRequests == defaults.PullRequests {
		p.PullRequests = ""
	}
	if p.Branches == defaults.Branches {
		p.Branches = ""
	}
	if p.DevelopmentEnvironmentsLimit == defaults.DevelopmentEnvironmentsLimit {
		p.DevelopmentEnvironmentsLimit = 0
	}
}

// minimiseConfig clears any configured fields from the config to minimise the
// marshalled YAML.
func minimiseConfig(c *Config, exclude map[string]bool) {
	// clear empty notifications
	if c.Notifications != nil &&
		len(c.Notifications.Slack) == 0 &&
		len(c.Notifications.RocketChat) == 0 {
		c.Notifications = nil
	}
	// handle exclusions
	if exclude["users"] {
		c.Users = nil
	}
	if exclude["groups"] {
		c.Groups = nil
	}
	if exclude["notifications"] {
		c.Notifications = nil
	}
}

// projectDefaults returns default Project values.
func projectDefaults() *ProjectConfig {
	// see https://github.com/amazeeio/lagoon/blob/
	// 817def93b3e15f5d96aa44e2b7bd33c15f18bd43
	// services/api/src/resources/project/resolvers.js#L233
	return &ProjectConfig{
		Project: Project{
			AddProjectInput: AddProjectInput{
				ActiveSystemsDeploy:          "lagoon_openshiftBuildDeploy",
				ActiveSystemsPromote:         "lagoon_openshiftBuildDeploy",
				ActiveSystemsRemove:          "lagoon_openshiftRemove",
				ActiveSystemsTask:            "lagoon_openshiftJob",
				PullRequests:                 "true",
				Branches:                     "true",
				DevelopmentEnvironmentsLimit: 5,
			},
		},
	}
}

// UnmarshalConfigYAML takes config data in YAML format and unmarshals it into
// a Config struct.
func UnmarshalConfigYAML(data []byte, config *Config) error {
	return yaml.Unmarshal(data, config)
}
