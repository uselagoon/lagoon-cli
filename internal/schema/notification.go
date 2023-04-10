package schema

import (
	"encoding/json"
	"fmt"
)

// AddNotificationRocketChatInput is based on the Lagoon API type.
type AddNotificationRocketChatInput struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// UpdateNotificationRocketChatPatchInput is based on the Lagoon API type.
type UpdateNotificationRocketChatInput struct {
	Name  string                         `json:"name"`
	Patch AddNotificationRocketChatInput `json:"patch"`
}

// NotificationRocketChat is based on the Lagoon API type.
type NotificationRocketChat struct {
	AddNotificationRocketChatInput
	ID uint `json:"id,omitempty"`
}

// AddNotificationSlackInput is based on the Lagoon API type.
type AddNotificationSlackInput struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
	Channel string `json:"channel,omitempty"`
}

// UpdateNotificationSlackPatchInput is based on the Lagoon API type.
type UpdateNotificationSlackInput struct {
	Name  string                    `json:"name"`
	Patch AddNotificationSlackInput `json:"patch"`
}

// NotificationSlack is based on the Lagoon API type.
type NotificationSlack struct {
	AddNotificationSlackInput
	ID uint `json:"id,omitempty"`
}

// AddNotificationEmailInput is based on the Lagoon API type.
type AddNotificationEmailInput struct {
	Name         string `json:"name,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}

// UpdateNotificationEmailPatchInput is based on the Lagoon API type.
type UpdateNotificationEmailInput struct {
	Name  string                    `json:"name"`
	Patch AddNotificationEmailInput `json:"patch"`
}

// NotificationEmail is based on the Lagoon API type.
type NotificationEmail struct {
	AddNotificationEmailInput
	ID uint `json:"id,omitempty"`
}

// AddNotificationMicrosoftTeamsInput is based on the Lagoon API type.
type AddNotificationMicrosoftTeamsInput struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
}

// UpdateNotificationMicrosoftTeamsPatchInput is based on the Lagoon API type.
type UpdateNotificationMicrosoftTeamsInput struct {
	Name  string                             `json:"name"`
	Patch AddNotificationMicrosoftTeamsInput `json:"patch"`
}

// NotificationMicrosoftTeams is based on the Lagoon API type.
type NotificationMicrosoftTeams struct {
	AddNotificationMicrosoftTeamsInput
	ID uint `json:"id,omitempty"`
}

// AddNotificationWebhookInput is based on the Lagoon API type.
type AddNotificationWebhookInput struct {
	Name    string `json:"name,omitempty"`
	Webhook string `json:"webhook,omitempty"`
}

// UpdateNotificationWebhookPatchInput is based on the Lagoon API type.
type UpdateNotificationWebhookInput struct {
	Name  string                      `json:"name"`
	Patch AddNotificationWebhookInput `json:"patch"`
}

// NotificationWebhook is based on the Lagoon API type.
type NotificationWebhook struct {
	AddNotificationWebhookInput
	ID uint `json:"id,omitempty"`
}

// DeleteNotification is the response.
type DeleteNotification struct {
	DeleteNotification string `json:"deleteNotification"`
}

// Notifications represents possible Lagoon notification types.
// These are unmarshalled from a projectByName query response.
type Notifications struct {
	Slack          []AddNotificationSlackInput          `json:"slack,omitempty"`
	RocketChat     []AddNotificationRocketChatInput     `json:"rocketchat,omitempty"`
	Email          []AddNotificationEmailInput          `json:"email,omitempty"`
	MicrosoftTeams []AddNotificationMicrosoftTeamsInput `json:"microsoftteams,omitempty"`
	Webhook        []AddNotificationWebhookInput        `json:"webhook,omitempty"`
}

// NotificationsConfig represents possible Lagoon notification types and
// (un)marshals to the config file format.
type NotificationsConfig struct {
	Notifications
}

// UnmarshalJSON unmashals a quoted json string to the Notification values
// returned from the Lagoon API.
func (n *Notifications) UnmarshalJSON(b []byte) error {
	var nArray []map[string]string
	err := json.Unmarshal(b, &nArray)
	if err != nil {
		return err
	}
	for _, nMap := range nArray {
		if len(nMap) == 0 {
			// Unsupported notification type returns an empty map.
			// This happens when the lagoon API being targeted is actually a higher
			// version than configured.
			continue
		}
		switch nMap["__typename"] {
		case "NotificationSlack":
			n.Slack = append(n.Slack,
				AddNotificationSlackInput{
					Name:    nMap["name"],
					Webhook: nMap["webhook"],
					Channel: nMap["channel"],
				})
		case "NotificationRocketChat":
			n.RocketChat = append(n.RocketChat,
				AddNotificationRocketChatInput{
					Name:    nMap["name"],
					Webhook: nMap["webhook"],
					Channel: nMap["channel"],
				})
		case "NotificationEmail":
			n.Email = append(n.Email,
				AddNotificationEmailInput{
					Name:         nMap["name"],
					EmailAddress: nMap["emailAddress"],
				})
		case "NotificationMicrosoftTeams":
			n.MicrosoftTeams = append(n.MicrosoftTeams,
				AddNotificationMicrosoftTeamsInput{
					Name:    nMap["name"],
					Webhook: nMap["webhook"],
				})
		case "NotificationWebhook":
			n.Webhook = append(n.Webhook,
				AddNotificationWebhookInput{
					Name:    nMap["name"],
					Webhook: nMap["webhook"],
				})
		default:
			return fmt.Errorf("unknown notification type: %v", nMap["__typename"])
		}
	}
	return nil
}

// MarshalJSON marshals the Notifications as a quoted json string.
func (n *NotificationsConfig) MarshalJSON() ([]byte, error) {
	nMap := map[string][]map[string]string{}
	for _, slack := range n.Slack {
		nMap["slack"] = append(nMap["slack"], map[string]string{
			"name":    slack.Name,
			"webhook": slack.Webhook,
			"channel": slack.Channel,
		})
	}
	for _, rocketChat := range n.RocketChat {
		nMap["rocketChat"] = append(nMap["rocketChat"], map[string]string{
			"name":    rocketChat.Name,
			"webhook": rocketChat.Webhook,
			"channel": rocketChat.Channel,
		})
	}
	for _, email := range n.Email {
		nMap["email"] = append(nMap["email"], map[string]string{
			"name":         email.Name,
			"emailAddress": email.EmailAddress,
		})
	}
	for _, microsoftTeams := range n.MicrosoftTeams {
		nMap["microsoftTeams"] = append(nMap["microsoftTeams"], map[string]string{
			"name":    microsoftTeams.Name,
			"webhook": microsoftTeams.Webhook,
		})
	}
	for _, webhook := range n.Webhook {
		nMap["webhook"] = append(nMap["webhook"], map[string]string{
			"name":    webhook.Name,
			"webhook": webhook.Webhook,
		})
	}
	return json.Marshal(nMap)
}

// UnmarshalJSON unmashals a quoted json string to the Notification values.
func (n *NotificationsConfig) UnmarshalJSON(b []byte) error {
	var nMap map[string][]map[string]string
	err := json.Unmarshal(b, &nMap)
	if err != nil {
		return err
	}
	for nType, nValues := range nMap {
		switch nType {
		case "slack":
			for _, slackMap := range nValues {
				n.Slack = append(n.Slack,
					AddNotificationSlackInput{
						Name:    slackMap["name"],
						Webhook: slackMap["webhook"],
						Channel: slackMap["channel"],
					})
			}
		case "rocketChat":
			for _, rocketChatMap := range nValues {
				n.RocketChat = append(n.RocketChat,
					AddNotificationRocketChatInput{
						Name:    rocketChatMap["name"],
						Webhook: rocketChatMap["webhook"],
						Channel: rocketChatMap["channel"],
					})
			}
		case "email":
			for _, emailMap := range nValues {
				n.Email = append(n.Email,
					AddNotificationEmailInput{
						Name:         emailMap["name"],
						EmailAddress: emailMap["emailAddress"],
					})
			}
		case "microsoftTeams":
			for _, microsoftTeamsMap := range nValues {
				n.MicrosoftTeams = append(n.MicrosoftTeams,
					AddNotificationMicrosoftTeamsInput{
						Name:    microsoftTeamsMap["name"],
						Webhook: microsoftTeamsMap["webhook"],
					})
			}
		case "webhook":
			for _, webhookMap := range nValues {
				n.Webhook = append(n.Webhook,
					AddNotificationWebhookInput{
						Name:    webhookMap["name"],
						Webhook: webhookMap["webhook"],
					})
			}
		default:
			return fmt.Errorf("unknown notification type: %v", nType)
		}
	}
	return nil
}
