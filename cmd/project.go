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
	Use:   "project",
	Short: "Delete a project",
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
	Use:   "project",
	Short: "Add a new project to lagoon",
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
	Use:   "project [project name]",
	Short: "Update a project",
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
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Branches of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Production environment of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsTask, "activeSystemsTask", "T", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsDeploy, "activeSystemsDeploy", "D", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsRemove, "activeSystemsRemove", "R", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsPromote, "activeSystemsPromote", "P", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Pull requests of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pull requests of the project")

	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Auto idle setting of the project")

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
	addProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Branches of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Production environment of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsTask, "activeSystemsTask", "T", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsDeploy, "activeSystemsDeploy", "D", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsRemove, "activeSystemsRemove", "R", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.ActiveSystemsPromote, "activeSystemsPromote", "P", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Pull requests of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pull requests of the project")

	addProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Auto idle setting of the project")

	addProjectCmd.Flags().MarkHidden("activeSystemsTask")
	addProjectCmd.Flags().MarkHidden("activeSystemsDeploy")
	addProjectCmd.Flags().MarkHidden("activeSystemsRemove")
	addProjectCmd.Flags().MarkHidden("activeSystemsPromote")
	addProjectCmd.Flags().MarkHidden("openshiftProjectPattern")
	addProjectCmd.Flags().MarkHidden("developmentEnvironmentsLimit")
	addProjectCmd.Flags().MarkHidden("openshift")
}
