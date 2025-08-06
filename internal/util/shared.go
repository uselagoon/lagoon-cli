package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/uselagoon/machinery/api/lagoon/client"
	"github.com/uselagoon/machinery/api/schema"
	"reflect"
	"regexp"
	"slices"
	"strconv"
)

type CreateConfig struct {
	Input               schema.AddProjectInput
	OrganizationDetails schema.Organization
	AutoIdle            bool
	AutoIdleProvided    bool
	StorageCalc         bool
	StorageCalcProvided bool
	DevEnvLimit         uint
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

func IsValidGitURL(url string) bool {
	const pattern = `^((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?$`

	re := regexp.MustCompile(pattern)
	return re.MatchString(url)
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

func GenerateCLICommand(config *CreateConfig) string {
	cmd := ""
	configFields := config.Input
	configType := reflect.TypeOf(configFields)
	configValue := reflect.ValueOf(configFields)
	boolFields := []string{
		"--auto-idle",
		"--storage-calc",
	}

	for i := 0; i < configType.NumField(); i++ {
		fieldName := configType.Field(i).Name
		fieldValue := configValue.Field(i)

		if !fieldValue.IsZero() {
			if flag, exists := fieldCmdMap[fieldName]; exists {
				if flag == "--owner" {
					cmd += " " + fmt.Sprintf("%s=%v", flag, *(fieldValue.Interface().(*bool)))
				} else {
					cmd += " " + fmt.Sprintf("%s %v", flag, fieldValue.Interface())
				}
			}
		}
	}

	configBoolType := reflect.TypeOf(*config)
	configBoolValue := reflect.ValueOf(*config)

	for i := 0; i < configBoolType.NumField(); i++ {
		fieldName := configBoolType.Field(i).Name
		fieldValue := configBoolValue.Field(i)

		var fieldProvided bool

		switch fieldName {
		case "AutoIdle":
			fieldProvided = config.AutoIdleProvided
		case "StorageCalc":
			fieldProvided = config.StorageCalcProvided
		}

		if fieldProvided {
			if flag, exists := fieldCmdMap[fieldName]; exists {
				if slices.Contains(boolFields, flag) {
					boolVal := fieldValue.Interface().(bool)
					cmd += " " + fmt.Sprintf("%s=%v", flag, boolVal)
				}
			}
		}
	}

	return cmd
}
