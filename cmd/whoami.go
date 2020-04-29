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

		showKeys, err := cmd.Flags().GetBool("keys")
		if err != nil {
			return err
		}

		showFingerprints, err := cmd.Flags().GetBool("fingerprints")
		if err != nil {
			return err
		}

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

		if showKeys || showFingerprints {
			// if we are only showing the users keys, leave the users email visible as part
			// of helping identify the user stil
			var keys []output.Data
			for _, key := range user.SSHKeys {
				if showFingerprints {
					keys = append(keys, []string{user.Email, key.Name, key.Created, fmt.Sprintf("%s", key.KeyType), key.KeyFingerprint})
				} else {
					keys = append(keys, []string{user.Email, key.Name, key.Created, fmt.Sprintf("%s", key.KeyType), key.KeyValue})
				}
			}
			output.RenderOutput(output.Table{
				Header: []string{
					"Email",
					"Name",
					"Created",
					"Type",
					"Value",
				},
				Data: keys,
			}, outputOptions)
		} else {
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Email",
					"FirstName",
					"LastName",
					"SSHKeys",
				},
				Data: []output.Data{
					[]string{
						returnNonEmptyString(fmt.Sprintf("%v", user.ID)),
						returnNonEmptyString(user.Email),
						returnNonEmptyString(user.FirstName),
						returnNonEmptyString(user.LastName),
						returnNonEmptyString(fmt.Sprintf("%v", len(user.SSHKeys))),
					},
				},
			}, outputOptions)
		}

		return nil
	},
}

func init() {
	whoamiCmd.Flags().Bool("keys", false,
		"Display your SSH keys")
	whoamiCmd.Flags().Bool("fingerprints", false,
		"Display your SSH keys fingerprints")
}
