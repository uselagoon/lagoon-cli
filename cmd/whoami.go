package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Aliases: []string{"w"},
	Hidden:  false,
	Short:   "Whoami will return your user information for lagoon",
	Long: `Whoami will return your user information for lagoon. 
This is useful if you have multiple keys or accounts in multiple lagoons and need to check which you are using.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)

		user, err := lagoon.GetMeInfo(context.TODO(), lc)
		if err != nil {
			if strings.Contains(err.Error(), "Cannot read property 'access_token' of null") {
				return fmt.Errorf("Unable to get user information, you may be using an administration token")
			}
			return err
		}

		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Email",
				"FirstName",
				"LastName",
			},
			Data: []output.Data{
				[]string{
					returnNonEmptyString(fmt.Sprintf("%v", user.ID)),
					returnNonEmptyString(user.Email),
					returnNonEmptyString(user.FirstName),
					returnNonEmptyString(user.LastName),
				},
			},
		}, outputOptions)

		return nil
	},
}

var canISSHCmd = &cobra.Command{
	Use:    "can-i-ssh", //@TODO: this may be annoying to type, open to suggestions
	Hidden: false,
	Short:  "Can I SSH will return the environment if you can access it",
	Long: `Can I SSH will return the environment if you can access it.
This is useful if you want to quickly check if you can SSH to an environment in lagoon.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(viper.GetString("current"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			return fmt.Errorf("Missing arguments: Project name or environment name is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		environment, err := lagoon.CanUserSSHToEnvironment(context.TODO(), fmt.Sprintf("%s-%s", cmdProjectName, sanitizeName(cmdProjectEnvironment)), lc)
		if err != nil {
			if strings.Contains(err.Error(), "Cannot read property 'access_token' of null") {
				return fmt.Errorf("Unable to get user information, you may be using an administration token")
			}
			return err
		}
		fmt.Println(environment.Name)
		return nil
	},
}
