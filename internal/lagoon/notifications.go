// Package lagoon implements high-level functions for interacting with the
// Lagoon API.
package lagoon

import (
	"context"

	"github.com/uselagoon/lagoon-cli/internal/schema"
)

// Notification interface contains methods for adding notifications in Lagoon.
type Notification interface {
	AddNotificationWebhook(ctx context.Context, input *schema.AddNotificationWebhookInput, result *schema.NotificationWebhook) error
	AddNotificationEmail(ctx context.Context, input *schema.AddNotificationEmailInput, result *schema.NotificationEmail) error
	AddNotificationRocketChat(ctx context.Context, input *schema.AddNotificationRocketChatInput, result *schema.NotificationRocketChat) error
	AddNotificationMicrosoftTeams(ctx context.Context, input *schema.AddNotificationMicrosoftTeamsInput, result *schema.NotificationMicrosoftTeams) error
	AddNotificationSlack(ctx context.Context, input *schema.AddNotificationSlackInput, result *schema.NotificationSlack) error
	AddNotificationToProject(context.Context, *schema.AddNotificationToProjectInput, *schema.Project) error
}

// AddNotificationWebhook adds a notification.
func AddNotificationWebhook(ctx context.Context, input *schema.AddNotificationWebhookInput, n Notification) (*schema.NotificationWebhook, error) {
	result := schema.NotificationWebhook{}
	return &result, n.AddNotificationWebhook(ctx, input, &result)
}

// AddNotificationEmail adds a notification.
func AddNotificationEmail(ctx context.Context, input *schema.AddNotificationEmailInput, n Notification) (*schema.NotificationEmail, error) {
	result := schema.NotificationEmail{}
	return &result, n.AddNotificationEmail(ctx, input, &result)
}

// AddNotificationRocketChat adds a notification.
func AddNotificationRocketChat(ctx context.Context, input *schema.AddNotificationRocketChatInput, n Notification) (*schema.NotificationRocketChat, error) {
	result := schema.NotificationRocketChat{}
	return &result, n.AddNotificationRocketChat(ctx, input, &result)
}

// AddNotificationMicrosoftTeams adds a notification.
func AddNotificationMicrosoftTeams(ctx context.Context, input *schema.AddNotificationMicrosoftTeamsInput, n Notification) (*schema.NotificationMicrosoftTeams, error) {
	result := schema.NotificationMicrosoftTeams{}
	return &result, n.AddNotificationMicrosoftTeams(ctx, input, &result)
}

// AddNotificationSlack adds a notification.
func AddNotificationSlack(ctx context.Context, input *schema.AddNotificationSlackInput, n Notification) (*schema.NotificationSlack, error) {
	result := schema.NotificationSlack{}
	return &result, n.AddNotificationSlack(ctx, input, &result)
}

// AddNotificationToProject adds a notification to project.
func AddNotificationToProject(ctx context.Context, input *schema.AddNotificationToProjectInput, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.AddNotificationToProject(ctx, input, &result)
}
