package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
)

var rawCmd = &cobra.Command{
	Use:     "raw",
	Aliases: []string{"r"},
	Short:   "Run a custom query or mutation",
	Long: `Run a custom query or mutation.
The output of this command will be the JSON response from the API`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		raw, err := cmd.Flags().GetString("raw")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Raw query or mutation", raw); err != nil {
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
		if err != nil {
			return err
		}
		rawResp, err := lc.ProcessRaw(context.TODO(), raw, nil)
		if err != nil {
			return err
		}
		r, err := json.Marshal(rawResp)
		if err != nil {
			return err
		}
		fmt.Println(string(r))
		return nil
	},
}

func init() {
	rawCmd.Flags().String("raw", "", "The raw query or mutation to run")
}
