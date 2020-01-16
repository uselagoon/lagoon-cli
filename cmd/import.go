package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/importer"
	"github.com/amazeeio/lagoon-cli/lagoon/parser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showExample bool
var example = `groups:
  - name: example-com
users:
  - user:
      email: usera@example.com
      sshkey: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIG/1WiXC+XSpGQsTBYhWMy8WCIIGGtq26GKHeXy9vySf usera@macbook.local
    groups:
      - name: example-com
        role: owner
  - user:
      email: userb@example.com
      sshkey: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIJ3qUs3GlmILI4ozHhCXPVq1WEv3gb0Mtc5FGu4l+qCl userb@macbook.local
    groups:
      - name: example-com
        role: developer
slack:
  - name: example-com-slack
    webhook: https://slack.com/webhook
    channel: example-com
rocketchat:
  - name: example-com-rocketchat
    webhook: https://rocketchat.com/webhook
    channel: example-com
  - name: example-com-api-rocketchat
    webhook: https://rocketchat.com/webhook
    channel: example-com-api
projects:
  - project:
      name: example-com
      giturl: "git@github.com:example/example-com.git"
      openshift: 1
      branches: master|develop|staging
      productionenvironment: master
    notifications:
      slack:
        - example-com-slack
      rocketchat:
        - example-com-rocketchat
    groups:
      - example-com
  - project:
      name: example-com-api
      giturl: "git@github.com:example/example-com-api.git"
      openshift: 1
      branches: master|develop|staging
      productionenvironment: master
    notifications:
      rocketchat:
        - example-com-api-rocketchat
    groups:
      - example-com`

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Hidden:  true,
	Short:   "Import a config from a yaml file",
	Long: `Import a config from a yaml file or from stdin, if using a file you will be prompted if you want to continue if you encounter any errors.
With stdin, there are no prompts and imports will fail if there are any errors.
You can avoid these and force it to continue on errors by specifying the '--force' flag`,
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		if showExample {
			fmt.Println(example)
			os.Exit(0)
		}
		importData, err := readStdInOrFile(importFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// if using stdin, then we need to force action as any prompts just fail as they expect stdin
			importer.ImportData(importData, forceAction)
		} else {
			// else we can prompt for failures
			if yesNo("Are you sure you want to import this data, it is potentially dangerous") {
				importer.ImportData(importData, forceAction)
			}
		}
	},
}

var parseCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p"},
	Hidden:  true,
	Short:   "Parse lagoon output to import yaml",
	Long: `Parse lagoon output to import yaml
If given the raw JSON output from a lagoon query, this will parse it into a yaml format that can then be used to import.`,
	Run: func(cmd *cobra.Command, args []string) {
		importJSON, err := readStdInOrFile(importFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		parser.ParseJSONImport(importJSON)
	},
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"e"},
	Hidden:  true,
	Short:   "Export lagoon output to yaml",
	Long: `Export lagoon output to yaml
You can specify to export a specific project by using the '-p <project-name>' flag`,
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		if cmdProjectName == "" {
			if yesNo("Are you sure you want to export lagoon output for all projects?") {
				_, _ = parser.ParseAllProjects()
				// fmt.Println(string(data))
			}
		} else {
			if yesNo("Are you sure you want to export lagoon output for " + cmdProjectName + "?") {
				_, _ = parser.ParseProject(cmdProjectName)
			}
		}

	},
}

func init() {
	importCmd.Flags().BoolVarP(&showExample, "example", "", false, "display example yaml")
	importCmd.Flags().StringVarP(&importFile, "import", "I", "", "path to the file to import")
	parseCmd.Flags().StringVarP(&importFile, "import", "I", "", "path to the file to import")
}
