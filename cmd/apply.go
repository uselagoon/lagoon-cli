package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/output"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type FileConfigRoot struct {
	Event    EventConfig          `json:"event,omitempty" yaml:"event,omitempty"`
	Tasks    []AdvancedTasksInput `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Workflow WorkflowConfig       `json:"workflow,omitempty" yaml:"workflow,omitempty"`
	Settings Settings             `json:"settings,omitempty" yaml:"settings,omitempty"`
}

type EventConfig struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type AdvancedTaskDefinitionArgument struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type Command struct {
	Cmd       string          `json:"cmd,omitempty" yaml:"cmd,omitempty"`
	Arguments CommandArgument `json:"arguments,omitempty" yaml:"arguments,omitempty"`
}

type CommandArgument struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type Settings struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

type WorkflowConfig struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

func ApplyAdvancedTaskDefinitions(lc *client.Client, fileConfig *FileConfigRoot) error {
	var advancedTaskDefinitionResult *schema.AdvancedTaskDefinition
	var data []output.Data

	// Preprocess validation
	advancedTasksJSON, err := json.Marshal(fileConfig.Tasks)
	if err != nil {
		return err
	}
	PreprocessAdvancedTaskDefinitionsInputValidation(advancedTasksJSON)

	// Add task definitions for each task found
	if len(fileConfig.Tasks) > 0 {
		for _, t := range fileConfig.Tasks {
			var hasTaskMatches = false

			// Get project environments from name
			project, err := lagoon.GetMinimalProjectByName(context.TODO(), t.Project, lc)
			if err != nil {
				return err
			}

			// Find environment ID if given
			var envID int
			if len(project.Environments) > 0 && t.Environment != "" {
				for _, e := range project.Environments {
					if e.Name == t.Environment {
						envID = int(e.ID)
					}
				}
			}

			// Get current tasks by Environment ID
			liveAdvancedTasks, err := lagoon.GetAdvancedTasksByEnvironment(context.TODO(), envID, lc)
			if err != nil {
				return err
			}

			// Check if given task already exists
			if liveAdvancedTasks != nil {
				for _, currAdvTask := range *liveAdvancedTasks {
					// Convert AdvancedTaskDefinition struct from file to input the update mutation can understand
					advancedTaskInput := &schema.AdvancedTaskDefinitionInput{
						ID:          currAdvTask.ID,
						Name:        t.Name,
						Description: t.Description,
						Type:        t.Type,
						Service:     t.Service,
						Image:       t.Image,
						Command:     t.Command,
						GroupName:   t.GroupName,
						Project:     currAdvTask.Project,
						Environment: currAdvTask.Environment,
						Permission:  t.Permission,
					}

					// Marshal to JSON to flip capitalisation of struct keys from yaml unmarshalling
					encodedJSONTaskInput, err := json.Marshal(advancedTaskInput)
					if err != nil {
						fmt.Errorf("Unable to marhsal yaml config to JSON '%s': %v", t, err)
					}

					// If matched task name exists, we diff and update
					if currAdvTask.Name == advancedTaskInput.Name {
						hasTaskMatches = true

						diffString, hasChanges, diffErr := DiffPatchChangesAgainstAPI(currAdvTask, encodedJSONTaskInput)
						if diffErr != nil {
							return diffErr
						}
						if !hasChanges {
							log.Printf("No changes found for task '%s'", advancedTaskInput.Name)
						}

						if hasChanges {
							log.Println("The following changes will be applied:\n", diffString)
							if forceAction || !forceAction && yesNo(fmt.Sprintf("Are you sure you want to update '%s'", advancedTaskInput.Name)) {
								// Update changes
								advancedTaskDefinitionResult, err = lagoon.UpdateAdvancedTaskDefinition(context.TODO(), int(advancedTaskInput.ID), &schema.AdvancedTaskDefinitionInput{
									Name:        advancedTaskInput.Name,
									Description: advancedTaskInput.Description,
									Type:        advancedTaskInput.Type,
									Service:     advancedTaskInput.Service,
									Image:       advancedTaskInput.Image,
									Command:     advancedTaskInput.Command,
									GroupName:   advancedTaskInput.GroupName,
									Project:     advancedTaskInput.Project,
									Environment: advancedTaskInput.Environment,
									Permission:  advancedTaskInput.Permission,
									//AdvancedTaskDefinitionArguments: advancedTaskInput.AdvancedTaskDefinitionArgument,
								}, lc)
								if err != nil {
									return err
								}
							}
						}
					} else if hasTaskMatches != true {
						hasTaskMatches = false
					}
				}
			}

			// If no match - we add a new one
			if !hasTaskMatches {
				if yesNo(fmt.Sprintf("You are attempting to add a new task '%s', are you sure?", t.Name)) {
					// Add TaskDefinition
					advancedTaskDefinitionResult, err = lagoon.AddAdvancedTaskDefinition(context.TODO(), &schema.AdvancedTaskDefinitionInput{
						Name:        t.Name,
						Description: t.Description,
						Type:        t.Type,
						Service:     t.Service,
						Image:       t.Image,
						Command:     t.Command,
						GroupName:   t.GroupName,
						Project:     int(project.ID),
						Environment: envID,
						Permission:  t.Permission,
					}, lc)
					if err != nil {
						return err
					}
					fmt.Println("successfully added task definition with ID:", advancedTaskDefinitionResult.ID)
				}
			}
		}
	}

	if advancedTaskDefinitionResult != nil {
		data = append(data, []string{
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Description)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Type)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Service)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.GroupName)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Project)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Environment)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Image)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Command)),
		})
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
				"Description",
				"Type",
				"Service",
				"Image",
				"Command",
				"GroupName",
				"Project",
				"Environment",
				"Permission",
			},
			Data: data,
		}, outputOptions)
	}

	return nil
}

func parseFile(file string) (*FileConfigRoot, error) {
	// Check file exists
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}

	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read config file contents '%s': %v", file, err)
	}

	// Unmarshal yaml
	parsedConfig := &FileConfigRoot{}
	err = yaml.Unmarshal(source, &parsedConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode config in file '%s': %v", file, err)
	}

	return parsedConfig, nil
}

var viewLastApplied = &cobra.Command{
	Use:   "view-last-applied",
	Short: "View the latest applied workflows or advanced task definitions for project/environment.",
	Long:  `View the latest applied workflows or advanced task definitions for project/environment.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		if cmdProjectEnvironment == "" || cmdProjectName == "" {
			fmt.Println("Missing arguments: Project name or environment name is not defined")
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

		// Get project environment from name
		project, err := lagoon.GetMinimalProjectByName(context.TODO(), cmdProjectName, lc)
		if err != nil {
			return err
		}

		var envID int
		if project.Environments != nil {
			if len(project.Environments) > 0 && cmdProjectEnvironment != "" {
				for _, e := range project.Environments {
					if e.Name == cmdProjectEnvironment {
						envID = int(e.ID)
					}
				}
			}
		}

		// Get current advanced tasks by Environment ID
		advancedTasks, err := lagoon.GetAdvancedTasksByEnvironment(context.TODO(), envID, lc)
		if err != nil {
			return err
		}
		if advancedTasks != nil {
			fmt.Printf("Found %d tasks for '%s:%s'\n", len(*advancedTasks), cmdProjectName, cmdProjectEnvironment)
			for _, task := range *advancedTasks {
				taskJSON, err := json.Marshal(task)
				if err != nil {
					return err
				}

				resultData := output.Result{
					Result: string(taskJSON),
				}
				output.RenderResult(resultData, outputOptions)
			}
		}
		return nil
	},
}

var setLastApplied = &cobra.Command{
	Use:   "set-last-applied -f FILENAME",
	Short: "View the latest applied workflows or advanced task definitions for project/environment.",
	Long:  `View the latest applied workflows or advanced task definitions for project/environment.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		fileName, err := cmd.Flags().GetString("file")
		// Read yaml config from file
		if fileName == "" {
			fmt.Errorf("Missing arguments: 'file' is not defined")
		}
		fileConfig, err := parseFile(fileName)
		if err != nil {
			log.Fatalln("Parsing config error:", err)
		}

		// Validate tasks input
		advancedTasksJSON, err := json.Marshal(fileConfig.Tasks)
		if err != nil {
			return err
		}
		PreprocessAdvancedTaskDefinitionsInputValidation(advancedTasksJSON)

		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		// Apply Advanced Tasks
		err = ApplyAdvancedTaskDefinitions(lc, fileConfig)
		if err != nil {
			return err
		}

		return nil
	},
}

var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"ap"},
	Short:   "Apply the configuration of workflows or tasks from a given yaml configuration file",
	Long: `Apply the configuration of workflows or tasks from a given yaml configuration file.
Workflows or task definitions will be created if they do not already exist.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")

		// Read yaml config from file
		fileName, err := cmd.Flags().GetString("file")
		fileConfig, err := parseFile(fileName)
		if err != nil {
			log.Fatalln("Parsing config error:", err)
		}

		// Create lagoon client
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		// Apply Advanced Tasks
		err = ApplyAdvancedTaskDefinitions(lc, fileConfig)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	applyCmd.PersistentFlags().StringP("file", "f", "", "lagoon apply (-f FILENAME) [options]")
	applyCmd.MarkFlagRequired("file")
	applyCmd.Flags().BoolVarP(&showAdvancedTasks, "advanced-tasks", "t", false, "Target advanced tasks only")
	applyCmd.Flags().BoolVarP(&showWorkflows, "workflows", "w", false, "Target workflows only")

	applyCmd.AddCommand(viewLastApplied)
	applyCmd.AddCommand(setLastApplied)
}
