package project

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/uselagoon/lagoon-cli/internal/util"
	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
	"strconv"
)

func RunCreateWizard(lc *client.Client) (*util.CreateConfig, error) {
	config := &util.CreateConfig{}
	var organizationConfirm bool
	initForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the project name").
				Value(&config.Input.Name).
				Validate(func(name string) error {
					valid, errMsg := util.IsValidProjectName(name)
					if !valid {
						return errors.New(errMsg)
					}
					if valid {
						proj, err := lagoon.GetMinimalProjectByName(context.TODO(), name, lc)
						if err != nil {
							return err
						}
						if proj.Name != "" {
							return errors.New(fmt.Sprintf("project: %s already exists", name))
						}
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Do you want to create this project in an Organization?").
				Value(&organizationConfirm),
		),
	).WithTheme(huh.ThemeCharm())

	formErr := initForm.Run()
	if formErr != nil {
		return nil, formErr
	}
	if organizationConfirm {
		raw := `query allOrgsWithProjects {
			allOrganizations {
				id
				name
				description
				friendlyName
				quotaProject
				quotaGroup
				quotaNotification
				quotaEnvironment
				quotaRoute
				projects {
					id
					name
				}
			}
		}`
		resp, err := lc.ProcessRaw(context.TODO(), raw, map[string]interface{}{})
		if err != nil {
			return nil, err
		}

		allOrgs := resp.(map[string]interface{})["allOrganizations"]
		o, err := json.Marshal(allOrgs)
		if err != nil {
			return nil, err
		}

		var organizations []schema.Organization
		err = json.Unmarshal(o, &organizations)
		if err != nil {
			return nil, err
		}

		orgValidationErr := false
		form2 := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select an organization").
					OptionsFunc(func() []huh.Option[string] {
						options := make([]huh.Option[string], len(organizations))
						for i, org := range organizations {
							orgProjectCount := len(org.Projects)
							if orgProjectCount >= org.QuotaProject && org.QuotaProject >= 0 {
								quotaFullLabel := lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Render(fmt.Sprintf("%s | Project Limit: %v/%v | ⚠️  Project Quota full, cannot assigned project to this organization.", org.Name, orgProjectCount, org.QuotaProject))
								options[i] = huh.NewOption(quotaFullLabel, org.Name)
							} else {
								options[i] = huh.NewOption(fmt.Sprintf("%s | Project Limit: %v/%v", org.Name, orgProjectCount, util.QuotaCheck(org.QuotaProject)), org.Name)
							}
						}

						return options
					}, &config.OrganizationDetails.Name).
					Value(&config.OrganizationDetails.Name).
					Validate(func(selectedOrg string) error {
						for _, org := range organizations {
							if org.Name == selectedOrg {
								orgProjectCount := len(org.Projects)
								if orgProjectCount >= org.QuotaProject && org.QuotaProject >= 0 {
									orgValidationErr = true
									return nil
									//return fmt.Errorf("Organization %s has reached its project quota (%d/%d)", org.Name, orgProjectCount, org.QuotaProject)
								}
							}
						}
						return nil
					}),
			),
		).WithTheme(huh.ThemeCharm())

		err = form2.Run()
		if err != nil {
			return nil, err
		}

		if orgValidationErr {
			var orgRetryResp string
			orgRetryRespForm := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Organization quota exceeded. What would you like to do?").
						Options(
							huh.NewOption("Go back to previous form", "back"),
							huh.NewOption("Cancel wizard", "cancel"),
						).
						Value(&orgRetryResp),
				),
			).WithTheme(huh.ThemeCharm())

			err := orgRetryRespForm.Run()
			if err != nil {
				return nil, err
			}

			switch orgRetryResp {
			case "back":
				return RunCreateWizard(lc)
			case "cancel":
				return nil, huh.ErrUserAborted
			}
		}
	}

	deploytargets, err := lagoon.ListDeployTargets(context.TODO(), lc)
	var additionalFields bool
	options := make([]huh.Option[uint], len(*deploytargets))
	for i, target := range *deploytargets {
		options[i] = huh.NewOption(target.Name, target.ID)
	}
	form3 := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the Git URL for your project").
				Value(&config.Input.GitURL).
				Validate(func(str string) error {
					if !util.IsValidGitURL(config.Input.GitURL) {
						return errors.New("Invalid Git URL")
					}
					return nil
				}),
			huh.NewInput().
				Title("Enter the branch name for the production environment").
				Value(&config.Input.ProductionEnvironment),
			huh.NewSelect[uint]().
				Title("Select a deploy target").
				Options(options...).
				Value(&config.Input.Openshift),
			huh.NewConfirm().
				Title("Do you want to define any other fields?").
				Value(&additionalFields),
		),
	).WithTheme(huh.ThemeCharm())

	err = form3.Run()
	if err != nil {
		return nil, err
	}

	if additionalFields {
		var fields []string
		additionalFieldsOptions := []huh.Option[string]{
			huh.NewOption("branches: Which branches should be deployed", "Branches"),
			huh.NewOption("pullrequests: Which Pull Requests should be deployed", "PullRequests"),
		}
		if organizationConfirm {
			additionalFieldsOptions = append(additionalFieldsOptions, huh.NewOption("owner (Only select if adding to an Organization)", "AddOrgOwner"))
		}
		form4 := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Select which fields you want to define").
					//Options(
					//	//huh.NewOption("auto-idle: Auto idle setting of the project.", "AutoIdle"),
					//	huh.NewOption("branches: Which branches should be deployed", "Branches"),
					//	//huh.NewOption("build-image: Build Image for the project", "BuildImage"),
					//	//huh.NewOption("deploytarget-project-pattern: Pattern of Deploytarget(Kubernetes) Project/Namespace that should be generated", "OpenshiftProjectPattern"),
					//	//huh.NewOption("development-environments-limit: How many environments can be deployed at one time", "DevelopmentEnvironmentsLimit"),
					//	huh.NewOption("owner (Only select if adding to an Organization)", "AddOrgOwner"),
					//	//huh.NewOption("private-key: Private key to use for the project", "PrivateKey"),
					//	huh.NewOption("pullrequests: Which Pull Requests should be deployed", "PullRequests"),
					//	//huh.NewOption("router-pattern: Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'", "RouterPattern"),
					//	huh.NewOption("standby-production-environment: Which environment(the name) should be marked as the standby production environment", "StandbyProductionEnvironment"),
					//	//huh.NewOption("storage-calc: Should storage for this environment be calculated.", "StorageCalc"),
					//	huh.NewOption("subfolder: Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository", "Subfolder"),
					//).
					Options(additionalFieldsOptions...).
					Value(&fields),
			),
		).WithTheme(huh.ThemeCharm())

		err = form4.Run()
		if err != nil {
			return nil, err
		}

		if len(fields) == 0 {
			return config, nil
		}

		var inputs []huh.Field
		var devEnvLimit string
		for _, field := range fields {
			switch field {
			case "StandbyProductionEnvironment":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.StandbyProductionEnvironment))
			case "Branches":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.Branches))
			case "PullRequests":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.PullRequests))
			case "OpenshiftProjectPattern":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.OpenshiftProjectPattern))
			case "Subfolder":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.Subfolder))
			case "PrivateKey":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.PrivateKey))
			case "BuildImage":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.BuildImage))
			case "RouterPattern":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.RouterPattern))

			case "DevelopmentEnvironmentsLimit":
				inputs = append(inputs, huh.NewInput().
					Title(fmt.Sprintf("Enter value for '%s'", field)).
					Value(&devEnvLimit).
					Validate(func(s string) error {
						if s == "" {
							return nil
						}
						_, err := strconv.ParseUint(s, 10, 64)
						return err
					}),
				)

			case "StorageCalc":
				config.StorageCalcProvided = true
				inputs = append(inputs, huh.NewConfirm().
					Title(fmt.Sprintf("Enable '%s'?", field)).
					Value(&config.StorageCalc))
			case "AutoIdle":
				config.AutoIdleProvided = true
				inputs = append(inputs, huh.NewConfirm().
					Title(fmt.Sprintf("Enable '%s'?", field)).
					Value(&config.AutoIdle))
			case "AddOrgOwner":
				config.Input.AddOrgOwner = new(bool)
				inputs = append(inputs, huh.NewConfirm().
					Title(fmt.Sprintf("Enable '%s'?", field)).
					Value(config.Input.AddOrgOwner))
			}
		}

		form5 := huh.NewForm(
			huh.NewGroup(inputs...),
		).WithTheme(huh.ThemeCharm())

		err = form5.Run()
		if err != nil {
			return nil, err
		}

		if devEnvLimit != "" {
			developmentEnvLimit, err := strconv.Atoi(devEnvLimit)
			if err != nil {
				fmt.Println("Error:", err)
				return nil, err
			}
			config.DevEnvLimit = uint(developmentEnvLimit)
		}

	}
	return config, nil
}
