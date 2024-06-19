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
