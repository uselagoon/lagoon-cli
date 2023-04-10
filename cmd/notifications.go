package cmd

import (
	"encoding/json"

	"github.com/spf13/pflag"
)

// NotificationFlags .
type NotificationFlags struct {
	Project             string `json:"project,omitempty"`
	NotificationName    string `json:"name,omitempty"`
	NotificationNewName string `json:"newname,omitempty"`
	NotificationOldName string `json:"old,omitempty"`
	NotificationWebhook string `json:"webhook,omitempty"`
	NotificationChannel string `json:"channel,omitempty"`
}

func parseNotificationFlags(flags pflag.FlagSet) NotificationFlags {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := NotificationFlags{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

func init() {
	addNotificationCmd.AddCommand(addProjectNotificationRocketChatCmd)
	addNotificationCmd.AddCommand(addProjectNotificationSlackCmd)
	addNotificationCmd.AddCommand(addProjectNotificationEmailCmd)
	addNotificationCmd.AddCommand(addProjectNotificationMicrosoftTeamsCmd)
	addNotificationCmd.AddCommand(addProjectNotificationWebhookCmd)

	addNotificationCmd.AddCommand(addNotificationRocketchatCmd)
	addNotificationCmd.AddCommand(addNotificationSlackCmd)
	addNotificationCmd.AddCommand(addNotificationEmailCmd)
	addNotificationCmd.AddCommand(addNotificationMicrosoftTeamsCmd)
	addNotificationCmd.AddCommand(addNotificationWebhookCmd)

	listNotificationCmd.AddCommand(listProjectSlacksCmd)
	listNotificationCmd.AddCommand(listProjectRocketChatsCmd)
	listNotificationCmd.AddCommand(listProjectEmailsCmd)
	listNotificationCmd.AddCommand(listProjectMicrosoftTeamsCmd)
	listNotificationCmd.AddCommand(listProjectWebhooksCmd)

	listNotificationCmd.AddCommand(listAllSlacksCmd)
	listNotificationCmd.AddCommand(listAllRocketChatsCmd)
	listNotificationCmd.AddCommand(listAllEmailsCmd)
	listNotificationCmd.AddCommand(listAllMicrosoftTeamsCmd)
	listNotificationCmd.AddCommand(listAllWebhooksCmd)

	deleteNotificationCmd.AddCommand(deleteProjectRocketChatNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteProjectSlackNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteProjectEmailNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteProjectMicrosoftTeamsNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteProjectWebhookNotificationCmd)

	deleteNotificationCmd.AddCommand(deleteRocketChatNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteSlackNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteEmailNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteMicrosoftTeamsNotificationCmd)
	deleteNotificationCmd.AddCommand(deleteWebhookNotificationCmd)

	updateNotificationCmd.AddCommand(updateRocketChatNotificationCmd)
	updateNotificationCmd.AddCommand(updateSlackNotificationCmd)
	updateNotificationCmd.AddCommand(updateEmailNotificationCmd)
	updateNotificationCmd.AddCommand(updateMicrosoftTeamsNotificationCmd)
	updateNotificationCmd.AddCommand(updateWebhookNotificationCmd)
}
