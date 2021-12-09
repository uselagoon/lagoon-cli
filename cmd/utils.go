package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"log"
	"os"
)

type AdvancedTasksInput struct {
	ID          uint                              `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string                            `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                            `json:"description,omitempty" yaml:"description,omitempty"`
	Type        schema.AdvancedTaskDefinitionType `json:"type,omitempty" yaml:"type,omitempty"`
	Command     string                            `json:"command,omitempty" yaml:"command,omitempty"`
	Image       string                            `json:"image,omitempty" yaml:"image,omitempty"`
	Service     string                            `json:"service,omitempty" yaml:"service,omitempty"`
	GroupName   string                            `json:"group,omitempty" yaml:"group,omitempty"`
	Project     string                            `json:"project,omitempty" yaml:"project,omitempty"`
	Environment string                            `json:"environment,omitempty" yaml:"environment,omitempty"`
	Permission  string                            `json:"permission,omitempty" yaml:"permission,omitempty"`
	//AdvancedTaskDefinitionArgument []AdvancedTaskDefinitionArgument `yaml:"arguments,omitempty"`
}

func PreprocessAdvancedTaskDefinitionsInputValidation(tasksInput []byte) {
	var tasks []AdvancedTasksInput
	err := json.Unmarshal(tasksInput, &tasks)
	if err != nil {
		log.Fatalln(err)
	}

	var hasNonProceedableErrors = false

	if len(tasks) > 0 {
		// Check and collate each task for validation issues
		for _, t := range tasks {
			// If project or environment arguments are given, use those.
			if cmdProjectName != "" {
				t.Project = cmdProjectName
			}
			if cmdProjectEnvironment != "" {
				t.Environment = cmdProjectEnvironment
			}

			// Required input checks
			if t.Name == "" {
				hasNonProceedableErrors = true
				fmt.Println("Task name is required")
			}
			if t.Project == "" {
				hasNonProceedableErrors = true
				fmt.Printf("Project name is required for task '%s'\n\n", t.Name)
			}
			if t.Environment == "" {
				hasNonProceedableErrors = true
				fmt.Printf("An Environment name is required for task '%s'\n\n", t.Name)
			}
			if t.Permission == "" {
				hasNonProceedableErrors = true
				fmt.Printf("Permission is required for task '%s'\n", t.Name)
			}
		}
	}

	if hasNonProceedableErrors {
		os.Exit(1)
	}
}
