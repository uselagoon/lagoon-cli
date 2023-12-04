package client

import (
	"context"
	"fmt"
	"regexp"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/schema"
)

var duplicate = regexp.MustCompile("^graphql: Duplicate entry ")

// wrapErr wraps a response error with a lagoon.ErrExist type if the response
// is due to an object already existing
func wrapErr(err error) error {
	if err != nil && duplicate.MatchString(err.Error()) {
		return fmt.Errorf("couldn't create object: %w: %v", lagoon.ErrExist, err)
	}
	return err
}

// AddDeployTarget adds a deploytarget (kubernetes/openshift).
func (c *Client) AddDeployTarget(ctx context.Context, in *schema.AddDeployTargetInput, out *schema.AddDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/addDeployTarget.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.AddDeployTargetResponse `json:"addDeployTarget"`
	}{
		Response: out,
	})
}

// UpdateDeployTarget updates a deploytarget (kubernetes/openshift).
func (c *Client) UpdateDeployTarget(ctx context.Context, in *schema.UpdateDeployTargetInput, out *schema.UpdateDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/updateDeployTarget.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateDeployTargetResponse `json:"updateDeployTarget"`
	}{
		Response: out,
	})
}

// DeleteDeployTarget deletes a deploytarget (kubernetes/openshift).
func (c *Client) DeleteDeployTarget(ctx context.Context, in *schema.DeleteDeployTargetInput, out *schema.DeleteDeployTargetResponse) error {
	req, err := c.newRequest("_lgraphql/deleteDeployTarget.graphql", in)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &out)
}

// AddGroup adds a group.
func (c *Client) AddGroup(
	ctx context.Context, in *schema.AddGroupInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addGroup"`
	}{
		Response: out,
	})
}

// AddUser adds a user.
func (c *Client) AddUser(
	ctx context.Context, in *schema.AddUserInput, out *schema.User) error {
	req, err := c.newRequest("_lgraphql/addUser.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"addUser"`
	}{
		Response: out,
	})
}

// AddSSHKey adds an SSH key to a user.
func (c *Client) AddSSHKey(
	ctx context.Context, in *schema.AddSSHKeyInput, out *schema.SSHKey) error {
	req, err := c.newRequest("_lgraphql/addSshKey.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.SSHKey `json:"addSshKey"`
	}{
		Response: out,
	})
}

// AddUserToGroup adds a user to a group.
func (c *Client) AddUserToGroup(
	ctx context.Context, in *schema.UserGroupRoleInput, out *schema.Group) error {
	req, err := c.newRequest("_lgraphql/addUserToGroup.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Group `json:"addUserToGroup"`
	}{
		Response: out,
	})
}

// AddProject adds a project.
func (c *Client) AddProject(
	ctx context.Context, in *schema.AddProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addProject.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addProject"`
	}{
		Response: out,
	}))
}

// AddEnvVariable adds an EnvVariable to an Environment or Project.
func (c *Client) AddEnvVariable(ctx context.Context,
	in *schema.EnvVariableInput, out *schema.EnvKeyValue) error {
	req, err := c.newRequest("_lgraphql/addEnvVariable.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.EnvKeyValue `json:"addEnvVariable"`
	}{
		Response: out,
	})
}

// AddOrUpdateEnvironment adds or updates a Project Environment.
func (c *Client) AddOrUpdateEnvironment(ctx context.Context,
	in *schema.AddEnvironmentInput, out *schema.Environment) error {
	req, err := c.newRequest("_lgraphql/addOrUpdateEnvironment.graphql", in)
	if err != nil {
		return err
	}
	return wrapErr(c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"addOrUpdateEnvironment"`
	}{
		Response: out,
	}))
}

// AddGroupsToProject adds Groups to a Project.
func (c *Client) AddGroupsToProject(ctx context.Context,
	in *schema.ProjectGroupsInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/addGroupsToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addGroupsToProject"`
	}{
		Response: out,
	})
}

// DeployEnvironmentLatest deploys a latest environment.
func (c *Client) DeployEnvironmentLatest(ctx context.Context,
	in *schema.DeployEnvironmentLatestInput, out *schema.DeployEnvironmentLatest) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentLatest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPullrequest deploys a pullreguest.
func (c *Client) DeployEnvironmentPullrequest(ctx context.Context,
	in *schema.DeployEnvironmentPullrequestInput, out *schema.DeployEnvironmentPullrequest) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentPullrequest.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentPromote promotes one environment into a new environment.
func (c *Client) DeployEnvironmentPromote(ctx context.Context,
	in *schema.DeployEnvironmentPromoteInput, out *schema.DeployEnvironmentPromote) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentPromote.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeployEnvironmentBranch deploys a branch.
func (c *Client) DeployEnvironmentBranch(ctx context.Context,
	in *schema.DeployEnvironmentBranchInput, out *schema.DeployEnvironmentBranch) error {
	req, err := c.newRequest("_lgraphql/deployEnvironmentBranch.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// AddRestore adds a restore.
func (c *Client) AddRestore(
	ctx context.Context, backupID string, out *schema.Restore) error {
	req, err := c.newRequest("_lgraphql/addRestore.graphql",
		map[string]interface{}{
			"backupid": backupID,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Restore `json:"addRestore"`
	}{
		Response: out,
	})
}

// UpdateDeployTargetConfiguration adds a deploytarget configuration to a project.
func (c *Client) UpdateDeployTargetConfiguration(ctx context.Context,
	in *schema.UpdateDeployTargetConfigInput, out *schema.DeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/updateDeployTargetConfig.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.DeployTargetConfig `json:"updateDeployTargetConfig"`
	}{
		Response: out,
	})
}

// DeleteDeployTargetConfig deletes a deploytarget config from a project.
func (c *Client) DeleteDeployTargetConfiguration(ctx context.Context,
	id int, project int, out *schema.DeleteDeployTargetConfig) error {
	req, err := c.newRequest("_lgraphql/deleteDeployTargetConfig.graphql",
		map[string]interface{}{
			"id":      id,
			"project": project,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// AddOrUpdateEnvVariableByName adds or updates an environment variable in the api
func (c *Client) AddOrUpdateEnvVariableByName(ctx context.Context, in *schema.EnvVariableByNameInput, out *schema.UpdateEnvVarResponse) error {
	req, err := c.newRequest("_lgraphql/variables/addOrUpdateEnvVariableByName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.UpdateEnvVarResponse `json:"addOrUpdateEnvVariableByName"`
	}{
		Response: out,
	})
}

// DeleteEnvVariableByName deletes an environment variable from the api
func (c *Client) DeleteEnvVariableByName(ctx context.Context, in *schema.DeleteEnvVariableByNameInput, out *schema.DeleteEnvVarResponse) error {
	req, err := c.newRequest("_lgraphql/variables/deleteEnvVariableByName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &out)
}
