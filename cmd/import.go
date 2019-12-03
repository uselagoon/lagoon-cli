package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/importer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Short:   "Import a config from a yaml file",
	Long: `Import a YAML file, see following example:
groups:
  - name: example-com
users:
  - user:
      email: usera@example.com
      sshkey: ~/usera.pub
    groups:
      - name: example-com
        role: owner
  - user:
      email: userb@example.com
      sshkey: ~/userb.pub
    groups:
      - name: example-com
        role: developere
projects:
  - name: example-com
    giturl: git@github.com:example/example-com.git
    openshift: 1
    branches: "master|develop|staging"
    productionenvironment: master
    groups:
      - example-com
`,
	Run: func(cmd *cobra.Command, args []string) {
		validateToken(viper.GetString("current")) // get a new token if the current one is invalid
		if importFile == "" {
			fmt.Println("Not enough arguments. Requires: path to file to import")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo("Are you sure you want to import this file, it is potentially dangerous") {
			importer.ImportData(importFile)
		}
	},
}

func init() {
	importCmd.Flags().StringVarP(&importFile, "import", "I", "", "path to the file to import")
}
