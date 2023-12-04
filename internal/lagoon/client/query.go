package client

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// ProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newVersionedRequest("_lgraphql/projectByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// Me queries the Lagoon API for me, and
// unmarshals the response into project.
func (c *Client) Me(
	ctx context.Context, user *schema.User) error {

	req, err := c.newRequest("_lgraphql/me.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.User `json:"me"`
	}{
		Response: user,
	})
}

// EnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) EnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/environmentByName.graphql",
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// BackupsForEnvironmentByName queries the Lagoon API for an environment by its name and
// parent projectID, and unmarshals the response into environment.
func (c *Client) BackupsForEnvironmentByName(ctx context.Context, name string,
	projectID uint, environment *schema.Environment) error {

	req, err := c.newRequest("_lgraphql/backupsForEnvironmentByName.graphql",
		map[string]interface{}{
			"name":    name,
			"project": projectID,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Environment `json:"environmentByName"`
	}{
		Response: environment,
	})
}

// LagoonAPIVersion queries the Lagoon API for its version, and
// unmarshals the response.
func (c *Client) LagoonAPIVersion(
	ctx context.Context, lagoonAPIVersion *schema.LagoonVersion) error {

	req, err := c.newRequest("_lgraphql/lagoonVersion.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &lagoonAPIVersion)
}

// LagoonSchema queries the Lagoon API for its schema, and
// unmarshals the response.
func (c *Client) LagoonSchema(
	ctx context.Context, lagoonSchema *schema.LagoonSchema) error {

	req, err := c.newRequest("_lgraphql/lagoonSchema.graphql",
		nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.LagoonSchema `json:"__schema"`
	}{
		Response: lagoonSchema,
	})
}

// MinimalProjectByName queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) MinimalProjectByName(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newVersionedRequest("_lgraphql/minimalProjectByName.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ProjectByNameMetadata queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) ProjectByNameMetadata(
	ctx context.Context, name string, project *schema.ProjectMetadata) error {

	req, err := c.newVersionedRequest("_lgraphql/projectByNameMetadata.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.ProjectMetadata `json:"projectByName"`
	}{
		Response: project,
	})
}

// DeployTargetConfigsByProjectID queries the Lagoon API for a projects deploytarget configs by its id, and
// unmarshals the response into deploytargetconfigs.
func (c *Client) DeployTargetConfigsByProjectID(
	ctx context.Context, project int, deploytargetconfigs *[]schema.DeployTargetConfig) error {

	req, err := c.newVersionedRequest("_lgraphql/deployTargetConfigsByProjectId.graphql",
		map[string]interface{}{
			"project": project,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.DeployTargetConfig `json:"deployTargetConfigsByProjectId"`
	}{
		Response: deploytargetconfigs,
	})
}

// SSHEndpointsByProject queries the Lagoon API for a project by its name, and
// unmarshals the response into project.
func (c *Client) SSHEndpointsByProject(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newVersionedRequest("_lgraphql/sshEndpointsByProject.graphql",
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"projectByName"`
	}{
		Response: project,
	})
}

// ListDeployTargets queries the Lagoon API for a deploytargets and unmarshals the response into deploytargets.
func (c *Client) ListDeployTargets(
	ctx context.Context, deploytargets *[]schema.DeployTarget) error {

	req, err := c.newRequest("_lgraphql/listDeployTargets.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.DeployTarget `json:"listDeployTargets"`
	}{
		Response: deploytargets,
	})
}

// GetEnvVariablesByProjectEnvironmentName queries the Lagoon API for a envvars by project environment and unmarshals the response.
func (c *Client) GetEnvVariablesByProjectEnvironmentName(
	ctx context.Context, in *schema.EnvVariableByProjectEnvironmentNameInput, envkeyvalue *[]schema.EnvKeyValue) error {

	req, err := c.newRequest("_lgraphql/variables/getEnvVariablesByProjectEnvironmentName.graphql",
		map[string]interface{}{
			"input": in,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.EnvKeyValue `json:"getEnvVariablesByProjectEnvironmentName"`
	}{
		Response: envkeyvalue,
	})
}
