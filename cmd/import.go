package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
)

var importCmd = &cobra.Command{
	Use:     "import",
	Aliases: []string{"i"},
	Hidden:  false,
	Short:   "Import a config from a yaml file",
	Long: `Import a config from a yaml file.
By default this command will exit on encountering an error (such as an existing object).
You can get it to continue anyway with --keep-going. To disable any prompts, use --force.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		importFile, err := cmd.Flags().GetString("import-file")
		if err != nil {
			return err
		}
		keepGoing, err := cmd.Flags().GetBool("keep-going")
		if err != nil {
			return err
		}
		openshiftID, err := cmd.Flags().GetUint("openshiftID")
		if err != nil {
			return err
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		if !yesNo(fmt.Sprintf(
			`Are you sure you want to import config from %s into "%s" lagoon?`,
			importFile, current)) {
			return nil // user cancelled
		}

		err = setConfigDefaultVersion(&lagoonCLIConfig, current, "1.0.0")
		if err != nil {
			return err
		}
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		file, err := os.Open(importFile)
		if err != nil {
			return fmt.Errorf("couldn't open file: %w", err)
		}

		return lagoon.Import(context.TODO(), lc, file, keepGoing, openshiftID)
	},
}

// convert a slice of strings to a set (as a map)
func sliceToMap(s []string) map[string]bool {
	m := map[string]bool{}
	for _, ss := range s {
		m[ss] = true
	}
	return m
}

var exportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"e"},
	Hidden:  false,
	Short:   "Export lagoon output to yaml",
	Long: `Export lagoon output to yaml
You must specify to export a specific project by using the '-p <project-name>' flag`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := cmd.Flags().GetString("project")
		if err != nil {
			return err
		}
		if len(project) == 0 {
			return fmt.Errorf("no project specified")
		}
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		exclude, err := cmd.Flags().GetStringSlice("exclude")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		if !yesNo(fmt.Sprintf(
			`Are you sure you want to export lagoon config for %s on "%s" lagoon?`,
			project, current)) {
			return nil // user cancelled
		}

		err = setConfigDefaultVersion(&lagoonCLIConfig, current, "1.0.0")
		if err != nil {
			return err
		}
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		conf, err := lagoon.ExportProject(
			context.TODO(), lc, project, sliceToMap(exclude))
		if err != nil {
			return err
		}

		_, err = fmt.Print(string(conf))
		return err
	},
}

func init() {
	importCmd.Flags().StringP("import-file", "I", "",
		"path to the file to import")
	importCmd.Flags().Bool("keep-going", false,
		"on error, just log and continue instead of aborting")
	importCmd.Flags().Uint("openshiftID", 0,
		"ID of the openshift to target for import")
	for _, flag := range []string{"import-file", "openshiftID"} {
		if err := importCmd.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}

	exportCmd.Flags().StringSlice("exclude", []string{"project-private-keys"},
		`Exclude data from the export. Valid options (others are ignored): users, project-users, groups, notifications, project-private-keys`)
}
