package cmd

import (
	"fmt"
	"os"

	pygmy "github.com/fubarhouse/pygmy-go/service/library"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var devConfig pygmy.Config

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "start, stop or check the status of dev",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var devCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Stop and remove all pygmy services regardless of state",
	Long: `Useful for debugging or system cleaning, this command will
remove all pygmy containers but leave the images in-tact.
This command does not check if the containers are running
because other checks do for speed convenience.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Clean(devConfig)
	},
}

var devStopCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop and remove all pygmy services",
	Long: `Check if any pygmy containers are running and removes
then if they are, it will not attempt to remove any
services which are not running.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Down(devConfig)
	},
}

var devRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart all pygmy containers.",
	Long:  `This command will trigger the Down and Up commands`,
	Run: func(cmd *cobra.Command, args []string) {

		devConfig.Key, _ = cmd.Flags().GetString("key")
		devConfig.SkipKey, _ = cmd.Flags().GetBool("no-addkey")
		devConfig.SkipResolver, _ = cmd.Flags().GetBool("no-resolver")
		pygmy.Restart(devConfig)
	},
}

var devStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report status of the pygmy services",
	Long: `Loop through all of pygmy's services and identify the present state.
This includes the docker services, the resolver and SSH key status`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Status(devConfig)
	},
}

var devUpCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up pygmy services (dnsmasq, haproxy, mailhog, resolv, ssh-agent)",
	Long: `Launch Pygmy - a set of containers and a resolver with very specific
configurations designed for use with Amazee.io local development.
It includes dnsmasq, haproxy, mailhog, resolv and ssh-agent.`,
	Run: func(cmd *cobra.Command, args []string) {

		devConfig.Key, _ = cmd.Flags().GetString("key")
		devConfig.SkipKey, _ = cmd.Flags().GetBool("no-addkey")
		devConfig.SkipResolver, _ = cmd.Flags().GetBool("no-resolver")
		pygmy.Up(devConfig)
	},
}

var devUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Pulls Docker Images and recreates the Containers",
	Long: `Pull all images Pygmy uses, as well as any images containing
the string 'amazeeio', which encompasses all lagoon images.`,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Update(devConfig)
	},
}

var devVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "# Check current installed version of Pygmy",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pygmy.Version(devConfig)
	},
}

func init() {
	devCmd.AddCommand(pygmyCleanCmd)
	devCmd.AddCommand(pygmyRestartCmd)
	devCmd.AddCommand(pygmyStatusCmd)
	devCmd.AddCommand(pygmyStopCmd)
	devCmd.AddCommand(pygmyUpCmd)
	devCmd.AddCommand(pygmyUpdateCmd)
	devCmd.AddCommand(pygmyVersionCmd)

	homedir, _ := homedir.Dir()
	keypath := fmt.Sprintf("%v%v.ssh%vid_rsa", homedir, string(os.PathSeparator), string(os.PathSeparator))

	pygmyUpCmd.Flags().StringP("key", "", keypath, "Path of SSH key to add")
	pygmyUpCmd.Flags().BoolP("no-addkey", "", false, "Skip adding the SSH key")
	pygmyUpCmd.Flags().BoolP("no-resolver", "", false, "Skip adding or removing the Resolver")

	pygmyRestartCmd.Flags().StringP("key", "", keypath, "Path of SSH key to add")
	pygmyRestartCmd.Flags().BoolP("no-addkey", "", false, "Skip adding the SSH key")
	pygmyRestartCmd.Flags().BoolP("no-resolver", "", false, "Skip adding or removing the Resolver")
}
