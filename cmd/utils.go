package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"os"
)

//matchEventToEventType Maps Event types from file input to event types that exist on Lagoon
func matchEventToEventType(event string) schema.EventType {
	switch event {
	case "DeployEnvironmentLatest":
		return schema.APIDeployEnvironmentLatest
	case "DeployEnvironmentBranch":
		return schema.APIDeployEnvironmentBranch
	case "APIDeleteEnvironment":
		return schema.APIDeleteEnvironment

	case "DeployOSFinished":
		return schema.DeployOSFinished
	case "DeployKubernetesFinished":
		return schema.DeployKubernetesFinished
	case "RemoveOSFinished":
		return schema.RemoveOSFinished
	case "RemoveKubernetesFinished":
		return schema.RemoveKubernetesFinished

	case "DeployErrorRemoveKubernetes":
		return schema.DeployErrorRemoveKubernetes
	case "DeployErrorRemoveOS":
		return schema.DeployErrorRemoveOS
	case "DeployErrorBuildDeployKubernetes":
		return schema.DeployErrorBuildDeployKubernetes
	case "DeployErrorBuildDeployOS":
		return schema.DeployErrorBuildDeployOS

	case "GithubPush":
		return schema.GithubPush
	case "GithubPROpened":
		return schema.GithubPROpened
	case "GithubPRUpdated":
		return schema.GithubPRUpdated
	case "GithubPRClosed":
		return schema.GithubPRClosed
	case "GithubPushSkip":
		return schema.GithubPushSkip
	case "GithubDeleteEnvironment":
		return schema.GithubDeleteEnvironment
	case "GithubPRNotDeleted":
		return schema.GithubPRNotDeleted
	case "GithubPushNotDeleted":
		return schema.GithubPushNotDeleted

	case "GitlabPush":
		return schema.GitlabPush
	case "GitlabPushSkip":
		return schema.GitlabPushSkip
	case "GitlabPROpened":
		return schema.GitlabPROpened
	case "GitlabPRUpdated":
		return schema.GitlabPRUpdated
	case "GitlabPRClosed":
		return schema.GitlabPRClosed
	case "GitlabDeleteEnvironment":
		return schema.GitlabDeleteEnvironment
	case "GitlabPushNotDeleted":
		return schema.GitlabPushNotDeleted

	case "BitbucketPush":
		return schema.BitbucketPush
	case "BitbucketPushSkip":
		return schema.BitbucketPushSkip
	case "BitbucketPROpened":
		return schema.BitbucketPROpened
	case "BitbucketPRCreatedOpened":
		return schema.BitbucketPRCreatedOpened
	case "BitbucketPRUpdated":
		return schema.BitbucketPRUpdated
	case "BitbucketPRUpdatedOpened":
		return schema.BitbucketPRUpdatedOpened
	case "BitbucketPRFulfilled":
		return schema.BitbucketPRFulfilled
	case "BitbucketPRRejected":
		return schema.BitbucketPRRejected
	case "BitbucketDeleteEnvironment":
		return schema.BitbucketDeleteEnvironment
	case "BitbucketPushNotDeleted":
		return schema.BitbucketPushNotDeleted

	default:
		return ""
	}
}

// PreprocessWorkflowsInputValidation Validates input from workflow configuration file
func PreprocessWorkflowsInputValidation(workflowsInput []WorkflowsFileInput) ([]WorkflowsFileInput, error) {
	var hasNonProceedableErrors = false

	if len(workflowsInput) > 0 {
		for _, w := range workflowsInput {
			if cmdProjectName != "" {
				w.Project = cmdProjectName
			}
			if cmdProjectEnvironment != "" {
				w.Environment = cmdProjectEnvironment
			}

			if w.Name == "" {
				hasNonProceedableErrors = true
				fmt.Printf("Workflow 'name' is required for '%s'\n\n", w.Name)
			}
			if w.Project == "" {
				hasNonProceedableErrors = true
				fmt.Printf("'project' name is required for workflow '%s'\n\n", w.Name)
			}
			if w.Environment == "" {
				hasNonProceedableErrors = true
				fmt.Printf("An 'environment' name is required for workflow '%s'\n\n", w.Name)
			}
			if w.Event == "" {
				hasNonProceedableErrors = true
				fmt.Printf("'event' name is required for workflow '%s'\n\n", w.Name)
			} else {
				// find a Lagoon event type match
				eventType := matchEventToEventType(w.Event)
				if eventType == "" {
					hasNonProceedableErrors = true
					fmt.Printf("Event '%s' did not match a Lagoon event type \n\n", w.Event)
				}
			}
		}
	}
	if hasNonProceedableErrors {
		return nil, fmt.Errorf("validation checks failed\n")
	}
	return workflowsInput, nil
}

// PreprocessAdvancedTaskDefinitionsInputValidation validates task inputs
func PreprocessAdvancedTaskDefinitionsInputValidation(tasksInput []AdvancedTasksFileInput) ([]AdvancedTasksFileInput, error) {
	var hasNonProceedableErrors = false

	if len(tasksInput) > 0 {
		// Check and collate each task for validation issues
		for _, t := range tasksInput {
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
		return nil, fmt.Errorf("validation checks failed")
	}
	return tasksInput, nil
}

// DiffPatchChangesAgainstAPI Diffs input config from patch against API config.
func DiffPatchChangesAgainstAPI(patchConfig []byte, apiConfig interface{}) (string, error) {
	currAPIJSON, _ := json.Marshal(apiConfig)

	differ := diff.New()
	d, err := differ.Compare(currAPIJSON, patchConfig)
	if err != nil {
		fmt.Printf("Failed to unmarshal file: %s\n", err.Error())
		os.Exit(3)
	}

	if !d.Modified() {
		return "", nil
	}

	var aJSON map[string]interface{}
	json.Unmarshal(currAPIJSON, &aJSON)

	var diffString string
	config := formatter.AsciiFormatterConfig{
		ShowArrayIndex: true,
		Coloring:       true,
	}

	formatter := formatter.NewAsciiFormatter(aJSON, config)
	diffString, err = formatter.Format(d)
	if err != nil {
		fmt.Printf("Failed to diff config: %s\n", err.Error())
		os.Exit(3)
	}

	return diffString, nil
}
