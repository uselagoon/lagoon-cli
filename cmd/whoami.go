package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var whoamiCmd = &cobra.Command{
	Use:     "whoami",
	Aliases: []string{"w"},
	Hidden:  false,
	Short:   "Whoami will return your user information for lagoon",
	Long: `Whoami will return your user information for lagoon. 
This is useful if you have multiple keys or accounts in multiple lagoons and need to check which you are using.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		showOpts, err := cmd.Flags().GetStringSlice("show-keys")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		user, err := lagoon.Me(context.TODO(), lc)
		if err != nil {
			if strings.Contains(err.Error(), "Cannot read properties of null (reading 'access_token')") {
				return fmt.Errorf("unable to get user information, you may be using an administration token")
			}
			return err
		}

		if len(showOpts) > 0 {
			// if we are only showing the users keys, leave the users email visible as part
			// of helping identify the user stil
			opts := sliceToMap(showOpts)
			header := []string{
				"Email",
				"Name",
			}
			var keys []output.Data
			if opts["created"] {
				header = append(header, "Created")
			}
			if opts["type"] {
				header = append(header, "Type")
			}
			if opts["key"] {
				header = append(header, "Key")
			}
			if opts["fingerprint"] {
				header = append(header, "Fingerprint")
			}
			for _, key := range user.SSHKeys {
				keyData := []string{user.Email, key.Name}
				if opts["created"] {
					keyData = append(keyData, key.Created)
				}
				if opts["type"] {
					keyData = append(keyData, string(key.KeyType))
				}
				if opts["key"] {
					keyData = append(keyData, key.KeyValue)
				}
				if opts["fingerprint"] {
					keyData = append(keyData, key.KeyFingerprint)
				}
				keys = append(keys, keyData)
			}
			r := output.RenderOutput(output.Table{
				Header: header,
				Data:   keys,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		} else {
			r := output.RenderOutput(output.Table{
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
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
		}

		return nil
	},
}

func init() {
	whoamiCmd.Flags().StringSlice("show-keys", []string{},
		`Select which fields to display when showing SSH keys. Valid options (others are ignored): type,created,key,fingerprint`)
}
