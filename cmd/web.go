package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Launch the web user interface",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Not enough arguments. Requires: project name")
			os.Exit(1)
		}
		projectName := args[0]

		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(viper.GetString("lagoons." + cmdLagoon + ".ui"))
		urlBuilder.WriteString(fmt.Sprintf("/project?name=%s", projectName))

		url := urlBuilder.String()
		fmt.Println(fmt.Sprintf("Opening %s", url))
		_ = browser.OpenURL(url)
	},
}
