package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	l "github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	s "github.com/uselagoon/machinery/api/schema"

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
	Short:   "Add a new project to Lagoon, or add a project to an organization",
	Long:    "To add a project to an organization, you'll need to include the `organization` flag and provide the name of the organization. You need to be an owner of this organization to do this.\nIf you're the organization owner and want to grant yourself ownership to this project to be able to deploy environments, specify the `owner` flag.",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		gitUrl, err := cmd.Flags().GetString("git-url")
		if err != nil {
			return err
		}
		productionEnvironment, err := cmd.Flags().GetString("production-environment")
		if err != nil {
			return err
		}
		openshift, err := cmd.Flags().GetUint("openshift")
		if err != nil {
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
		orgOwner, err := cmd.Flags().GetBool("owner")
		if err != nil {
			return err
		}
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Project name", cmdProjectName, "GitURL", gitUrl, "Production environment", productionEnvironment, "Openshift", strconv.Itoa(int(openshift))); err != nil {
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

		projectInput := s.AddProjectInput{
			Name:                         cmdProjectName,
			AddOrgOwner:                  orgOwner,
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
		}

		if organizationName != "" {
			organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
			if err != nil {
				return err
			}
			projectInput.Organization = organization.ID
		}

		project := s.Project{}
		err = lc.AddProject(context.TODO(), &projectInput, &project)
		if err != nil {
			return err
		}
		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": project.Name,
				"GitURL":       gitUrl,
			},
		}
		if organizationName != "" {
			resultData.ResultData["Organization"] = organizationName
		}
		output.RenderResult(resultData, outputOptions)
		return nil
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
			return fmt.Errorf("missing arguments: key is not defined")
		}
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		projects, err := l.GetProjectsByMetadata(context.TODO(), key, value, lc)
		if err != nil {
			return err
		}
		if len(*projects) == 0 {
			if value != "" {
				outputOptions.Error = fmt.Sprintf("No projects found with metadata key '%s' and value '%s'\n", key, value)
			}
			outputOptions.Error = fmt.Sprintf("No projects found with metadata key '%s'\n", key)
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
			outputOptions.Error = fmt.Sprintf("There is no metadata for project '%s'\n", cmdProjectName)
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
			return fmt.Errorf("missing arguments: Project name or key is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to update key '%s' for project '%s' metadata, are you sure?", key, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
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
			return fmt.Errorf("missing arguments: Project name or key is not defined")
		}
		if yesNo(fmt.Sprintf("You are attempting to delete key '%s' from project '%s' metadata, are you sure?", key, cmdProjectName)) {
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
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

var removeProjectFromOrganizationCmd = &cobra.Command{
	Use:     "organization-project",
	Aliases: []string{"org-p"},
	Short:   "Remove a project from an Organization",
	Long:    "Removes a project from an Organization, but does not delete the project.\nThis is used by platform administrators to be able to reset a project.",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		organizationName, err := cmd.Flags().GetString("organization-name")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName, "Organization name", organizationName); err != nil {
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

		project, err := l.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		organization, err := l.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}

		projectInput := s.RemoveProjectFromOrganizationInput{
			Project:      project.ID,
			Organization: organization.ID,
		}

		if yesNo(fmt.Sprintf("You are attempting to remove project '%s' from organization '%s'. This will return the project to a state where it has no groups or notifications associated, are you sure?", cmdProjectName, organization.Name)) {
			_, err := l.RemoveProjectFromOrganization(context.TODO(), &projectInput, lc)
			if err != nil {
				return err
			}
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
	updateProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "gitUrl", "", "", "GitURL of the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.GitURL, "git-url", "g", "", "Git URL of the project")
	updateProjectCmd.Flags().MarkDeprecated("gitUrl", "please use --git-url instead")
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "privateKey", "", "", "Private key to use for the project")
	updateProjectCmd.Flags().StringVarP(&projectPatch.PrivateKey, "private-key", "I", "", "Private key to use for the project")
	updateProjectCmd.Flags().MarkDeprecated("privateKey", "please use --private-key instead")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Subfolder, "subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	updateProjectCmd.Flags().StringVarP(&projectPatch.RouterPattern, "routerPattern", "", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	updateProjectCmd.Flags().StringVarP(&projectPatch.RouterPattern, "router-pattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	updateProjectCmd.Flags().MarkDeprecated("routerPattern", "please use --router-pattern instead")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Branches, "branches", "b", "", "Which branches should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Name, "name", "N", "", "Change the name of the project by specifying a new name (careful!)")
	updateProjectCmd.Flags().StringVarP(&projectPatch.Pullrequests, "pullrequests", "m", "", "Which Pull Requests should be deployed")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "productionEnvironment", "", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().StringVarP(&projectPatch.ProductionEnvironment, "production-environment", "E", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().MarkDeprecated("productionEnvironment", "please use --production-environment instead")
	updateProjectCmd.Flags().StringVar(&projectPatch.StandbyProductionEnvironment, "standbyProductionEnvironment", "", "Which environment(the name) should be marked as the standby production environment")
	updateProjectCmd.Flags().StringVar(&projectPatch.StandbyProductionEnvironment, "standby-production-environment", "", "Which environment(the name) should be marked as the standby production environment")
	updateProjectCmd.Flags().MarkDeprecated("standbyProductionEnvironment", "please use --standby-production-environment instead")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshiftProjectPattern", "", "", "Pattern of OpenShift Project/Namespace that should be generated")
	updateProjectCmd.Flags().StringVarP(&projectPatch.OpenshiftProjectPattern, "openshift-project-pattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")
	updateProjectCmd.Flags().MarkDeprecated("openshiftProjectPattern", "please use --openshift-project-pattern instead")

	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "autoIdle", "", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().IntVarP(&projectAutoIdle, "auto-idle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().MarkDeprecated("autoIdle", "please use --auto-idle instead")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storageCalc", "", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().IntVarP(&projectStorageCalc, "storage-calc", "C", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().MarkDeprecated("storageCalc", "please use --storage-calc instead")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "developmentEnvironmentsLimit", "", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().IntVarP(&projectDevelopmentEnvironmentsLimit, "development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().MarkDeprecated("developmentEnvironmentsLimit", "please use --development-environments-limit instead")
	updateProjectCmd.Flags().IntVarP(&projectOpenshift, "openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")
	updateProjectCmd.Flags().IntVarP(&projectDeploymentsDisabled, "deploymentsDisabled", "", 0, "Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable")
	updateProjectCmd.Flags().IntVarP(&projectDeploymentsDisabled, "deployments-disabled", "", 0, "Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable")
	updateProjectCmd.Flags().MarkDeprecated("deploymentsDisabled", "please use --deployments-disabled instead")

	updateProjectCmd.Flags().IntVarP(&factsUi, "factsUi", "", 0, "Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().IntVarP(&factsUi, "facts-ui", "", 0, "Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().MarkDeprecated("factsUi", "please use --facts-ui instead")
	updateProjectCmd.Flags().IntVarP(&problemsUi, "problemsUi", "", 0, "Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().IntVarP(&problemsUi, "problems-ui", "", 0, "Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().MarkDeprecated("problemsUi", "please use --problems-ui instead")

	addProjectCmd.Flags().StringP("json", "j", "", "JSON string to patch")

	addProjectCmd.Flags().StringP("gitUrl", "", "", "GitURL of the project")
	addProjectCmd.Flags().StringP("git-url", "g", "", "GitURL of the project")
	addProjectCmd.Flags().MarkDeprecated("gitUrl", "please use --git-url instead")
	addProjectCmd.Flags().StringP("privateKey", "", "", "Private key to use for the project")
	addProjectCmd.Flags().StringP("private-key", "I", "", "Private key to use for the project")
	addProjectCmd.Flags().MarkDeprecated("privateKey", "please use --private-key instead")
	addProjectCmd.Flags().StringP("subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	addProjectCmd.Flags().StringP("routerPattern", "", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	addProjectCmd.Flags().StringP("router-pattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	addProjectCmd.Flags().MarkDeprecated("routerPattern", "please use --router-pattern instead")
	addProjectCmd.Flags().StringP("branches", "b", "", "Which branches should be deployed")
	addProjectCmd.Flags().StringP("pullrequests", "m", "", "Which Pull Requests should be deployed")
	addProjectCmd.Flags().StringP("productionEnvironment", "", "", "Which environment(the name) should be marked as the production environment")
	addProjectCmd.Flags().StringP("production-environment", "E", "", "Which environment(the name) should be marked as the production environment")
	addProjectCmd.Flags().MarkDeprecated("productionEnvironment", "please use --production-environment instead")
	addProjectCmd.Flags().String("standbyProductionEnvironment", "", "Which environment(the name) should be marked as the standby production environment")
	addProjectCmd.Flags().String("standby-production-environment", "", "Which environment(the name) should be marked as the standby production environment")
	addProjectCmd.Flags().MarkDeprecated("standbyProductionEnvironment", "please use --standby-production-environment instead")
	addProjectCmd.Flags().StringP("openshiftProjectPattern", "", "", "Pattern of OpenShift Project/Namespace that should be generated")
	addProjectCmd.Flags().StringP("openshift-project-pattern", "o", "", "Pattern of OpenShift Project/Namespace that should be generated")
	addProjectCmd.Flags().MarkDeprecated("openshiftProjectPattern", "please use --openshift-project-pattern instead")

	addProjectCmd.Flags().UintP("autoIdle", "", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().UintP("auto-idle", "a", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().MarkDeprecated("autoIdle", "please use --auto-idle instead")
	addProjectCmd.Flags().UintP("storageCalc", "", 0, "Should storage for this environment be calculated")
	addProjectCmd.Flags().UintP("storage-calc", "C", 0, "Should storage for this environment be calculated")
	addProjectCmd.Flags().MarkDeprecated("storageCalc", "please use --storage-calc instead")
	addProjectCmd.Flags().UintP("developmentEnvironmentsLimit", "", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().UintP("development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().MarkDeprecated("developmentEnvironmentsLimit", "please use --development-environments-limit instead")
	addProjectCmd.Flags().UintP("openshift", "S", 0, "Reference to OpenShift Object this Project should be deployed to")
	addProjectCmd.Flags().Bool("owner", false, "Add the user as an owner of the project")
	addProjectCmd.Flags().StringP("organization-name", "O", "", "Name of the Organization to add the project to")

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

	removeProjectFromOrganizationCmd.Flags().StringP("organization-name", "O", "", "Name of the Organization to remove the project from")
}
