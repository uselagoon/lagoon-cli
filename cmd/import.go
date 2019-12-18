package cmd

import (
	"bufio"
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
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		if showExample {
			fmt.Println(example)
			os.Exit(0)
		}
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// check if we are getting data froms stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		} else {
			// otherwise we can read from a file
			if importFile == "" {
				fmt.Println("Not enough arguments. Requires: path to file to import")
				cmd.Help()
				os.Exit(1)
			}
			if yesNo("Are you sure you want to import this file, it is potentially dangerous") {
				importer.ImportData(importFile, forceAction)
			}
		}
	},
}

var parseCmd = &cobra.Command{
	Use:     "parse",
	Aliases: []string{"p"},
	Hidden:  true,
	Short:   "Parse lagoon output to import yml",
	Run: func(cmd *cobra.Command, args []string) {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// check if we are getting data froms stdin
			fmt.Println("data is being piped to stdin")
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				fmt.Println(scanner.Text()) // Println will add back the final '\n'
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}
		} else {
			// otherwise we can read from a file
			if importFile == "" {
				fmt.Println("Not enough arguments. Requires: path to file to import")
				cmd.Help()
				os.Exit(1)
			}
			parser.ParseJSONImport(importFile)
		}
	},
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"e"},
	Hidden:  true,
	Short:   "Export lagoon output to yml",
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
