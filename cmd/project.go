package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/lagoon/projects"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
		var projectName string
		if len(args) < 1 {
			if cmdProject.Name != "" {
				projectName = cmdProject.Name
			} else {
				fmt.Println("Not enough arguments. Requires: project name")
				cmd.Help()
				os.Exit(1)
			}
		} else {
			projectName = args[0]
		}

		if !outputOptions.JSON {
			fmt.Println(fmt.Sprintf("Deleting %s", projectName))
		}

		if yesNo() {
			deleteResult, err := projects.DeleteProject(projectName)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
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
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}

		jsonPatch, _ := json.Marshal(projectFlags)
		addResult, err := projects.AddProject(cmdProjectName, string(jsonPatch))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var addedProject api.Project
		err = json.Unmarshal([]byte(addResult), &addedProject)

		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}

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
			fmt.Println("Not enough arguments. Requires: project name")
			cmd.Help()
			os.Exit(1)
		}

		jsonPatch, _ := json.Marshal(projectFlags)
		projectUpdateID, err := projects.UpdateProject(cmdProjectName, string(jsonPatch))
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		var updatedProject api.Project
		err = json.Unmarshal([]byte(projectUpdateID), &updatedProject)
		if err != nil {
			output.RenderError(err.Error(), outputOptions)
			os.Exit(1)
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": updatedProject.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)

	},
}

func init() {
	deleteCmd.AddCommand(deleteProjectCmd)

	updateCmd.AddCommand(updateProjectCmd)
	updateProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	// @TODO this seems needlessly busy, maybe see if cobra supports grouping flags and applying them to commands easier?
	updateProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder Usefull if you have multiple Lagoon projects per Git Repository")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsTask, "activeSystemsTask", "T", "", "Which internal Lagoon System is responsible for tasks ")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsDeploy, "activeSystemsDeploy", "D", "", "Which internal Lagoon System is responsible for deploying ")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsRemove, "activeSystemsRemove", "R", "", "Which internal Lagoon System is responsible for promoting")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsPromote, "activeSystemsPromote", "P", "", "Which internal Lagoon System is responsible for promoting")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")

	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")

	updateProjectCmd.Flags().MarkHidden("activeSystemsTask")
	updateProjectCmd.Flags().MarkHidden("activeSystemsDeploy")
	updateProjectCmd.Flags().MarkHidden("activeSystemsRemove")
	updateProjectCmd.Flags().MarkHidden("activeSystemsPromote")
	updateProjectCmd.Flags().MarkHidden("openshiftProjectPattern")
	updateProjectCmd.Flags().MarkHidden("developmentEnvironmentsLimit")
	updateProjectCmd.Flags().MarkHidden("openshift")

	addCmd.AddCommand(addProjectCmd)
	addProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	addProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder Usefull if you have multiple Lagoon projects per Git Repository")
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

	addProjectCmd.Flags().MarkHidden("activeSystemsTask")
	addProjectCmd.Flags().MarkHidden("activeSystemsDeploy")
	addProjectCmd.Flags().MarkHidden("activeSystemsRemove")
	addProjectCmd.Flags().MarkHidden("activeSystemsPromote")
	addProjectCmd.Flags().MarkHidden("openshiftProjectPattern")
	addProjectCmd.Flags().MarkHidden("developmentEnvironmentsLimit")
	addProjectCmd.Flags().MarkHidden("openshift")
}
