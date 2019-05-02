package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Launch the web user interface",
	Run: func(cmd *cobra.Command, args []string) {
		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(viper.GetString("lagoon_ui"))
		if cmdProjectName != "" {
			urlBuilder.WriteString(fmt.Sprintf("/project?name=%s", cmdProjectName))
		}

		url := urlBuilder.String()
		fmt.Println(fmt.Sprintf("Opening %s", url))
		_ = browser.OpenURL(url)
	},
}
