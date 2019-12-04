package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/importer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showExample bool
var example = `groups:
  - name: example-com
users:
  - user:
      email: usera@example.com
      sshkey: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC+qXkqFFmZiQzEE9yhwEbHSuge187PRbH+MhEo3xyKpSODSa5Arp3P/ZPjxOOPm/KU2m9S/k4MKwvtZehqQtTDss9KaDZcZmwdYQSyJDgUzcAcMtEAFO8+9bVDIzgRXsRwsgkaJM627BRrutjpaaAryMMQJyuF/QMDwAfZ6ZBKNw== usera@macbook.local
    groups:
      - name: example-com
        role: owner
  - user:
      email: userb@example.com
      sshkey: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC19Ap4udFpuIW7OGP3E0xSMr30c56uzp26thVmonvPWunU6pwcdV/hGioShSSUEWZdUwPhhCkUG4c/IrD0g6W6ZG/8JoLkUkIQ0OMSRl6zDGy0woTrbURyktwHLfzV12WLFmJR6l7egHov/FgNVx4gRRvVLIEwfpbGHvNCG6lQnQ== userb@macbook.local
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
	Short:   "Import a config from a yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		if showExample {
			fmt.Println(example)
			os.Exit(0)
		}
		if importFile == "" {
			fmt.Println("Not enough arguments. Requires: path to file to import")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo("Are you sure you want to import this file, it is potentially dangerous") {
			importer.ImportData(importFile, forceAction)
		}
	},
}

func init() {
	importCmd.Flags().BoolVarP(&showExample, "example", "", false, "display example yaml")
	importCmd.Flags().StringVarP(&importFile, "import", "I", "", "path to the file to import")
}
