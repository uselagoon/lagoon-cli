package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/amazeeio/lagoon-cli/pkg/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var projectPatch api.ProjectPatch

var projectAutoIdle int
var projectStorageCalc int
var projectDevelopmentEnvironmentsLimit int
var projectOpenshift int

func parseProjectFlags(flags pflag.FlagSet) api.ProjectPatch {
	configMap := make(map[string]interface{})
	flags.VisitAll(func(f *pflag.Flag) {
		if flags.Changed(f.Name) {
			configMap[f.Name] = &f.Value
		}
	})
	jsonStr, _ := json.Marshal(configMap)
	parsedFlags := api.ProjectPatch{}
	json.Unmarshal(jsonStr, &parsedFlags)
	return parsedFlags
}

var deleteProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Delete a project",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		if yesNo(fmt.Sprintf("You are attempting to delete project '%s', are you sure?", cmdProjectName)) {
			deleteResult, err := pClient.DeleteProject(cmdProjectName)
			handleError(err)
			resultData := output.Result{
				Result: string(deleteResult),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var addProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Add a new project to lagoon",
	Run: func(cmd *cobra.Command, args []string) {
		projectFlags := parseProjectFlags(*cmd.Flags())
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		jsonPatch, _ := json.Marshal(projectFlags)
		addResult, err := pClient.AddProject(cmdProjectName, string(jsonPatch))
		handleError(err)
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)
		handleError(err)

		if err != nil {
			output.RenderError(err.Error(), outputOptions)
		} else {
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"Project Name": addedProject.Name,
					"GitURL":       addedProject.GitURL,
				},
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

var updateProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Update a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectFlags := parseProjectFlags(*cmd.Flags())
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}

		jsonPatch, _ := json.Marshal(projectFlags)
		projectUpdateID, err := pClient.UpdateProject(cmdProjectName, string(jsonPatch))
		handleError(err)
		var updatedProject api.Project
		err = json.Unmarshal([]byte(projectUpdateID), &updatedProject)
		handleError(err)
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": updatedProject.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)
	},
}

var getProjectByMetadata = &cobra.Command{
	Use:     "project-by-metadata",
	Aliases: []string{"pm", "projectmeta"},
	Short:   "Checks the current Lagoon for its version and sets it in the config file",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			return err
		}
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			return err
		}
		if key == "" {
			return fmt.Errorf("Missing arguments: key is not defined")
		}
		current := viper.GetString("current")
		lc := client.New(
			viper.GetString("lagoons."+current+".graphql"),
			viper.GetString("lagoons."+current+".token"),
			viper.GetString("lagoons."+current+".version"),
			lagoonCLIVersion,
			debug)
		projects, err := lagoon.GetProjectsByMetadata(context.TODO(), key, value, lc)
		if err != nil {
			return err
		}
		data := []output.Data{}
		for _, project := range *projects {
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
			})
		}
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
			},
			Data: data,
		}, outputOptions)
		return nil
	},
}

func init() {
	updateProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	// @TODO this seems needlessly busy, maybe see if cobra supports grouping flags and applying them to commands easier?
	updateProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsTask, "activeSystemsTask", "T", "", "Which internal Lagoon System is responsible for tasks ")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsDeploy, "activeSystemsDeploy", "D", "", "Which internal Lagoon System is responsible for deploying ")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsRemove, "activeSystemsRemove", "R", "", "Which internal Lagoon System is responsible for promoting")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsPromote, "activeSystemsPromote", "P", "", "Which internal Lagoon System is responsible for promoting")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Name, "name", "N", "", "Change the name of the project by specifying a new name (careful!)")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")

	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")

	addProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	addProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsTask, "activeSystemsTask", "T", "", "Which internal Lagoon System is responsible for tasks ")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsDeploy, "activeSystemsDeploy", "D", "", "Which internal Lagoon System is responsible for deploying ")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsRemove, "activeSystemsRemove", "R", "", "Which internal Lagoon System is responsible for promoting")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsPromote, "activeSystemsPromote", "P", "", "Which internal Lagoon System is responsible for promoting")
	addProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	addProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	addProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Which environment(the name) should be marked as the production environment")
	addProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")

	addProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Should storage for this environment be calculated")
	addProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")

	getCmd.AddCommand(getProjectByMetadata)
	getProjectByMetadata.Flags().String("key", "", "The key name of the metadata value you are querying on")
	getProjectByMetadata.Flags().String("value", "", "The value for the key you are querying on")
}
