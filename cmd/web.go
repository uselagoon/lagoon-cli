package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"w"},
	Short:   "Launch the web user interface",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			_ = cmd.Help()
			os.Exit(1)
		}

		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].UI)
		if lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].UI != "" {
			urlBuilder.WriteString(fmt.Sprintf("/projects/%s", cmdProjectName))
		} else {
			handleError(fmt.Errorf("unable to determine url for ui, is one set?"))
		}

		url := urlBuilder.String()
		fmt.Printf("Opening %s\n", url)
		_ = browser.OpenURL(url)
	},
}

var kibanaCmd = &cobra.Command{
	Use:     "kibana",
	Aliases: []string{"k"},
	Short:   "Launch the kibana interface",
	Run: func(cmd *cobra.Command, args []string) {
		urlBuilder := strings.Builder{}
		urlBuilder.WriteString(lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].Kibana)
		if lagoonCLIConfig.Lagoons[lagoonCLIConfig.Current].Kibana == "" {
			handleError(fmt.Errorf("unable to determine url for kibana, is one set?"))
		}

		url := urlBuilder.String()
		fmt.Printf("Opening %s\n", url)
		_ = browser.OpenURL(url)
	},
}
