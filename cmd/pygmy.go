package cmd

import (
	"fmt"
	"os"

	pygmy "github.com/fubarhouse/pygmy-go/service/library"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var pygmyConfig pygmy.Config

var pygmyCmd = &cobra.Command{
	Use:   "pygmy",
	Short: "start, stop or check the status of pygmy",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var pygmyCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Stop and remove all pygmy services regardless of state",
	Long: `Useful for debugging or system cleaning, this command will
remove all pygmy containers but leave the images in-tact.
This command does not check if the containers are running
because other checks do for speed convenience.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Clean(pygmyConfig)
	},
}

var pygmyStopCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove all pygmy services",
	Long: `Check if any pygmy containers are running and removes
then if they are, it will not attempt to remove any
services which are not running.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Down(pygmyConfig)
	},
}

var pygmyRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart all pygmy containers.",
	Long:  `This command will trigger the Down and Up commands`,
	Run: func(cmd *cobra.Command, args []string) {

		pygmyConfig.Key, _ = cmd.Flags().GetString("key")
		pygmyConfig.SkipKey, _ = cmd.Flags().GetBool("no-addkey")
		pygmyConfig.SkipResolver, _ = cmd.Flags().GetBool("no-resolver")
		pygmy.Restart(pygmyConfig)
	},
}

var pygmyStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report status of the pygmy services",
	Long: `Loop through all of pygmy's services and identify the present state.
This includes the docker services, the resolver and SSH key status`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Status(pygmyConfig)
	},
}

var pygmyUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up pygmy services (dnsmasq, haproxy, mailhog, resolv, ssh-agent)",
	Long: `Launch Pygmy - a set of containers and a resolver with very specific
configurations designed for use with Amazee.io local development.
It includes dnsmasq, haproxy, mailhog, resolv and ssh-agent.`,
	Run: func(cmd *cobra.Command, args []string) {

		pygmyConfig.Key, _ = cmd.Flags().GetString("key")
		pygmyConfig.SkipKey, _ = cmd.Flags().GetBool("no-addkey")
		pygmyConfig.SkipResolver, _ = cmd.Flags().GetBool("no-resolver")
		pygmy.Up(pygmyConfig)
	},
}

var pygmyUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pulls Docker Images and recreates the Containers",
	Long: `Pull all images Pygmy uses, as well as any images containing
the string 'amazeeio', which encompasses all lagoon images.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Update(pygmyConfig)
	},
}

var pygmyVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "# Check current installed version of pygmy",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Version(pygmyConfig)
	},
}

func init() {
	pygmyCmd.AddCommand(pygmyCleanCmd)
	pygmyCmd.AddCommand(pygmyRestartCmd)
	pygmyCmd.AddCommand(pygmyStatusCmd)
	pygmyCmd.AddCommand(pygmyStopCmd)
	pygmyCmd.AddCommand(pygmyUpCmd)
	pygmyCmd.AddCommand(pygmyUpdateCmd)
	pygmyCmd.AddCommand(pygmyVersionCmd)

	homedir, _ := homedir.Dir()
	keypath := fmt.Sprintf("%v%v.ssh%vid_rsa", homedir, string(os.PathSeparator), string(os.PathSeparator))

	pygmyUpCmd.Flags().StringP("key", "", keypath, "Path of SSH key to add")
	pygmyUpCmd.Flags().BoolP("no-addkey", "", false, "Skip adding the SSH key")
	pygmyUpCmd.Flags().BoolP("no-resolver", "", false, "Skip adding or removing the Resolver")

	pygmyRestartCmd.Flags().StringP("key", "", keypath, "Path of SSH key to add")
	pygmyRestartCmd.Flags().BoolP("no-addkey", "", false, "Skip adding the SSH key")
	pygmyRestartCmd.Flags().BoolP("no-resolver", "", false, "Skip adding or removing the Resolver")
}
