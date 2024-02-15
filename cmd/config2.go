package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	config "github.com/uselagoon/machinery/utils/config"
	"golang.org/x/oauth2"
)

var config2Cmd = &cobra.Command{
	Use:     "config2",
	Aliases: []string{"c2"},
	Short:   "Configure Lagoon CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
}

var config2ListContextsCmd = &cobra.Command{
	Use:     "list-contexts",
	Aliases: []string{"lc"},
	Short:   "View all configured Lagoon contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := []output.Data{}
		for _, con := range c.Contexts {
			defa := false
			if con.Name == c.DefaultContext {
				defa = true
			}
			selected := ""
			if con.Name == cmdLagoon {
				selected = " (selected)"
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v%s", con.Name, selected)),
				returnNonEmptyString(fmt.Sprintf("%v", defa)),
				returnNonEmptyString(fmt.Sprintf("%v", con.User)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var config2ListUsersCmd = &cobra.Command{
	Use:     "list-users",
	Aliases: []string{"lu"},
	Short:   "View all configured Lagoon context users",
	RunE: func(cmd *cobra.Command, args []string) error {
		data := []output.Data{}
		for _, user := range c.Users {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

var config2AddContextCmd = &cobra.Command{
	Use:     "add-context",
	Aliases: []string{"ac"},
	Short:   "Add a new Lagoon context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		cUser, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}
		cAPIHost, err := cmd.Flags().GetString("api-hostname")
		if err != nil {
			return err
		}
		cTokenHost, err := cmd.Flags().GetString("token-hostname")
		if err != nil {
			return err
		}
		cTokenPort, err := cmd.Flags().GetInt("token-port")
		if err != nil {
			return err
		}
		cAuthHost, err := cmd.Flags().GetString("authentication-hostname")
		if err != nil {
			return err
		}
		// create the requested context
		cConfig := config.ContextConfig{
			Name: cName,
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "api-hostname" {
				cConfig.APIHostname = cAPIHost
			}
			if f.Name == "token-hostname" {
				cConfig.TokenHost = cTokenHost
			}
			if f.Name == "token-port" {
				cConfig.TokenPort = cTokenPort
			}
			if f.Name == "authentication-hostname" {
				cConfig.AuthenticationEndpoint = cAuthHost
			}
		})
		err = c.NewContext(cConfig, cUser)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range c.Contexts {
			defa := false
			if con.Name == c.DefaultContext {
				defa = true
			}
			if con.Name == cName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", con.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", defa)),
					returnNonEmptyString(fmt.Sprintf("%v", con.User)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
				})
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		return c.WriteConfig()
	},
}

var config2UpdateContextCmd = &cobra.Command{
	Use:     "update-context",
	Aliases: []string{"uc"},
	Short:   "Update a Lagoon context",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		cUser, err := cmd.Flags().GetString("user")
		if err != nil {
			return err
		}
		cAPIHost, err := cmd.Flags().GetString("api-hostname")
		if err != nil {
			return err
		}
		cTokenHost, err := cmd.Flags().GetString("token-hostname")
		if err != nil {
			return err
		}
		cTokenPort, err := cmd.Flags().GetInt("token-port")
		if err != nil {
			return err
		}
		cAuthHost, err := cmd.Flags().GetString("authentication-hostname")
		if err != nil {
			return err
		}
		// get the requested user
		cConfig, err := c.GetContext(cName)
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "api-hostname" {
				cConfig.ContextConfig.APIHostname = cAPIHost
			}
			if f.Name == "token-hostname" {
				cConfig.ContextConfig.TokenHost = cTokenHost
			}
			if f.Name == "token-port" {
				cConfig.ContextConfig.TokenPort = cTokenPort
			}
			if f.Name == "authentication-hostname" {
				cConfig.ContextConfig.AuthenticationEndpoint = cAuthHost
			}
			if f.Name == "user" {
				cConfig.User = cUser
			}
		})
		err = c.UpdateContext(cConfig.ContextConfig, cConfig.User)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range c.Contexts {
			defa := false
			if con.Name == c.DefaultContext {
				defa = true
			}
			if con.Name == cName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", con.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", defa)),
					returnNonEmptyString(fmt.Sprintf("%v", con.User)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
					returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
				})
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		return c.WriteConfig()
	},
}

var config2AddUserCmd = &cobra.Command{
	Use:     "add-user",
	Aliases: []string{"au"},
	Short:   "Add a new Lagoon context user",
	RunE: func(cmd *cobra.Command, args []string) error {
		uName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		uSSHKey, err := cmd.Flags().GetString("ssh-key")
		if err != nil {
			return err
		}
		// create the requested user
		uConfig := config.UserConfig{
			Name:  uName,
			Grant: &oauth2.Token{},
		}
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "ssh-key" {
				uConfig.SSHKey = uSSHKey
			}
		})
		err = c.NewUser(uConfig)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range c.Users {
			if user.Name == uName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
				})
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		return c.WriteConfig()
	},
}

var config2UpdateUserCmd = &cobra.Command{
	Use:     "update-user",
	Aliases: []string{"uu"},
	Short:   "Update a Lagoon context user",
	RunE: func(cmd *cobra.Command, args []string) error {
		uName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		uSSHKey, err := cmd.Flags().GetString("ssh-key")
		if err != nil {
			return err
		}
		// get the requested user
		u, err := c.GetUser(uName)
		// visit the flags and check for any defined flags to set
		cmd.Flags().Visit(func(f *pflag.Flag) {
			if f.Name == "ssh-key" {
				u.UserConfig.SSHKey = uSSHKey
			}
		})
		err = c.UpdateUser(u.UserConfig)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, user := range c.Users {
			if user.Name == uName {
				data = append(data, []string{
					returnNonEmptyString(fmt.Sprintf("%v", user.Name)),
					returnNonEmptyString(fmt.Sprintf("%v", user.UserConfig.SSHKey)),
				})
			}
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"SSH-Key",
			},
			Data: data,
		}, outputOptions)
		return c.WriteConfig()
	},
}

var config2GetConfigPathCmd = &cobra.Command{
	Use:     "config-path",
	Aliases: []string{"cp"},
	Short:   "Get the path of where the config file lives",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := config.Config{}
		path, err := c.GetConfigLocation()
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	},
}

var config2SetDefaultCmd = &cobra.Command{
	Use:     "default-context",
	Aliases: []string{"dc"},
	Short:   "Change which context is the default",
	RunE: func(cmd *cobra.Command, args []string) error {
		cName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err = c.SetDefaultContext(cName); err != nil {
			return err
		}
		if err = c.WriteConfig(); err != nil {
			return err
		}
		data := []output.Data{}
		for _, con := range c.Contexts {
			defa := false
			if con.Name == c.DefaultContext {
				defa = true
			}
			selected := ""
			if con.Name == cmdLagoon {
				selected = " (selected)"
			}
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v%s", con.Name, selected)),
				returnNonEmptyString(fmt.Sprintf("%v", defa)),
				returnNonEmptyString(fmt.Sprintf("%v", con.User)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.APIHostname)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenHost)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.TokenPort)),
				returnNonEmptyString(fmt.Sprintf("%v", con.ContextConfig.AuthenticationEndpoint)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"Name",
				"Default",
				"User",
				"API-Hostname",
				"Token-Hostname",
				"Token-Port",
				"Authentication-Hostname",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

func init() {
	config2Cmd.AddCommand(config2ListContextsCmd)
	config2Cmd.AddCommand(config2ListUsersCmd)
	config2Cmd.AddCommand(config2GetConfigPathCmd)

	config2Cmd.AddCommand(config2SetDefaultCmd)
	config2SetDefaultCmd.Flags().String("name", "", "The name of the context to be default")

	config2Cmd.AddCommand(config2AddContextCmd)
	config2AddContextCmd.Flags().String("name", "", "The name to reference this context as")
	config2AddContextCmd.Flags().String("api-hostname", "", "Lagoon API hostname (eg: https://api.lagoon.sh)")
	config2AddContextCmd.Flags().String("token-hostname", "", "Lagoon Token endpoint hostname (eg: token.lagoon.sh)")
	config2AddContextCmd.Flags().Int("token-port", 0, "Lagoon Token endpoint port (eg: 22)")
	config2AddContextCmd.Flags().String("authentication-hostname", "", "Lagoon authentication hostname (eg: https://keycloak.lagoon.sh)")
	config2AddContextCmd.Flags().String("user", "", "The user to associate to this context")

	config2Cmd.AddCommand(config2AddUserCmd)
	config2AddUserCmd.Flags().String("name", "", "The name to reference this user as")
	config2AddUserCmd.Flags().String("ssh-key", "", "The full path to this users ssh-key")

	config2Cmd.AddCommand(config2UpdateUserCmd)
	config2UpdateUserCmd.Flags().String("name", "", "The name to reference this user as")
	config2UpdateUserCmd.Flags().String("ssh-key", "", "The full path to this users ssh-key")

	config2Cmd.AddCommand(config2UpdateContextCmd)
	config2UpdateContextCmd.Flags().String("name", "", "The name to reference this context as")
	config2UpdateContextCmd.Flags().String("api-hostname", "", "Lagoon API hostname (eg: https://api.lagoon.sh)")
	config2UpdateContextCmd.Flags().String("token-hostname", "", "Lagoon Token endpoint hostname (eg: token.lagoon.sh)")
	config2UpdateContextCmd.Flags().Int("token-port", 0, "Lagoon Token endpoint port (eg: 22)")
	config2UpdateContextCmd.Flags().String("authentication-hostname", "", "Lagoon authentication hostname (eg: https://keycloak.lagoon.sh)")
	config2UpdateContextCmd.Flags().String("user", "", "The user to associate to this context")
}
