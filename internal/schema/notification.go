package schema

import (
	"encoding/json"
	"fmt"
)

// AddNotificationRocketChatInput is based on the Lagoon API type.
type AddNotificationRocketChatInput struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

// NotificationRocketChat is based on the Lagoon API type.
type NotificationRocketChat struct {
	AddNotificationRocketChatInput
	ID uint `json:"id,omitempty"`
}

// AddNotificationSlackInput is based on the Lagoon API type.
type AddNotificationSlackInput struct {
	Name    string `json:"name"`
	Webhook string `json:"webhook"`
	Channel string `json:"channel"`
}

// NotificationSlack is based on the Lagoon API type.
type NotificationSlack struct {
	AddNotificationSlackInput
	ID uint `json:"id,omitempty"`
}

// Notifications represents possible Lagoon notification types.
type Notifications struct {
	Slack      []AddNotificationSlackInput
	RocketChat []AddNotificationRocketChatInput
}

// NotificationsConfig represents possible Lagoon notification types and
// (un)marshals to the config file format.
type NotificationsConfig struct {
	Notifications
}

// UnmarshalJSON unmashals a quoted json string to the Notification values.
func (n *Notifications) UnmarshalJSON(b []byte) error {
	var nArray []map[string]string
	err := json.Unmarshal(b, &nArray)
	if err != nil {
		return err
	}
	for _, nMap := range nArray {
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
		default:
			return fmt.Errorf("unknown notification type: %v", nType)
		}
	}
	return nil
}
