package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
)

type CreateConfig struct {
	Input               schema.AddProjectInput
	OrganizationDetails schema.Organization
}

var fieldCmdMap = map[string]string{
	"Name":                         "--project",
	"GitURL":                       "--git-url",
	"Subfolder":                    "--subfolder",
	"Openshift":                    "--deploytarget",
	"OpenshiftProjectPattern":      "--deploytarget-project-pattern",
	"Branches":                     "--branches",
	"PullRequests":                 "--pullrequests",
	"ProductionEnvironment":        "--production-environment",
	"StandbyProductionEnvironment": "--standby-production-environment",
	"Availability":                 "--availability",
	"AutoIdle":                     "--auto-idle",
	"StorageCalc":                  "--storage-calc",
	"DevelopmentEnvironmentsLimit": "--development-environments-limit",
	"PrivateKey":                   "--private-key",
	"BuildImage":                   "--build-image",
	"Organization":                 "--organization-id",
	"AddOrgOwner":                  "--owner",
	"RouterPattern":                "--router-pattern",
	"ProblemsUI":                   "--problems-ui",
	"FactsUI":                      "--facts-ui",
	"ProductionBuildPriority":      "--production-build-priority",
	"DevelopmentBuildPriority":     "--development-build-priority",
	"DeploymentsDisabled":          "--deployments-disabled",
}

type reflectFields struct {
	fieldType  reflect.Type
	fieldValue reflect.Value
}

func IsValidGitURL(gitUrl string) bool {
	if strings.TrimSpace(gitUrl) == "" {
		return false
	}

	const sshPattern = `^[\w.-]+@[\w.-]+:[\w.-]+/[\w.-]+(?:\.git)?$`
	re := regexp.MustCompile(sshPattern)
	if re.MatchString(gitUrl) {
		return true
	}

	parsedUrl, err := url.Parse(gitUrl)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return false
	}

	if parsedUrl.Host == "" || parsedUrl.Scheme == "" {
		return false
	}

	validProtocols := map[string]bool{"git": true, "ssh": true, "http": true, "https": true}
	if !validProtocols[strings.ToLower(parsedUrl.Scheme)] {
		return false
	}

	pathParts := strings.Split(strings.Trim(parsedUrl.Path, "/"), "/")
	return len(pathParts) >= 2
}

func QuotaCheck(quota int) string {
	quotaRoute := strconv.Itoa(quota)
	if quota < 0 {
		quotaRoute = "∞"
	}
	return quotaRoute
}

func IsValidProjectName(name string) (bool, string) {
	casePattern := regexp.MustCompile("[^0-9a-z-]")
	dashPattern := regexp.MustCompile("--")

	if name == "" {
		return false, "Project name is required"
	}
	if casePattern.MatchString(name) {
		return false, "Project name is invalid, only lowercase characters, numbers and dashes allowed for name"
	}
	if dashPattern.MatchString(name) {
		return false, "Multiple consecutive dashes are not allowed for name"
	}
	return true, ""
}

func GetOrgDeployTargets(lc *client.Client, orgName string) ([]schema.DeployTarget, error) {
	var orgDeployTargets []schema.DeployTarget
	rawOrgByName := `query organizationByNameWithDeployTargets($name: String!) {
			organizationByName(name: $name) {
				id
				name
				deployTargets{
				  id
				  name
				}
			}
		}`

	orgResp, err := lc.ProcessRaw(context.TODO(), rawOrgByName, map[string]interface{}{
		"name": orgName,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println()

	orgData, ok := orgResp.(map[string]interface{})["organizationByName"]
	if !ok {
		return nil, errors.New("organization not found")
	}

	orgDT, exists := orgData.(map[string]interface{})["deployTargets"]
	if !exists {
		return nil, errors.New("no deployTargets found in organization")
	}

	odt, err := json.Marshal(orgDT)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(odt, &orgDeployTargets)
	if err != nil {
		return nil, err
	}
	return orgDeployTargets, nil
}

func processFields(fieldName string, fieldValue reflect.Value) string {
	cmd := ""

	if fieldName == "OrganizationDetails" {
		addOrgInputField := fieldValue.FieldByName("AddOrganizationInput")
		nameField := addOrgInputField.FieldByName("Name")
		if nameField.IsValid() {
			cmd += " " + fmt.Sprintf("%s=%v", "--organization-name", nameField.Interface())
		}
	}

	if flag, exists := fieldCmdMap[fieldName]; exists {
		if flag == "--owner" {
			cmd += " " + fmt.Sprintf("%s=%v", flag, *(fieldValue.Interface().(*bool)))
		} else {
			cmd += " " + fmt.Sprintf("%s %v", flag, fieldValue.Interface())
		}
	}
	return cmd
}

func GenerateCLICommand(config *CreateConfig) string {
	commands := ""

	configFields := []reflectFields{
		{
			fieldType:  reflect.TypeOf(config.Input),
			fieldValue: reflect.ValueOf(config.Input),
		},
		{
			fieldType:  reflect.TypeOf(*config),
			fieldValue: reflect.ValueOf(*config),
		},
	}

	for _, field := range configFields {
		for i := 0; i < field.fieldType.NumField(); i++ {
			fieldName := field.fieldType.Field(i).Name
			fieldValue := field.fieldValue.Field(i)
			if !fieldValue.IsZero() {
				commands += processFields(fieldName, fieldValue)
			}
		}
	}
	return commands
}
