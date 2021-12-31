package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"

	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/lagoon/client"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/output"
)

type FileConfigRoot struct {
	Tasks     []AdvancedTasksFileInput `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Workflows []WorkflowsFileInput     `json:"workflows,omitempty" yaml:"workflows,omitempty"`
	Settings  Settings                 `json:"settings,omitempty" yaml:"settings,omitempty"`
}

type WorkflowsFileInput struct {
	ID                       uint   `json:"id,omitempty" yaml:"id,omitempty"`
	Name                     string `json:"name,omitempty" yaml:"name,omitempty"`
	Event                    string `json:"event,omitempty" yaml:"event,omitempty"`
	Project                  string `json:"project,omitempty" yaml:"project,omitempty"`
	Environment              string `json:"environment,omitempty" yaml:"environment,omitempty"`
	AdvancedTaskDefinition   string `json:"task,omitempty" yaml:"task,omitempty"`
	AdvancedTaskDefinitionID int
	Target                   string   `json:"target,omitempty" yaml:"target,omitempty"`
	Cron                     string   `json:"cron,omitempty" yaml:"cron,omitempty"`
	Settings                 Settings `json:"settings,omitempty" yaml:"settings,omitempty"`
}

type AdvancedTasksFileInput struct {
	ID            uint                                    `json:"id,omitempty" yaml:"id,omitempty"`
	Name          string                                  `json:"name,omitempty" yaml:"name,omitempty"`
	Description   string                                  `json:"description,omitempty" yaml:"description,omitempty"`
	Type          schema.AdvancedTaskDefinitionType       `json:"type,omitempty" yaml:"type,omitempty"`
	Command       string                                  `json:"command,omitempty" yaml:"command,omitempty"`
	Image         string                                  `json:"image,omitempty" yaml:"image,omitempty"`
	Service       string                                  `json:"service,omitempty" yaml:"service,omitempty"`
	GroupName     string                                  `json:"group,omitempty" yaml:"group,omitempty"`
	Project       string                                  `json:"project,omitempty" yaml:"project,omitempty"`
	Environment   string                                  `json:"environment,omitempty" yaml:"environment,omitempty"`
	EnvironmentID int                                     `json:"environmentID,omitempty" yaml:"environmentID,omitempty"`
	Permission    string                                  `json:"permission,omitempty" yaml:"permission,omitempty"`
	Arguments     []schema.AdvancedTaskDefinitionArgument `json:"arguments,omitempty" yaml:"arguments,omitempty"`
}

type Settings struct {
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}

var fileName string
var debug bool

// ApplyWorkflows Apply workflow configurations from file.
func ApplyWorkflows(lc *client.Client, workflows []WorkflowsFileInput) error {
	var workflowsResult *schema.WorkflowResponse
	var data []output.Data

	if len(workflows) > 0 {
		for _, w := range workflows {
			var hasWorkflowMatches = false

			project, err := lagoon.GetMinimalProjectByName(context.TODO(), w.Project, lc)
			if err != nil {
				return err
			}

			// Find environment ID if given
			var envID int
			if len(project.Environments) > 0 && w.Environment != "" {
				for _, e := range project.Environments {
					if e.Name == w.Environment {
						envID = int(e.ID)
					}
				}
			}

			// Get current workflows by Environment ID
			liveWorkflows, err := lagoon.GetWorkflowsByEnvironment(context.TODO(), envID, lc)
			if err != nil {
				return err
			}

			// Match Event input string to a known Lagoon event type
			eventType := matchEventToEventType(w.Event)
			if eventType == "" {
				return fmt.Errorf("event type did not match a Lagoon event type")
			}

			// Match Advanced Task Definition string input to associated Lagoon int id.
			liveAdvancedTasks, err := lagoon.GetAdvancedTasksByEnvironment(context.TODO(), envID, lc)
			if err != nil {
				return err
			}

			if liveAdvancedTasks != nil {
				for _, task := range *liveAdvancedTasks {
					if task.Name == w.AdvancedTaskDefinition {
						w.AdvancedTaskDefinitionID = int(task.ID)
					}
				}
			}

			// If advanced task given does not exist inside project, then exit.
			if w.AdvancedTaskDefinitionID == 0 {
				fmt.Printf("task '%s' does not match one in Lagoon resource '%s', you should create it first\n", w.AdvancedTaskDefinition, w.Project)
				os.Exit(1)
			}

			// Check if given workflows already exists - if they do, we attempt to update them.
			if liveWorkflows != nil {
				for _, currWorkflow := range *liveWorkflows {
					// Convert struct from file config to input the update mutation can understand
					workflowInput := &schema.WorkflowInput{
						ID:                     currWorkflow.ID,
						Name:                   w.Name,
						Event:                  eventType,
						AdvancedTaskDefinition: w.AdvancedTaskDefinitionID,
						Project:                currWorkflow.Project,
					}

					encodedDiffWorkflowInput, err := json.Marshal(struct {
						ID                     uint                               `json:"id,omitempty"`
						Name                   string                             `json:"name,omitempty"`
						Event                  schema.EventType                   `json:"event,omitempty"`
						AdvancedTaskDefinition schema.AdvancedTaskDefinitionInput `json:"advancedTaskDefinition,omitempty"`
						Project                int                                `json:"project,omitempty"`
					}{
						ID:    currWorkflow.ID,
						Name:  w.Name,
						Event: eventType,
						AdvancedTaskDefinition: schema.AdvancedTaskDefinitionInput{
							ID: currWorkflow.AdvancedTaskDefinition.ID,
						},
						Project: currWorkflow.Project,
					})
					if err != nil {
						return fmt.Errorf("unable to marhsal yaml config to JSON '%v': %w", w, err)
					}

					// If matched workflow exists, we diff and update
					if currWorkflow.Name == workflowInput.Name {
						hasWorkflowMatches = true

						diffString, diffErr := DiffPatchChangesAgainstAPI(encodedDiffWorkflowInput, currWorkflow)
						if diffErr != nil {
							return diffErr
						}

						if diffString == "" {
							log.Printf("No changes found for workflow '%s'", currWorkflow.Name)
						}

						if diffString != "" {
							log.Println("The following changes will be applied:\n", diffString)
							if forceAction || !forceAction && yesNo(fmt.Sprintf("Are you sure you want to update '%s'", workflowInput.Name)) {
								// Unset ID as it's not required for patch.
								workflowInput.ID = 0
								workflowsResult, err = lagoon.UpdateWorkflow(context.TODO(), int(currWorkflow.ID), workflowInput, lc)
								if err != nil {
									return err
								}
							}
						}
					} else if !hasWorkflowMatches {
						hasWorkflowMatches = false
					}
				}
			}

			// If no match - we add a new one
			if !hasWorkflowMatches {
				if yesNo(fmt.Sprintf("You are attempting to add a new workflow '%s', are you sure?", w.Name)) { // Add TaskDefinition
					workflowsResult, err = lagoon.AddWorkflow(context.TODO(), &schema.WorkflowInput{
						Name:                   w.Name,
						Event:                  eventType,
						Project:                int(project.ID),
						AdvancedTaskDefinition: w.AdvancedTaskDefinitionID,
					}, lc)
					if err != nil {
						return err
					}
					fmt.Println("successfully added workflow with ID:", workflowsResult.ID)
				}
			}
		}
	}

	if workflowsResult != nil {
		data = append(data, []string{
			returnNonEmptyString(fmt.Sprintf("%v", workflowsResult.ID)),
			returnNonEmptyString(fmt.Sprintf("%v", workflowsResult.Name)),
			returnNonEmptyString(fmt.Sprintf("%v", workflowsResult.Event)),
			returnNonEmptyString(fmt.Sprintf("%v", workflowsResult.Project)),
			returnNonEmptyString(fmt.Sprintf("%v", workflowsResult.AdvancedTaskDefinition.ID)),
		})
		output.RenderOutput(output.Table{
			Header: []string{
				"ID",
				"Name",
				"Event",
				"Project",
				"Advanced Task",
			},
			Data: data,
		}, outputOptions)
	}

	return nil
}

// ApplyAdvancedTaskDefinitions Apply advanced task definition configurations from file.
func ApplyAdvancedTaskDefinitions(lc *client.Client, tasks []AdvancedTasksFileInput) error {
	var advancedTaskDefinitionResult *schema.AdvancedTaskDefinitionResponse
	var data []output.Data

	// Add task definitions for each task found
	if len(tasks) > 0 {
		for _, t := range tasks {
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
					// Convert AdvancedTaskDefinitionResponse struct from file to input the update mutation can understand
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
						Arguments:   t.Arguments,
					}

					// Marshal to JSON to flip capitalisation of struct keys from yaml unmarshalling
					encodedJSONTaskInput, err := json.Marshal(advancedTaskInput)
					if err != nil {
						return fmt.Errorf("unable to marhsal yaml config to JSON '%v': %w", t, err)
					}

					// If matched task name exists, we diff and update
					if currAdvTask.Name == advancedTaskInput.Name {
						hasTaskMatches = true

						diffString, diffErr := DiffPatchChangesAgainstAPI(encodedJSONTaskInput, currAdvTask)
						if diffErr != nil {
							return diffErr
						}

						if diffString == "" {
							log.Printf("No changes found for task '%s'", advancedTaskInput.Name)
						}

						if diffString != "" {
							log.Println("The following changes will be applied:\n", diffString)
							if forceAction || !forceAction && yesNo(fmt.Sprintf("Are you sure you want to update '%s'", advancedTaskInput.Name)) {
								// Unset ID as it's not required for task patch.
								advancedTaskInput.ID = 0
								// Update changes
								advancedTaskDefinitionResult, err = lagoon.UpdateAdvancedTaskDefinition(context.TODO(), int(currAdvTask.ID), advancedTaskInput, lc)
								if err != nil {
									return err
								}
							}
						}
					} else if !hasTaskMatches {
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
						Environment: t.EnvironmentID,
						Permission:  t.Permission,
						Arguments:   t.Arguments,
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
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Image)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Command)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.GroupName)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Project)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Environment)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Permission)),
			returnNonEmptyString(fmt.Sprintf("%v", advancedTaskDefinitionResult.Arguments)),
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
				"Arguments",
			},
			Data: data,
		}, outputOptions)
	}

	return nil
}

// ReadConfigFile Checks if file exists, reads yaml config from file name and returns the config.
func ReadConfigFile(fileName string) (*FileConfigRoot, error) {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist")
		}
	}
	fileConfig, err := parseFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("parsing config error %w", err)
	}

	return fileConfig, nil
}

// Attempts to parse the configuration
func parseFile(file string) (*FileConfigRoot, error) {
	source, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read config file contents '%s': %v", file, err)
	}

	// Unmarshal yaml
	parsedConfig := &FileConfigRoot{}
	err = yaml.Unmarshal(source, &parsedConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config in file '%s': %v", file, err)
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
		fileName, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		var tasksInput []AdvancedTasksFileInput
		if fileName != "" {
			fileConfig, err := ReadConfigFile(fileName)
			if err != nil {
				log.Fatalln("Parsing config error:", err)
			}

			// Preprocess validation
			tasksInput, err = PreprocessAdvancedTaskDefinitionsInputValidation(fileConfig.Tasks)
			if err != nil {
				return err
			}
		} else {
			tasksInput = append(tasksInput, AdvancedTasksFileInput{
				Project:     cmdProjectName,
				Environment: cmdProjectEnvironment,
			})
		}

		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		for _, t := range tasksInput {
			// Get project environment from name
			project, err := lagoon.GetMinimalProjectByName(context.TODO(), t.Project, lc)
			if err != nil {
				return err
			}

			if project.Environments != nil {
				if len(project.Environments) > 0 && t.Environment != "" {
					for _, e := range project.Environments {
						if e.Name == t.Environment {
							t.EnvironmentID = int(e.ID)
						}
					}
				}
			}

			// Get current advanced tasks by Environment ID
			advancedTasks, err := lagoon.GetAdvancedTasksByEnvironment(context.TODO(), t.EnvironmentID, lc)
			if err != nil {
				return err
			}

			if advancedTasks != nil {
				var tasksJSON []byte
				if cmdProjectName != "" && fileName == "" {
					fmt.Printf("Found tasks for '%s:%s'\n", t.Project, t.Environment)
					tasksJSON, err = json.MarshalIndent(advancedTasks, "", "  ")
					if err != nil {
						return err
					}

				} else {
					fmt.Printf("Found '%s' for '%s:%s'\n", t.Name, t.Project, t.Environment)
					for _, task := range *advancedTasks {
						if task.Name == t.Name {
							tasksJSON, err = json.Marshal(task)
							if err != nil {
								return err
							}
						}
					}
				}

				resultData := output.Result{
					Result: string(tasksJSON),
				}
				output.RenderResult(resultData, outputOptions)
			}
		}
		return nil
	},
}

// setLastApplied Finds a resource match in the API and updates the configuration with local input
var setLastApplied = &cobra.Command{
	Use:   "set-last-applied -f FILENAME",
	Short: "Set the latest applied workflows or advanced task definitions for project/environment.",
	Long:  `Finds latest configuration match by workflow/task definition 'Name' and sets the latest applied workflow or advanced task definition for project/environment with the contents of file.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(cmdLagoon)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}
		if fileName == "" {
			return fmt.Errorf("no file name given to apply (pass a configuration file as 'apply set-last-applied -f [FILENAME])")
		}

		fileConfig, err := ReadConfigFile(fileName)
		if err != nil {
			log.Fatalln("Parsing config error:", err)
		}

		// Validate tasks input
		_, err = PreprocessAdvancedTaskDefinitionsInputValidation(fileConfig.Tasks)
		if err != nil {
			return err
		}

		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		// Apply Advanced Tasks
		err = ApplyAdvancedTaskDefinitions(lc, fileConfig.Tasks)
		if err != nil {
			return err
		}

		return nil
	},
}

// Applies the latest configuration from a given yaml file. A file is required.
var applyCmd = &cobra.Command{
	Use:     "apply",
	Aliases: []string{"ap"},
	Short:   "Apply the configuration of workflows or tasks from a given yaml configuration file",
	Long: `Apply the configuration of workflows or tasks from a given yaml configuration file.
Workflows or advanced task definitions will be created if they do not already exist.`,
	PreRunE: func(_ *cobra.Command, _ []string) error {
		return validateTokenE(lagoonCLIConfig.Current)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		debug, err := cmd.Flags().GetBool("debug")
		if err != nil {
			return err
		}

		fileName, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}

		if fileName == "" {
			return fmt.Errorf("no file name given to apply (pass a configuration file as 'apply -f [FILENAME])")
		}

		// Read yaml config from file
		fileConfig, err := ReadConfigFile(fileName)
		if err != nil {
			log.Fatalln("Parsing config error:", err)
		}

		// Validate input
		_, err = PreprocessWorkflowsInputValidation(fileConfig.Workflows)
		if err != nil {
			return err
		}

		_, err = PreprocessAdvancedTaskDefinitionsInputValidation(fileConfig.Tasks)
		if err != nil {
			return err
		}

		// Create lagoon client
		current := lagoonCLIConfig.Current
		lc := client.New(
			lagoonCLIConfig.Lagoons[current].GraphQL,
			lagoonCLIConfig.Lagoons[current].Token,
			lagoonCLIConfig.Lagoons[current].Version,
			lagoonCLIVersion,
			debug)

		// Apply Advanced Tasks first since they can be added to workflows.
		err = ApplyAdvancedTaskDefinitions(lc, fileConfig.Tasks)
		if err != nil {
			return err
		}

		// Apply Workflows
		err = ApplyWorkflows(lc, fileConfig.Workflows)
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
	viewLastApplied.InheritedFlags()
	setLastApplied.InheritedFlags()
}
