package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/pkg/api"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var projectPatch api.ProjectPatch

var projectAutoIdle int
var projectStorageCalc int
var projectDevelopmentEnvironmentsLimit int
var projectOpenshift int
var projectDeploymentsDisabled int
var factsUi int
var problemsUi int

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
	Short:   "Add a new project to Lagoon",
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

		var jsonPatchFromProjectFlags string
		if string(jsonPatch) != "" {
			jsonPatchFromProjectFlags = jsonPatch
		} else {
			jsonMarshalPatch, _ := json.Marshal(projectFlags)
			jsonPatchFromProjectFlags = string(jsonMarshalPatch)
		}

		projectUpdateID, err := pClient.UpdateProject(cmdProjectName, jsonPatchFromProjectFlags)
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

var listProjectByMetadata = &cobra.Command{
	Use:     "projects-by-metadata",
	Aliases: []string{"pm", "projectmeta"},
	Short:   "List projects by a given metadata key or key:value",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		showMetadata, err := cmd.Flags().GetBool("show-metadata")
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
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)
		projects, err := l.GetProjectsByMetadata(context.TODO(), key, value, lc)
		if err != nil {
			return err
		}
		if len(*projects) == 0 {
			if value != "" {
				return fmt.Errorf(fmt.Sprintf("No projects found with metadata key %s and value %s", key, value))
			}
			return fmt.Errorf(fmt.Sprintf("No projects found with metadata key %s", key))
		}
		data := []output.Data{}
		for _, project := range *projects {
			projectData := []string{
				returnNonEmptyString(fmt.Sprintf("%v", project.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", project.Name)),
			}
			if showMetadata {
				metaData, _ := json.Marshal(project.Metadata)
				projectData = append(projectData, returnNonEmptyString(fmt.Sprintf("%v", string(metaData))))
			}
			data = append(data, projectData)
		}
		header := []string{
			"ID",
			"Name",
		}
		if showMetadata {
			header = append(header, "Metadata")
		}
		output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		return nil
	},
}

var getProjectMetadata = &cobra.Command{
	Use:     "project-metadata",
	Aliases: []string{"pm", "projectmeta"},
	Short:   "Get all metadata for a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name is not defined")
			cmd.Help()
			os.Exit(1)
		}
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)
		project, err := lagoon.GetProjectMetadata(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(project.Metadata) == 0 {
			output.RenderInfo(fmt.Sprintf("There is no metadata for project '%s'", cmdProjectName), outputOptions)
			return nil
		}
		data := []output.Data{}
		for metaKey, metaVal := range project.Metadata {
			metadataData := []string{
				returnNonEmptyString(fmt.Sprintf("%v", metaKey)),
				returnNonEmptyString(fmt.Sprintf("%v", metaVal)),
			}
			data = append(data, metadataData)
		}
		header := []string{
			"Key",
			"Value",
		}
		output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		return nil
	},
}

var updateProjectMetadata = &cobra.Command{
	Use:     "project-metadata",
	Aliases: []string{"pm", "meta", "projectmeta"},
	Short:   "Update a projects metadata with a given key or key:value",
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
		if key == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: Project name or key is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to update key '%s' for project '%s' metadata, are you sure?", key, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				&token,
				debug)
			project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			projectResult, err := l.UpdateProjectMetadata(context.TODO(), int(project.ID), key, value, lc)
			if err != nil {
				return err
			}
			data := []output.Data{}
			metaData, _ := json.Marshal(projectResult.Metadata)
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", string(metaData))),
			})
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Metadata",
				},
				Data: data,
			}, outputOptions)
		}
		return nil
	},
}

var deleteProjectMetadataByKey = &cobra.Command{
	Use:     "project-metadata",
	Aliases: []string{"pm", "meta", "projectmeta"},
	Short:   "Delete a key from a projects metadata",
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
		if key == "" || cmdProjectName == "" {
			return fmt.Errorf("Missing arguments: Project name or key is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to delete key '%s' from project '%s' metadata, are you sure?", key, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				&token,
				debug)
			project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			projectResult, err := l.RemoveProjectMetadataByKey(context.TODO(), int(project.ID), key, lc)
			if err != nil {
				return err
			}
			data := []output.Data{}
			metaData, _ := json.Marshal(projectResult.Metadata)
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", string(metaData))),
			})
			output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Metadata",
				},
				Data: data,
			}, outputOptions)
		}
		return nil
	},
}

var addProjectToOrganizationCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Add a project to an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		handleError(err)

		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
			return err
		}
		gitUrl, err := cmd.Flags().GetString("git-url")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("gitUrl", gitUrl); err != nil {
			return err
		}
		productionEnvironment, err := cmd.Flags().GetString("production-environment")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Production Environment", productionEnvironment); err != nil {
			return err
		}
		openshift, err := cmd.Flags().GetUint("openshift")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("openshift", strconv.Itoa(int(openshift))); err != nil {
			return err
		}
		standbyProductionEnvironment, err := cmd.Flags().GetString("standby-production-environment")
		if err != nil {
			return err
		}
		branches, err := cmd.Flags().GetString("branches")
		if err != nil {
			return err
		}
		pullrequests, err := cmd.Flags().GetString("pullrequests")
		if err != nil {
			return err
		}
		openshiftProjectPattern, err := cmd.Flags().GetString("openshift-project-pattern")
		if err != nil {
			return err
		}
		developmentEnvironmentsLimit, err := cmd.Flags().GetUint("development-environments-limit")
		if err != nil {
			return err
		}
		storageCalc, err := cmd.Flags().GetUint("storage-calc")
		if err != nil {
			return err
		}
		autoIdle, err := cmd.Flags().GetUint("auto-idle")
		if err != nil {
			return err
		}
		subfolder, err := cmd.Flags().GetString("subfolder")
		if err != nil {
			return err
		}
		privateKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return err
		}
		orgOwner, err := cmd.Flags().GetBool("org-owner")
		if err != nil {
			return err
		}
		buildImage, err := cmd.Flags().GetString("build-image")
		if err != nil {
			return err
		}
		availability, err := cmd.Flags().GetString("availability")
		if err != nil {
			return err
		}
		factsUi, err := cmd.Flags().GetUint("facts-ui")
		if err != nil {
			return err
		}
		problemsUi, err := cmd.Flags().GetUint("problems-ui")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		deploymentsDisabled, err := cmd.Flags().GetUint("deployments-disabled")
		if err != nil {
			return err
		}
		ProductionBuildPriority, err := cmd.Flags().GetUint("production-build-priority")
		if err != nil {
			return err
		}
		DevelopmentBuildPriority, err := cmd.Flags().GetUint("development-build-priority")
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		handleError(err)

		projectInput := s.AddProjectInput{
			Name:                         cmdProjectName,
			Organization:                 organization.ID,
			AddOrgOwner:                  orgOwner,
			BuildImage:                   buildImage,
			Availability:                 s.ProjectAvailability(availability),
			GitURL:                       gitUrl,
			ProductionEnvironment:        productionEnvironment,
			StandbyProductionEnvironment: standbyProductionEnvironment,
			Branches:                     branches,
			PullRequests:                 pullrequests,
			OpenshiftProjectPattern:      openshiftProjectPattern,
			Openshift:                    openshift,
			DevelopmentEnvironmentsLimit: developmentEnvironmentsLimit,
			StorageCalc:                  storageCalc,
			AutoIdle:                     autoIdle,
			Subfolder:                    subfolder,
			PrivateKey:                   privateKey,
			RouterPattern:                routerPattern,
			ProblemsUI:                   problemsUi,
			FactsUI:                      factsUi,
			ProductionBuildPriority:      ProductionBuildPriority,
			DevelopmentBuildPriority:     DevelopmentBuildPriority,
			DeploymentsDisabled:          deploymentsDisabled,
		}
		project := s.Project{}
		err = lc.AddProject(context.TODO(), &projectInput, &project)
		handleError(err)

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name":      project.Name,
				"Organization Name": organizationName,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
	},
}

var RemoveProjectFromOrganizationCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Remove a project from an Organization",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		handleError(err)

		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}
		organizationName, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Organization name", organizationName); err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			&token,
			debug)

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		handleError(err)
		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		handleError(err)

		projectInput := s.RemoveProjectFromOrganizationInput{
			Project:      project.ID,
			Organization: organization.ID,
		}

		if yesNo(fmt.Sprintf("You are attempting to remove project '%s' from organization '%s'. This will return the project to a state where it has no groups or notifications associated, are you sure?", cmdProjectName, organization.Name)) {
			_, err := l.RemoveProjectFromOrganization(context.TODO(), &projectInput, lc)
			handleError(err)
			resultData := output.Result{
				Result: "success",
				ResultData: map[string]interface{}{
					"Project Name":      cmdProjectName,
					"Organization Name": organizationName,
				},
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

func init() {
	updateProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	// @TODO this seems needlessly busy, maybe see if cobra supports grouping flags and applying them to commands easier?
	updateProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	updateProjectCmd.Flags().StringVarP(&projectPatch.RouterPattern, "routerPattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Name, "name", "N", "", "Change the name of the project by specifying a new name (careful!)")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().StringVar(&projectPatch.StandbyProductionEnvironment, "standbyProductionEnvironment", "", "Which environment(the name) should be marked as the standby production environment")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")

	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")
	updateProjectCmd.Flags().IntVarP(&projectDeploymentsDisabled, "deploymentsDisabled", "", 0, "Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable")

	updateProjectCmd.Flags().IntVarP(&factsUi, "factsUi", "", 0, "Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().IntVarP(&problemsUi, "problemsUi", "", 0, "Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable")

	addProjectCmd.Flags().StringVarP(&jsonPatch, "json", "j", "", "JSON string to patch")

	addProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "g", "", "GitURL of the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "I", "", "Private key to use for the project")
	addProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	addProjectCmd.Flags().StringVarP(&projectPatch.RouterPattern, "routerPattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	addProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	addProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	addProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "E", "", "Which environment(the name) should be marked as the production environment")
	addProjectCmd.Flags().StringVar(&projectPatch.StandbyProductionEnvironment, "standbyProductionEnvironment", "", "Which environment(the name) should be marked as the standby production environment")
	addProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")

	addProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "a", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "C", 0, "Should storage for this environment be calculated")
	addProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "L", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")

	listCmd.AddCommand(listProjectByMetadata)
	listProjectByMetadata.Flags().StringP("key", "K", "", "The key name of the metadata value you are querying on")
	listProjectByMetadata.Flags().StringP("value", "V", "", "The value for the key you are querying on")
	listProjectByMetadata.Flags().Bool("show-metadata", false, "Show the metadata for each project as another field (this could be a lot of data)")

	updateCmd.AddCommand(updateProjectMetadata)
	updateProjectMetadata.Flags().StringP("key", "K", "", "The key name of the metadata value you are querying on")
	updateProjectMetadata.Flags().StringP("value", "V", "", "The value for the key you are querying on")

	deleteCmd.AddCommand(deleteProjectMetadataByKey)
	deleteProjectMetadataByKey.Flags().StringP("key", "K", "", "The key name of the metadata value you are querying on")

	getCmd.AddCommand(getProjectMetadata)

	addProjectToOrganizationCmd.Flags().String("build-image", "", "Build Image for the project")
	addProjectToOrganizationCmd.Flags().String("availability", "", "Availability of the project")
	addProjectToOrganizationCmd.Flags().String("git-url", "", "GitURL of the project")
	addProjectToOrganizationCmd.Flags().String("production-environment", "", "Production Environment for the project")
	addProjectToOrganizationCmd.Flags().String("standby-production-environment", "", "Standby Production Environment for the project")
	addProjectToOrganizationCmd.Flags().String("subfolder", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	addProjectToOrganizationCmd.Flags().String("private-key", "", "Private key to use for the project")
	addProjectToOrganizationCmd.Flags().String("branches", "", "branches")
	addProjectToOrganizationCmd.Flags().String("pullrequests", "", "Which Pull Requests should be deployed")
	addProjectToOrganizationCmd.Flags().StringP("name", "O", "", "Name of the Organization to add the project to")
	addProjectToOrganizationCmd.Flags().String("openshift-project-pattern", "", "Pattern of OpenShift Project/Namespace that should be generated")
	addProjectToOrganizationCmd.Flags().String("router-pattern", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")

	addProjectToOrganizationCmd.Flags().Uint("openshift", 0, "Reference to OpenShift Object this Project should be deployed to")
	addProjectToOrganizationCmd.Flags().Uint("auto-idle", 0, "Auto idle setting of the project")
	addProjectToOrganizationCmd.Flags().Uint("storage-calc", 0, "Should storage for this environment be calculated")
	addProjectToOrganizationCmd.Flags().Uint("development-environments-limit", 0, "How many environments can be deployed at one time")
	addProjectToOrganizationCmd.Flags().Uint("facts-ui", 0, "Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable")
	addProjectToOrganizationCmd.Flags().Uint("problems-ui", 0, "Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable")
	addProjectToOrganizationCmd.Flags().Uint("deployments-disabled", 0, "Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable")
	addProjectToOrganizationCmd.Flags().Uint("production-build-priority", 0, "Set the priority of the production build")
	addProjectToOrganizationCmd.Flags().Uint("development-build-priority", 0, "Set the priority of the development build")

	addProjectToOrganizationCmd.Flags().Bool("org-owner", false, "Add the user as an owner of the project")

	RemoveProjectFromOrganizationCmd.Flags().StringP("name", "O", "", "Name of the Organization to remove the project from")
}
