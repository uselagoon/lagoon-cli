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
