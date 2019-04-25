package cmd

import (
	"github.com/spf13/viper"
)

func ValidateToken() bool {
	apiToken := viper.GetString("lagoon_token")
	return apiToken != ""
}
