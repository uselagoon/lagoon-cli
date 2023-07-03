package client

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// AddNotificationSlack defines a Slack notification.
func (c *Client) AddNotificationSlack(ctx context.Context,
	in *schema.AddNotificationSlackInput, out *schema.NotificationSlack) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationSlack.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationSlack `json:"addNotificationSlack"`
	}{
		Response: out,
	})
}

// AddNotificationRocketChat defines a RocketChat notification.
func (c *Client) AddNotificationRocketChat(ctx context.Context,
	in *schema.AddNotificationRocketChatInput,
	out *schema.NotificationRocketChat) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationRocketChat.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationRocketChat `json:"addNotificationRocketChat"`
	}{
		Response: out,
	})
}

// AddNotificationEmail defines an Email notification.
func (c *Client) AddNotificationEmail(ctx context.Context,
	in *schema.AddNotificationEmailInput,
	out *schema.NotificationEmail) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationEmail.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationEmail `json:"addNotificationEmail"`
	}{
		Response: out,
	})
}

// AddNotificationMicrosoftTeams defines a MicrosoftTeams notification.
func (c *Client) AddNotificationMicrosoftTeams(ctx context.Context,
	in *schema.AddNotificationMicrosoftTeamsInput,
	out *schema.NotificationMicrosoftTeams) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationMicrosoftTeams.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationMicrosoftTeams `json:"addNotificationMicrosoftTeams"`
	}{
		Response: out,
	})
}

// AddNotificationWebhook defines a Webhook notification.
func (c *Client) AddNotificationWebhook(ctx context.Context,
	in *schema.AddNotificationWebhookInput,
	out *schema.NotificationWebhook) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationWebhook.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationWebhook `json:"addNotificationWebhook"`
	}{
		Response: out,
	})
}

// UpdateNotificationEmail updates an email notification.
func (c *Client) UpdateNotificationEmail(
	ctx context.Context, input *schema.UpdateNotificationEmailInput, out *schema.NotificationEmail) error {

	req, err := c.newVersionedRequest("_lgraphql/notifications/updateNotificationEmail.graphql",
		map[string]interface{}{
			"name":  input.Name,
			"patch": input.Patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationEmail `json:"updateNotificationEmail"`
	}{
		Response: out,
	})
}

// UpdateNotificationSlack updates a slack notification.
func (c *Client) UpdateNotificationSlack(
	ctx context.Context, input *schema.UpdateNotificationSlackInput, out *schema.NotificationSlack) error {

	req, err := c.newVersionedRequest("_lgraphql/notifications/updateNotificationSlack.graphql",
		map[string]interface{}{
			"name":  input.Name,
			"patch": input.Patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationSlack `json:"updateNotificationSlack"`
	}{
		Response: out,
	})
}

// UpdateNotificationRocketChat updates a rocket chat notification.
func (c *Client) UpdateNotificationRocketChat(
	ctx context.Context, input *schema.UpdateNotificationRocketChatInput, out *schema.NotificationRocketChat) error {

	req, err := c.newVersionedRequest("_lgraphql/notifications/updateNotificationRocketChat.graphql",
		map[string]interface{}{
			"name":  input.Name,
			"patch": input.Patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationRocketChat `json:"updateNotificationRocketChat"`
	}{
		Response: out,
	})
}

// UpdateNotificationMicrosoftTeams updates a microsoft teams notification.
func (c *Client) UpdateNotificationMicrosoftTeams(
	ctx context.Context, input *schema.UpdateNotificationMicrosoftTeamsInput, out *schema.NotificationMicrosoftTeams) error {

	req, err := c.newVersionedRequest("_lgraphql/notifications/updateNotificationMicrosoftTeams.graphql",
		map[string]interface{}{
			"name":  input.Name,
			"patch": input.Patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationMicrosoftTeams `json:"updateNotificationMicrosoftTeams"`
	}{
		Response: out,
	})
}

// UpdateNotificationWebhook updates a webhook notification.
func (c *Client) UpdateNotificationWebhook(
	ctx context.Context, input *schema.UpdateNotificationWebhookInput, out *schema.NotificationWebhook) error {

	req, err := c.newVersionedRequest("_lgraphql/notifications/updateNotificationWebhook.graphql",
		map[string]interface{}{
			"name":  input.Name,
			"patch": input.Patch,
		})
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *schema.NotificationWebhook `json:"updateNotificationWebhook"`
	}{
		Response: out,
	})
}

// DeleteNotificationSlack deletes a Slack notification.
func (c *Client) DeleteNotificationSlack(ctx context.Context,
	name string,
	out *schema.DeleteNotification) error {
	req, err := c.newRequest("_lgraphql/notifications/deleteNotificationSlack.graphql", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeleteNotificationRocketChat deletes a RocketChat notification.
func (c *Client) DeleteNotificationRocketChat(ctx context.Context,
	name string,
	out *schema.DeleteNotification) error {
	req, err := c.newRequest("_lgraphql/notifications/deleteNotificationRocketChat.graphql", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeleteNotificationEmail deletes an Email notification.
func (c *Client) DeleteNotificationEmail(ctx context.Context,
	name string,
	out *schema.DeleteNotification) error {
	req, err := c.newRequest("_lgraphql/notifications/deleteNotificationEmail.graphql", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeleteNotificationMicrosoftTeams deletes a MicrosoftTeams notification.
func (c *Client) DeleteNotificationMicrosoftTeams(ctx context.Context,
	name string,
	out *schema.DeleteNotification) error {
	req, err := c.newRequest("_lgraphql/notifications/deleteNotificationMicrosoftTeams.graphql", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// DeleteNotificationWebhook deletes a Webhook notification.
func (c *Client) DeleteNotificationWebhook(ctx context.Context,
	name string,
	out *schema.DeleteNotification) error {
	req, err := c.newRequest("_lgraphql/notifications/deleteNotificationWebhook.graphql", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &out)
}

// AddNotificationToProject adds a Notification to a Project.
func (c *Client) AddNotificationToProject(ctx context.Context,
	in *schema.AddNotificationToProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/notifications/addNotificationToProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"addNotificationToProject"`
	}{
		Response: out,
	})
}

// RemoveNotificationFromProject removes a Notification from a Project.
func (c *Client) RemoveNotificationFromProject(ctx context.Context,
	in *schema.RemoveNotificationFromProjectInput, out *schema.Project) error {
	req, err := c.newRequest("_lgraphql/notifications/removeNotificationFromProject.graphql", in)
	if err != nil {
		return err
	}
	return c.client.Run(ctx, req, &struct {
		Response *schema.Project `json:"removeNotificationFromProject"`
	}{
		Response: out,
	})
}

// GetProjectNotificationEmail queries the Lagoon API for notifications of the requested type
func (c *Client) GetProjectNotificationEmail(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/projectNotificationEmail.graphql",
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

// GetAllNotificationEmail queries the Lagoon API for notifications of the requested type
func (c *Client) GetAllNotificationEmail(
	ctx context.Context, projects *[]schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/listAllNotificationEmail.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Project `json:"allProjects"`
	}{
		Response: projects,
	})
}

// GetProjectNotificationWebhook queries the Lagoon API for notifications of the requested type
func (c *Client) GetProjectNotificationWebhook(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/projectNotificationWebhook.graphql",
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

// GetAllNotificationWebhook queries the Lagoon API for notifications of the requested type
func (c *Client) GetAllNotificationWebhook(
	ctx context.Context, projects *[]schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/listAllNotificationWebhook.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Project `json:"allProjects"`
	}{
		Response: projects,
	})
}

// GetProjectNotificationSlack queries the Lagoon API for notifications of the requested type
func (c *Client) GetProjectNotificationSlack(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/projectNotificationSlack.graphql",
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

// GetAllNotificationSlack queries the Lagoon API for notifications of the requested type
func (c *Client) GetAllNotificationSlack(
	ctx context.Context, projects *[]schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/listAllNotificationSlack.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Project `json:"allProjects"`
	}{
		Response: projects,
	})
}

// GetProjectNotificationRocketChat queries the Lagoon API for notifications of the requested type
func (c *Client) GetProjectNotificationRocketChat(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/projectNotificationRocketChat.graphql",
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

// GetAllNotificationRocketChat queries the Lagoon API for notifications of the requested type
func (c *Client) GetAllNotificationRocketChat(
	ctx context.Context, projects *[]schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/listAllNotificationRocketChat.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Project `json:"allProjects"`
	}{
		Response: projects,
	})
}

// GetProjectNotificationMicrosoftTeams queries the Lagoon API for notifications of the requested type
func (c *Client) GetProjectNotificationMicrosoftTeams(
	ctx context.Context, name string, project *schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/projectNotificationMicrosoftTeams.graphql",
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

// GetAllNotificationMicrosoftTeams queries the Lagoon API for notifications of the requested type
func (c *Client) GetAllNotificationMicrosoftTeams(
	ctx context.Context, projects *[]schema.Project) error {

	req, err := c.newRequest("_lgraphql/notifications/listAllNotificationMicrosoftTeams.graphql", nil)
	if err != nil {
		return err
	}

	return c.client.Run(ctx, req, &struct {
		Response *[]schema.Project `json:"allProjects"`
	}{
		Response: projects,
	})
}
