package project

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/uselagoon/lagoon-cli/internal/util"
	"github.com/uselagoon/machinery/api/lagoon"
	"github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
	"log"
	"strconv"
)

type CreateConfig struct {
	Input               schema.AddProjectInput
	OrganizationName    string // need to update current schema in Machinery to utilize organizationDetails
	AutoIdle            bool
	AutoIdleProvided    bool
	StorageCalc         bool
	StorageCalcProvided bool
	DevEnvLimit         uint
}

func RunCreateWizard(lc *client.Client) (*CreateConfig, error) {
	config := &CreateConfig{}
	var organizationConfirm bool
	initForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the project name").
				Value(&config.Input.Name).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("Project name is required")
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Do you want to create this project in an Organization?").
				Value(&organizationConfirm),
		),
	).WithTheme(huh.ThemeDracula())

	formErr := initForm.Run()
	if formErr != nil {
		log.Fatal(formErr)
	}
	if organizationConfirm {
		organizations, err := lagoon.AllOrganizations(context.TODO(), lc)
		if err != nil {
			return nil, err
		}
		form2 := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select an organization").
					OptionsFunc(func() []huh.Option[string] {
						options := make([]huh.Option[string], len(*organizations))
						for i, org := range *organizations {
							options[i] = huh.NewOption(org.Name, org.Name)
						}

						return options
					}, &config.OrganizationName).
					Value(&config.OrganizationName),
			),
		).WithTheme(huh.ThemeDracula())

		err = form2.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	deploytargets, err := lagoon.ListDeployTargets(context.TODO(), lc)
	var additionalFields bool
	form3 := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter the Git URL for your project").
				Value(&config.Input.GitURL).
				Validate(func(str string) error {
					if !util.IsValidGitURL(config.Input.GitURL) {
						return errors.New("Invalid Git URL")
					}
					return nil
				}),
			huh.NewInput().
				Title("Enter the branch name for the production environment").
				Value(&config.Input.ProductionEnvironment),
			huh.NewSelect[uint]().
				Title("Select a deploy target").
				OptionsFunc(func() []huh.Option[uint] {
					options := make([]huh.Option[uint], len(*deploytargets))
					for i, target := range *deploytargets {
						options[i] = huh.NewOption(target.Name, target.ID)
					}

					return options
				}, &config.Input.Openshift).
				Value(&config.Input.Openshift),
			huh.NewConfirm().
				Title("Do you want to define any other fields?").
				Value(&additionalFields),
		),
	).WithTheme(huh.ThemeDracula())

	err = form3.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}

	if additionalFields {
		var fields []string
		form4 := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					Title("Select which fields you want to define").
					Options(
						huh.NewOption("standby-production-environment", "StandbyProductionEnvironment"),
						huh.NewOption("branches", "Branches"),
						huh.NewOption("pullrequests", "PullRequests"),
						huh.NewOption("deploytarget-project-pattern", "OpenshiftProjectPattern"),
						huh.NewOption("development-environments-limit", "DevelopmentEnvironmentsLimit"),
						huh.NewOption("auto-idle", "AutoIdle"),
						huh.NewOption("subfolder", "Subfolder"),
						huh.NewOption("private-key", "PrivateKey"),
						huh.NewOption("build-image", "BuildImage"),
						huh.NewOption("router-pattern", "RouterPattern"),
						huh.NewOption("owner (Only select if adding to an Organization)", "AddOrgOwner"),
						huh.NewOption("storage-calc", "StorageCalc "),
					).
					Value(&fields),
			),
		).WithTheme(huh.ThemeDracula())

		err = form4.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}

		if len(fields) == 0 {
			fmt.Println("No fields selected.")
		}

		var inputs []huh.Field
		var devEnvLimit string
		for _, field := range fields {
			switch field {
			case "StandbyProductionEnvironment":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.StandbyProductionEnvironment))
			case "Branches":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.Branches))
			case "PullRequests":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.PullRequests))
			case "OpenshiftProjectPattern":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.OpenshiftProjectPattern))
			case "Subfolder":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.Subfolder))
			case "PrivateKey":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.PrivateKey))
			case "BuildImage":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.BuildImage))
			case "RouterPattern":
				inputs = append(inputs, huh.NewInput().Title(fmt.Sprintf("Enter value for '%s'", field)).Value(&config.Input.RouterPattern))

			case "DevelopmentEnvironmentsLimit":
				inputs = append(inputs, huh.NewInput().
					Title(fmt.Sprintf("Enter value for '%s'", field)).
					Value(&devEnvLimit).
					Validate(func(s string) error {
						if s == "" {
							return nil
						}
						_, err := strconv.ParseUint(s, 10, 64)
						return err
					}),
				)

			case "StorageCalc":
				config.StorageCalcProvided = true
				inputs = append(inputs, huh.NewConfirm().
					Title(field).
					Value(&config.StorageCalc))
			case "AutoIdle":
				config.AutoIdleProvided = true
				inputs = append(inputs, huh.NewConfirm().
					Title(field).
					Value(&config.AutoIdle))
			case "AddOrgOwner":
				inputs = append(inputs, huh.NewConfirm().Title(field).Value(config.Input.AddOrgOwner))
			}
		}

		form5 := huh.NewForm(
			huh.NewGroup(inputs...),
		).WithTheme(huh.ThemeDracula())

		err = form5.Run()
		if err != nil {
			fmt.Println("Error:", err)
		}

		if devEnvLimit != "" {
			developmentEnvLimit, err := strconv.Atoi(devEnvLimit)
			if err != nil {
				fmt.Println("Error:", err)
			}
			config.DevEnvLimit = uint(developmentEnvLimit)
		}

	}

	return config, nil
}
