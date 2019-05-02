package cmd

import (
	"fmt"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var kibanaCmd = &cobra.Command{
	Use:   "kibana",
	Short: "Launch the web user interface",
	Run: func(cmd *cobra.Command, args []string) {
		url := viper.GetString("lagoon_kibana")
		fmt.Println(fmt.Sprintf("Opening %s", url))
		_ = browser.OpenURL(url)
	},
}

func init() {
	rootCmd.AddCommand(kibanaCmd)
}
