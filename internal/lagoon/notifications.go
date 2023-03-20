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

	RemoveNotificationFromProject(context.Context, *schema.RemoveNotificationFromProjectInput, *schema.Project) error

	DeleteNotificationSlack(ctx context.Context, name string, project *schema.DeleteNotification) error
	DeleteNotificationRocketChat(ctx context.Context, name string, project *schema.DeleteNotification) error
	DeleteNotificationMicrosoftTeams(ctx context.Context, name string, project *schema.DeleteNotification) error
	DeleteNotificationEmail(ctx context.Context, name string, project *schema.DeleteNotification) error
	DeleteNotificationWebhook(ctx context.Context, name string, project *schema.DeleteNotification) error

	UpdateNotificationWebhook(ctx context.Context, input *schema.UpdateNotificationWebhookInput, result *schema.NotificationWebhook) error
	UpdateNotificationEmail(ctx context.Context, input *schema.UpdateNotificationEmailInput, result *schema.NotificationEmail) error
	UpdateNotificationRocketChat(ctx context.Context, input *schema.UpdateNotificationRocketChatInput, result *schema.NotificationRocketChat) error
	UpdateNotificationMicrosoftTeams(ctx context.Context, input *schema.UpdateNotificationMicrosoftTeamsInput, result *schema.NotificationMicrosoftTeams) error
	UpdateNotificationSlack(ctx context.Context, input *schema.UpdateNotificationSlackInput, result *schema.NotificationSlack) error

	GetAllNotificationEmail(ctx context.Context, project *[]schema.Project) error
	GetAllNotificationWebhook(ctx context.Context, project *[]schema.Project) error
	GetAllNotificationMicrosoftTeams(ctx context.Context, project *[]schema.Project) error
	GetAllNotificationSlack(ctx context.Context, project *[]schema.Project) error
	GetAllNotificationRocketChat(ctx context.Context, project *[]schema.Project) error

	GetProjectNotificationSlack(ctx context.Context, name string, project *schema.Project) error
	GetProjectNotificationRocketChat(ctx context.Context, name string, project *schema.Project) error
	GetProjectNotificationMicrosoftTeams(ctx context.Context, name string, project *schema.Project) error
	GetProjectNotificationEmail(ctx context.Context, name string, project *schema.Project) error
	GetProjectNotificationWebhook(ctx context.Context, name string, project *schema.Project) error
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

// RemoveNotificationFromProject removes a notification from a project.
func RemoveNotificationFromProject(ctx context.Context, input *schema.RemoveNotificationFromProjectInput, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.RemoveNotificationFromProject(ctx, input, &result)
}

// UpdateNotificationWebhook updates a notification.
func UpdateNotificationWebhook(ctx context.Context, input *schema.UpdateNotificationWebhookInput, n Notification) (*schema.NotificationWebhook, error) {
	result := schema.NotificationWebhook{}
	return &result, n.UpdateNotificationWebhook(ctx, input, &result)
}

// UpdateNotificationEmail updates a notification.
func UpdateNotificationEmail(ctx context.Context, input *schema.UpdateNotificationEmailInput, n Notification) (*schema.NotificationEmail, error) {
	result := schema.NotificationEmail{}
	return &result, n.UpdateNotificationEmail(ctx, input, &result)
}

// UpdateNotificationRocketChat updates a notification.
func UpdateNotificationRocketChat(ctx context.Context, input *schema.UpdateNotificationRocketChatInput, n Notification) (*schema.NotificationRocketChat, error) {
	result := schema.NotificationRocketChat{}
	return &result, n.UpdateNotificationRocketChat(ctx, input, &result)
}

// UpdateNotificationMicrosoftTeams updates a notification.
func UpdateNotificationMicrosoftTeams(ctx context.Context, input *schema.UpdateNotificationMicrosoftTeamsInput, n Notification) (*schema.NotificationMicrosoftTeams, error) {
	result := schema.NotificationMicrosoftTeams{}
	return &result, n.UpdateNotificationMicrosoftTeams(ctx, input, &result)
}

// UpdateNotificationSlack updates a notification.
func UpdateNotificationSlack(ctx context.Context, input *schema.UpdateNotificationSlackInput, n Notification) (*schema.NotificationSlack, error) {
	result := schema.NotificationSlack{}
	return &result, n.UpdateNotificationSlack(ctx, input, &result)
}

// GetAllNotificationEmail gets all notifications of type.
func GetAllNotificationEmail(ctx context.Context, n Notification) (*[]schema.Project, error) {
	result := []schema.Project{}
	return &result, n.GetAllNotificationEmail(ctx, &result)
}

// GetAllNotificationWebhook gets all notifications of type.
func GetAllNotificationWebhook(ctx context.Context, n Notification) (*[]schema.Project, error) {
	result := []schema.Project{}
	return &result, n.GetAllNotificationWebhook(ctx, &result)
}

// GetAllNotificationSlack gets all notifications of type.
func GetAllNotificationSlack(ctx context.Context, n Notification) (*[]schema.Project, error) {
	result := []schema.Project{}
	return &result, n.GetAllNotificationSlack(ctx, &result)
}

// GetAllNotificationRocketChat gets all notifications of type.
func GetAllNotificationRocketChat(ctx context.Context, n Notification) (*[]schema.Project, error) {
	result := []schema.Project{}
	return &result, n.GetAllNotificationRocketChat(ctx, &result)
}

// GetAllNotificationMicrosoftTeams gets all notifications of type.
func GetAllNotificationMicrosoftTeams(ctx context.Context, n Notification) (*[]schema.Project, error) {
	result := []schema.Project{}
	return &result, n.GetAllNotificationMicrosoftTeams(ctx, &result)
}

// GetProjectNotificationEmail gets all notifications of type in project.
func GetProjectNotificationEmail(ctx context.Context, name string, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.GetProjectNotificationEmail(ctx, name, &result)
}

// GetProjectNotificationWebhook gets all notifications of type in project.
func GetProjectNotificationWebhook(ctx context.Context, name string, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.GetProjectNotificationWebhook(ctx, name, &result)
}

// GetProjectNotificationRocketChat gets all notifications of type in project.
func GetProjectNotificationRocketChat(ctx context.Context, name string, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.GetProjectNotificationRocketChat(ctx, name, &result)
}

// GetProjectNotificationSlack gets all notifications of type in project.
func GetProjectNotificationSlack(ctx context.Context, name string, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.GetProjectNotificationSlack(ctx, name, &result)
}

// GetProjectNotificationMicrosoftTeams gets all notifications of type in project.
func GetProjectNotificationMicrosoftTeams(ctx context.Context, name string, n Notification) (*schema.Project, error) {
	result := schema.Project{}
	return &result, n.GetProjectNotificationMicrosoftTeams(ctx, name, &result)
}

// DeleteNotificationEmail deletes notification.
func DeleteNotificationEmail(ctx context.Context, name string, n Notification) (*schema.DeleteNotification, error) {
	result := schema.DeleteNotification{}
	return &result, n.DeleteNotificationEmail(ctx, name, &result)
}

// DeleteNotificationWebhook deletes notification.
func DeleteNotificationWebhook(ctx context.Context, name string, n Notification) (*schema.DeleteNotification, error) {
	result := schema.DeleteNotification{}
	return &result, n.DeleteNotificationWebhook(ctx, name, &result)
}

// DeleteNotificationRocketChat deletes notification.
func DeleteNotificationRocketChat(ctx context.Context, name string, n Notification) (*schema.DeleteNotification, error) {
	result := schema.DeleteNotification{}
	return &result, n.DeleteNotificationRocketChat(ctx, name, &result)
}

// DeleteNotificationSlack deletes notification.
func DeleteNotificationSlack(ctx context.Context, name string, n Notification) (*schema.DeleteNotification, error) {
	result := schema.DeleteNotification{}
	return &result, n.DeleteNotificationSlack(ctx, name, &result)
}

// DeleteNotificationMicrosoftTeams deletes notification.
func DeleteNotificationMicrosoftTeams(ctx context.Context, name string, n Notification) (*schema.DeleteNotification, error) {
	result := schema.DeleteNotification{}
	return &result, n.DeleteNotificationMicrosoftTeams(ctx, name, &result)
}
