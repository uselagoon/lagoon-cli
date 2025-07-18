package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/uselagoon/lagoon-cli/internal/util"
	"github.com/uselagoon/lagoon-cli/internal/wizard/project"
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

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		if yesNo(fmt.Sprintf("You are attempting to delete project '%s', are you sure?", cmdProjectName)) {
			_, err := lagoon.DeleteProject(context.TODO(), cmdProjectName, lc)
			if err != nil {
				return err
			}
			resultData := output.Result{
				Result: "success",
			}
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		developmentEnvironmentsLimitProvided := cmd.Flags().Lookup("development-environments-limit").Changed
		storageCalc, err := cmd.Flags().GetBool("storage-calc")
		if err != nil {
			return err
		}
		storageCalcProvided := cmd.Flags().Lookup("storage-calc").Changed
		autoIdle, err := cmd.Flags().GetBool("auto-idle")
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
		orgOwner, err := cmd.Flags().GetBool("owner")
		if err != nil {
			return err
		}
		orgOwnerProvided := cmd.Flags().Lookup("owner").Changed
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		interactive, err := cmd.Flags().GetBool("interactive")
		if err != nil {
			return err
		}
		generatedCommand := ""

		if !interactive {
			if err := requiredInputCheck("Project name", cmdProjectName, "git-url", gitUrl, "Production environment", productionEnvironment, "Deploytarget", strconv.Itoa(int(deploytarget))); err != nil {
				return err
			}
		}

		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)

		projectInput := schema.AddProjectInput{}
		if interactive {
			config, err := project.RunCreateWizard(lc)
			if err != nil {
				if errors.Is(err, huh.ErrUserAborted) {
					return nil
				}
				return err
			}
			projectInput = config.Input
			if config.OrganizationDetails.Name != "" {
				organizationName = config.OrganizationDetails.Name
			}
			if config.AutoIdleProvided {
				autoIdle = config.AutoIdle
				autoIdleProvided = config.AutoIdleProvided
			}
			if config.StorageCalcProvided {
				storageCalc = config.StorageCalc
				storageCalcProvided = config.StorageCalcProvided
			}
			if config.DevEnvLimit != 0 {
				developmentEnvironmentsLimit = config.DevEnvLimit
				developmentEnvironmentsLimitProvided = true
			}
			generatedCommand = util.GenerateCLICommand(config)
		}

		if !interactive {
			projectInput = schema.AddProjectInput{
				Name:                         cmdProjectName,
				GitURL:                       gitUrl,
				ProductionEnvironment:        productionEnvironment,
				StandbyProductionEnvironment: standbyProductionEnvironment,
				Branches:                     branches,
				PullRequests:                 pullrequests,
				OpenshiftProjectPattern:      deploytargetProjectPattern,
				Openshift:                    deploytarget,
				Subfolder:                    subfolder,
				PrivateKey:                   privateKey,
				BuildImage:                   buildImage,
				RouterPattern:                routerPattern,
			}
		}
		if orgOwnerProvided {
			projectInput.AddOrgOwner = &orgOwner
		}
		if storageCalcProvided {
			projectInput.StorageCalc = nullBoolToUint(storageCalc)
		}
		if autoIdleProvided {
			projectInput.AutoIdle = nullBoolToUint(autoIdle)
		}
		if developmentEnvironmentsLimitProvided {
			projectInput.DevelopmentEnvironmentsLimit = nullUintCheck(developmentEnvironmentsLimit)
		}
		// if organizationid is provided, use it over the name
		if organizationID != 0 {
			projectInput.Organization = &organizationID
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
			projectInput.Organization = &organization.ID
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
				"GitURL":       projectInput.GitURL,
			},
		}
		if organizationName != "" {
			resultData.ResultData["Organization"] = organizationName
		}
		if interactive {
			resultData.ResultData["Generated Command"] = "lagoon add project" + generatedCommand
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		developmentEnvironmentsLimitProvided := cmd.Flags().Lookup("development-environments-limit").Changed
		if err != nil {
			return err
		}
		storageCalc, err := cmd.Flags().GetBool("storage-calc")
		if err != nil {
			return err
		}
		storageCalcProvided := cmd.Flags().Lookup("storage-calc").Changed
		autoIdle, err := cmd.Flags().GetBool("auto-idle")
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
		factsUi, err := cmd.Flags().GetBool("facts-ui")
		if err != nil {
			return err
		}
		factsUIProvided := cmd.Flags().Lookup("facts-ui").Changed
		problemsUi, err := cmd.Flags().GetBool("problems-ui")
		if err != nil {
			return err
		}
		problemsUIProvided := cmd.Flags().Lookup("problems-ui").Changed
		routerPattern, err := cmd.Flags().GetString("router-pattern")
		if err != nil {
			return err
		}
		deploymentsDisabled, err := cmd.Flags().GetBool("deployments-disabled")
		if err != nil {
			return err
		}
		deploymentsDisabledProvided := cmd.Flags().Lookup("deployments-disabled").Changed
		productionBuildPriority, err := cmd.Flags().GetUint("production-build-priority")
		if err != nil {
			return err
		}
		productionBuildPriorityProvided := cmd.Flags().Lookup("production-build-priority").Changed
		developmentBuildPriority, err := cmd.Flags().GetUint("development-build-priority")
		if err != nil {
			return err
		}
		developmentBuildPriorityProvided := cmd.Flags().Lookup("development-build-priority").Changed

		if err := requiredInputCheck("Project name", cmdProjectName); err != nil {
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

		projectPatch := schema.UpdateProjectPatchInput{
			GitURL:                       nullStrCheck(gitUrl),
			ProductionEnvironment:        nullStrCheck(productionEnvironment),
			Openshift:                    nullUintCheck(deploytarget),
			StandbyProductionEnvironment: nullStrCheck(standbyProductionEnvironment),
			Branches:                     nullStrCheck(branches),
			Pullrequests:                 nullStrCheck(pullrequests),
			OpenshiftProjectPattern:      nullStrCheck(deploytargetProjectPattern),
			Subfolder:                    nullStrCheck(subfolder),
			PrivateKey:                   nullStrCheck(privateKey),
			RouterPattern:                nullStrCheck(routerPattern),
			Name:                         nullStrCheck(projectName),
		}

		if availability != "" {
			ProjectAvailability := schema.ProjectAvailability(strings.ToUpper(availability))
			projectPatch.Availability = &ProjectAvailability
		}
		if storageCalcProvided {
			projectPatch.StorageCalc = nullBoolToUint(storageCalc)
		}
		if autoIdleProvided {
			projectPatch.AutoIdle = nullBoolToUint(autoIdle)
		}
		if deploymentsDisabledProvided {
			projectPatch.DeploymentsDisabled = nullBoolToUint(deploymentsDisabled)
		}
		if factsUIProvided {
			projectPatch.FactsUI = nullBoolToUint(factsUi)
		}
		if problemsUIProvided {
			projectPatch.ProblemsUI = nullBoolToUint(problemsUi)
		}
		if productionBuildPriorityProvided {
			projectPatch.ProductionBuildPriority = nullUintCheck(productionBuildPriority)
		}
		if developmentBuildPriorityProvided {
			projectPatch.DevelopmentBuildPriority = nullUintCheck(developmentBuildPriority)
		}
		if developmentEnvironmentsLimitProvided {
			projectPatch.DevelopmentEnvironmentsLimit = nullUintCheck(developmentEnvironmentsLimit)
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

		projectUpdate, err := lagoon.UpdateProjectByName(context.TODO(), cmdProjectName, projectPatch, lc)
		if err != nil {
			return fmt.Errorf("%v: check if the project exists", err.Error())
		}

		resultData := output.Result{
			Result: "success",
			ResultData: map[string]interface{}{
				"Project Name": projectUpdate.Name,
			},
		}
		r := output.RenderResult(resultData, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		projects, err := lagoon.GetProjectsByMetadata(context.TODO(), key, value, lc)
		if err != nil {
			return err
		}
		if len(*projects) == 0 {
			if value != "" {
				return handleNilResults("No projects found with metadata key '%s' and value '%s'\n", cmd, key, value)
			}
			return handleNilResults("No projects found with metadata key '%s'\n", cmd, key)
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
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
		current := lagoonCLIConfig.Current
		token := lagoonCLIConfig.Lagoons[current].Token
		lc := lclient.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIVersion,
			lagoonCLIConfig.Lagoons[current].Version,
			&token,
			debug)
		project, err := lagoon.GetProjectMetadata(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if len(project.Metadata) == 0 {
			return handleNilResults("There is no metadata for project '%s'\n", cmd, cmdProjectName)
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
		r := output.RenderOutput(output.Table{
			Header: header,
			Data:   data,
		}, outputOptions)
		fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)

			projectResult, err := lagoon.UpdateProjectMetadataByName(context.TODO(), cmdProjectName, key, value, lc)
			if err != nil {
				return fmt.Errorf("%v: check if the project exists", err.Error())
			}
			data := []output.Data{}
			metaData, _ := json.Marshal(projectResult.Metadata)
			data = append(data, []string{
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.ID)),
				returnNonEmptyString(fmt.Sprintf("%v", projectResult.Name)),
				returnNonEmptyString(fmt.Sprintf("%v", string(metaData))),
			})
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Metadata",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
			current := lagoonCLIConfig.Current
			token := lagoonCLIConfig.Lagoons[current].Token
			lc := lclient.New(
				lagoonCLIConfig.Lagoons[current].GraphQL,
				lagoonCLIVersion,
				lagoonCLIConfig.Lagoons[current].Version,
				&token,
				debug)

			projectResult, err := lagoon.RemoveProjectMetadataByKeyByName(context.TODO(), cmdProjectName, key, lc)
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
			r := output.RenderOutput(output.Table{
				Header: []string{
					"ID",
					"Name",
					"Metadata",
				},
				Data: data,
			}, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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

		// this can stay for now as `removeProjectFromOrganization` is platform only scoped
		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}
		if project.Name == "" {
			return handleNilResults("No project found for '%s'\n", cmd, cmdProjectName)
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
			r := output.RenderResult(resultData, outputOptions)
			fmt.Fprintf(cmd.OutOrStdout(), "%s", r)
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
	updateProjectCmd.Flags().BoolP("auto-idle", "a", false, "Auto idle setting of the project. Set to enable, --auto-idle=false to disable")
	updateProjectCmd.Flags().BoolP("storage-calc", "C", false, "Should storage for this environment be calculated. Set to enable, --storage-calc=false to disable")
	updateProjectCmd.Flags().UintP("development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	updateProjectCmd.Flags().UintP("deploytarget", "S", 0, "Reference to Deploytarget(Kubernetes) this Project should be deployed to")
	updateProjectCmd.Flags().BoolP("deployments-disabled", "", false, "Admin only flag for disabling deployments on a project. Set to disable deployments, --deployments-disabled=false to enable")

	updateProjectCmd.Flags().BoolP("facts-ui", "", false, "Enables the Lagoon insights Facts tab in the UI. Set to enable, --facts-ui=false to disable")
	updateProjectCmd.Flags().BoolP("problems-ui", "", false, "Enables the Lagoon insights Problems tab in the UI. Set to enable, --problems-ui=false to disable")

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

	addProjectCmd.Flags().BoolP("auto-idle", "a", false, "Auto idle setting of the project. Set to enable, --auto-idle=false to disable")
	addProjectCmd.Flags().BoolP("storage-calc", "C", false, "Should storage for this environment be calculated. Set to enable, --storage-calc=false to disable")
	addProjectCmd.Flags().UintP("development-environments-limit", "L", 0, "How many environments can be deployed at one time")
	addProjectCmd.Flags().UintP("deploytarget", "S", 0, "Reference to Deploytarget(Kubernetes) target this Project should be deployed to")
	addProjectCmd.Flags().StringP("build-image", "", "", "Build Image for the project")
	addProjectCmd.Flags().Bool("owner", false, "Add the user as an owner of the project")
	addProjectCmd.Flags().StringP("organization-name", "O", "", "Name of the Organization to add the project to")
	addProjectCmd.Flags().UintP("organization-id", "", 0, "ID of the Organization to add the project to")
	addProjectCmd.Flags().Bool("interactive", false, "Set Interactive mode for the project creation wizard.")

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
