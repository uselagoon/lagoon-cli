package environments

import (
	"encoding/json"
	"testing"

	"github.com/uselagoon/lagoon-cli/pkg/api"
)

func TestListEnvironmentVariables(t *testing.T) {
	var all = `{
		"envVariables":[
			{"id":10,"name":"TEST_VAR10","value":"value10","scope":"runtime"},
			{"id":13,"name":"TEST_VAR20","value":"value20","scope":"build"},
			{"id":15,"name":"TEST_VAR30","value":"value30","scope":"build"},
			{"id":17,"name":"TEST_VAR50","value":"value50","scope":"build"}
		],"name":"master","openshiftProjectName":"high-cotton-master"}`
	var allSuccess = `{"header":["ID","Project","Environment","Scope","VariableName"],"data":[["10","high-cotton","master","runtime","TEST_VAR10"],["13","high-cotton","master","build","TEST_VAR20"],["15","high-cotton","master","build","TEST_VAR30"],["17","high-cotton","master","build","TEST_VAR50"]]}`
	var allSuccess2 = `{"header":["ID","Project","Environment","Scope","VariableName","VariableValue"],"data":[["10","high-cotton","master","runtime","TEST_VAR10","value10"],["13","high-cotton","master","build","TEST_VAR20","value20"],["15","high-cotton","master","build","TEST_VAR30","value30"],["17","high-cotton","master","build","TEST_VAR50","value50"]]}`

	env := api.Environment{}
	json.Unmarshal([]byte(all), &env)
	testResult, err := processEnvironmentVariables(env, "high-cotton", false)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "projectInfo processing failed")
	}
	testResult, err = processEnvironmentVariables(env, "high-cotton", true)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess2 {
		checkEqual(t, string(testResult), allSuccess2, "projectInfo processing failed")
	}
}
