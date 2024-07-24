package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"strings"

	"github.com/uselagoon/machinery/api/lagoon"
	lclient "github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"

	"github.com/guregu/null"
	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

var deleteProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Delete a project",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete project '%s', are you sure?", cmdProjectName)) {
			_, err := lagoon.DeleteProject(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			output.RenderResult(resultData, outputOptions)
		}
		return nil
	},
}

var addProjectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"p"},
	Short:   "Add a new project to Lagoon, or add a project to an organization",
	Long:    "To add a project to an organization, you'll need to include the `organization` flag and provide the name of the organization. You need to be an owner of this organization to do this.\nIf you're the organization owner and want to grant yourself ownership to this project to be able to deploy environments, specify the `owner` flag.",
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lContext.Name)
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
		organizationID, err := cmd.Flags().GetUint("organization-id")
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
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
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
		deploytargetProjectPattern, err := cmd.Flags().GetString("deploytarget-project-pattern")
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
		buildImage, err := cmd.Flags().GetString("build-image")
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

		if err := requiredInputCheck("Project name", cmdProjectName, "git-url", gitUrl, "Production environment", productionEnvironment, "Deploytarget", strconv.Itoa(int(deploytarget))); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		projectInput := schema.AddProjectInput{
			Name:                         cmdProjectName,
			AddOrgOwner:                  orgOwner,
			GitURL:                       gitUrl,
			ProductionEnvironment:        productionEnvironment,
			StandbyProductionEnvironment: standbyProductionEnvironment,
			Branches:                     branches,
			PullRequests:                 pullrequests,
			OpenshiftProjectPattern:      deploytargetProjectPattern,
			Openshift:                    deploytarget,
			DevelopmentEnvironmentsLimit: developmentEnvironmentsLimit,
			StorageCalc:                  storageCalc,
			AutoIdle:                     autoIdle,
			Subfolder:                    subfolder,
			PrivateKey:                   privateKey,
			BuildImage:                   buildImage,
			RouterPattern:                routerPattern,
		}
		// if organizationid is provided, use it over the name
		if organizationID != 0 {
			projectInput.Organization = organizationID
		}
		// otherwise if name is provided use it
		if organizationName != "" && organizationID == 0 {
			organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
			if err != nil {
				return err
			}
			// since getorganizationbyname returns null response if an organization doesn't exist
			// check if the result has a name
			if organization.Name == "" {
				return fmt.Errorf("error querying organization by name")
			}
			projectInput.Organization = organization.ID
		}

		project := schema.Project{}
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
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}
		projectName, err := cmd.Flags().GetString("name")
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
		deploytarget, err := cmd.Flags().GetUint("deploytarget")
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
		deploytargetProjectPattern, err := cmd.Flags().GetString("deploytarget-project-pattern")
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
		autoIdleProvided := cmd.Flags().Lookup("auto-idle").Changed
		subfolder, err := cmd.Flags().GetString("subfolder")
		if err != nil {
			return err
		}
		privateKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return err
		}
		buildImage, err := cmd.Flags().GetString("build-image")
		if err != nil {
			return err
		}
		buildImageProvided := cmd.Flags().Lookup("build-image").Changed
		availability, err := cmd.Flags().GetString("availability")
		if err != nil {
			return err
		}
		factsUi, err := cmd.Flags().GetUint("facts-ui")
		if err != nil {
			return err
		}
		factsUIProvided := cmd.Flags().Lookup("facts-ui").Changed
		problemsUi, err := cmd.Flags().GetUint("problems-ui")
		if err != nil {
			return err
		}
		problemsUIProvided := cmd.Flags().Lookup("problems-ui").Changed
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		deploymentsDisabled, err := cmd.Flags().GetUint("deployments-disabled")
		if err != nil {
			return err
		}
		deploymentsDisabledProvided := cmd.Flags().Lookup("deployments-disabled").Changed
		ProductionBuildPriority, err := cmd.Flags().GetUint("production-build-priority")
		if err != nil {
			return err
		}
		DevelopmentBuildPriority, err := cmd.Flags().GetUint("development-build-priority")
		if err != nil {
			return err
		}

		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		projectPatch := schema.UpdateProjectPatchInput{
			GitURL:                       nullStrCheck(gitUrl),
			ProductionEnvironment:        nullStrCheck(productionEnvironment),
			Openshift:                    nullUintCheck(deploytarget),
			StandbyProductionEnvironment: nullStrCheck(standbyProductionEnvironment),
			Branches:                     nullStrCheck(branches),
			Pullrequests:                 nullStrCheck(pullrequests),
			OpenshiftProjectPattern:      nullStrCheck(deploytargetProjectPattern),
			DevelopmentEnvironmentsLimit: nullUintCheck(developmentEnvironmentsLimit),
			StorageCalc:                  nullUintCheck(storageCalc),
			AutoIdle:                     nullUintCheck(autoIdle),
			Subfolder:                    nullStrCheck(subfolder),
			PrivateKey:                   nullStrCheck(privateKey),
			FactsUI:                      nullUintCheck(factsUi),
			ProblemsUI:                   nullUintCheck(problemsUi),
			RouterPattern:                nullStrCheck(routerPattern),
			DeploymentsDisabled:          nullUintCheck(deploymentsDisabled),
			ProductionBuildPriority:      nullUintCheck(ProductionBuildPriority),
			DevelopmentBuildPriority:     nullUintCheck(DevelopmentBuildPriority),
			Name:                         nullStrCheck(projectName),
		}

		if availability != "" {
			ProjectAvailability := schema.ProjectAvailability(strings.ToUpper(availability))
			projectPatch.Availability = &ProjectAvailability
		}
		if autoIdleProvided {
			projectPatch.AutoIdle = &autoIdle
		}
		if deploymentsDisabledProvided {
			projectPatch.DeploymentsDisabled = &deploymentsDisabled
		}
		if factsUIProvided {
			projectPatch.FactsUI = &factsUi
		}
		if problemsUIProvided {
			projectPatch.ProblemsUI = &problemsUi
		}
		if buildImageProvided {
			if buildImage == "null" {
				nullBuildImage := null.String{}
				projectPatch.BuildImage = &nullBuildImage
			} else {
				buildImg := null.StringFrom(buildImage)
				projectPatch.BuildImage = &buildImg
			}
		}

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			outputOptions.Error = fmt.Sprintf("Project '%s' not found\n", cmdProjectName)
			output.RenderError(outputOptions.Error, outputOptions)
			return nil
		}
		projectUpdate, err := lagoon.UpdateProject(context.TODO(), int(project.ID), projectPatch, lc)
		if err != nil {
			return err
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": projectUpdate.Name,
			},
		}
		output.RenderResult(resultData, outputOptions)
		return nil
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
		if err := requiredInputCheck("Key", key); err != nil {
			return err
		}
		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)
		projects, err := lagoon.GetProjectsByMetadata(context.TODO(), key, value, lc)
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
		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
			return err
		}
		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Key", key); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to update key '%s' for project '%s' metadata, are you sure?", key, cmdProjectName)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			projectResult, err := lagoon.UpdateProjectMetadata(context.TODO(), int(project.ID), key, value, lc)
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
		if err := requiredInputCheck("Project name", cmdProjectName, "Key", key); err != nil {
			return err
		}
		if yesNo(fmt.Sprintf("You are attempting to delete key '%s' from project '%s' metadata, are you sure?", key, cmdProjectName)) {
			utoken := lUser.UserConfig.Grant.AccessToken
			lc := lclient.New(
				fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
				lagoonCLIVersion,
				lContext.ContextConfig.Version,
				&utoken,
				debug)
			project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			projectResult, err := lagoon.RemoveProjectMetadataByKey(context.TODO(), int(project.ID), key, lc)
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
		return validateTokenE(lContext.Name)
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

		utoken := lUser.UserConfig.Grant.AccessToken
		lc := lclient.New(
			fmt.Sprintf("%s/graphql", lContext.ContextConfig.APIHostname),
			lagoonCLIVersion,
			lContext.ContextConfig.Version,
			&utoken,
			debug)

		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		organization, err := lagoon.GetOrganizationByName(context.TODO(), organizationName, lc)
		if err != nil {
			return err
		}
		if organization.Name == "" {
			return fmt.Errorf("error querying organization by name")
		}

		projectInput := schema.RemoveProjectFromOrganizationInput{
			Project:      project.ID,
			Organization: organization.ID,
		}

		if yesNo(fmt.Sprintf("You are attempting to remove project '%s' from organization '%s'. This will return the project to a state where it has no groups or notifications associated, are you sure?", cmdProjectName, organization.Name)) {
			_, err := lagoon.RemoveProjectFromOrganization(context.TODO(), &projectInput, lc)
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
	updateProjectCmd.Flags().StringP("git-url", "g", "", "GitURL of the project")
	updateProjectCmd.Flags().StringP("private-key", "I", "", "Private key to use for the project")
	updateProjectCmd.Flags().StringP("subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	updateProjectCmd.Flags().StringP("router-pattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	updateProjectCmd.Flags().StringP("branches", "b", "", "Which branches should be deployed")
	updateProjectCmd.Flags().StringP("name", "N", "", "Change the name of the project by specifying a new name (careful!)")
	updateProjectCmd.Flags().StringP("pullrequests", "m", "", "Which Pull Requests should be deployed")
	updateProjectCmd.Flags().StringP("production-environment", "E", "", "Which environment(the name) should be marked as the production environment")
	updateProjectCmd.Flags().String("standby-production-environment", "", "Which environment(the name) should be marked as the standby production environment")
	updateProjectCmd.Flags().StringP("deploytarget-project-pattern", "o", "", "Pattern of Deploytarget(Kubernetes) Project/Namespace that should be generated")
	updateProjectCmd.Flags().StringP("build-image", "", "", "Build Image for the project. Set to 'null' to remove the build image")
	updateProjectCmd.Flags().StringP("availability", "", "", "Availability of the project")

	updateProjectCmd.Flags().Uint("production-build-priority", 0, "Set the priority of the production build")
	updateProjectCmd.Flags().Uint("development-build-priority", 0, "Set the priority of the development build")
	updateProjectCmd.Flags().UintP("auto-idle", "a", 0, "Auto idle setting of the project")
	updateProjectCmd.Flags().UintP("storage-calc", "C", 0, "Should storage for this environment be calculated")
	updateProjectCmd.Flags().UintP("development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().UintP("deploytarget", "S", 0, "Reference to Deploytarget(Kubernetes) this Project should be deployed to")
	updateProjectCmd.Flags().UintP("deployments-disabled", "", 0, "Admin only flag for disabling deployments on a project, 1 to disable deployments, 0 to enable")

	updateProjectCmd.Flags().UintP("facts-ui", "", 0, "Enables the Lagoon insights Facts tab in the UI. Set to 1 to enable, 0 to disable")
	updateProjectCmd.Flags().UintP("problems-ui", "", 0, "Enables the Lagoon insights Problems tab in the UI. Set to 1 to enable, 0 to disable")

	addProjectCmd.Flags().StringP("json", "j", "", "JSON string to patch")

	addProjectCmd.Flags().StringP("git-url", "g", "", "GitURL of the project")
	addProjectCmd.Flags().StringP("private-key", "I", "", "Private key to use for the project")
	addProjectCmd.Flags().StringP("subfolder", "s", "", "Set if the .lagoon.yml should be found in a subfolder useful if you have multiple Lagoon projects per Git Repository")
	addProjectCmd.Flags().StringP("router-pattern", "Z", "", "Router pattern of the project, e.g. '${service}-${environment}-${project}.lagoon.example.com'")
	addProjectCmd.Flags().StringP("branches", "b", "", "Which branches should be deployed")
	addProjectCmd.Flags().StringP("pullrequests", "m", "", "Which Pull Requests should be deployed")
	addProjectCmd.Flags().StringP("production-environment", "E", "", "Which environment(the name) should be marked as the production environment")
	addProjectCmd.Flags().String("standby-production-environment", "", "Which environment(the name) should be marked as the standby production environment")
	addProjectCmd.Flags().StringP("deploytarget-project-pattern", "", "", "Pattern of Deploytarget(Kubernetes) Project/Namespace that should be generated")

	addProjectCmd.Flags().UintP("auto-idle", "a", 0, "Auto idle setting of the project")
	addProjectCmd.Flags().UintP("storage-calc", "C", 0, "Should storage for this environment be calculated")
	addProjectCmd.Flags().UintP("development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().UintP("deploytarget", "S", 0, "Reference to Deploytarget(Kubernetes) target this Project should be deployed to")
	addProjectCmd.Flags().StringP("build-image", "", "", "Build Image for the project")
	addProjectCmd.Flags().Bool("owner", false, "Add the user as an owner of the project")
	addProjectCmd.Flags().StringP("organization-name", "O", "", "Name of the Organization to add the project to")
	addProjectCmd.Flags().UintP("organization-id", "", 0, "ID of the Organization to add the project to")

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
