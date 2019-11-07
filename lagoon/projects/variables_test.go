package projects

import (
	"testing"
)

func TestListProjectVariables(t *testing.T) {
	var all = `{
		"envVariables":[
			{"id":1,"name":"TEST_VAR1","value":"value1","scope":"runtime"},
			{"id":3,"name":"TEST_VAR2","value":"value2","scope":"build"},
			{"id":5,"name":"TEST_VAR3","value":"value3","scope":"build"},
			{"id":7,"name":"TEST_VAR5","value":"value5","scope":"build"}
		],"id":18,"name":"high-cotton"
	}`
	var allSuccess = `{"header":["ID","Project","Scope","VariableName"],"data":[["1","high-cotton","runtime","TEST_VAR1"],["3","high-cotton","build","TEST_VAR2"],["5","high-cotton","build","TEST_VAR3"],["7","high-cotton","build","TEST_VAR5"]]}`
	var allSuccess2 = `{"header":["ID","Project","Scope","VariableName","VariableValue"],"data":[["1","high-cotton","runtime","TEST_VAR1","value1"],["3","high-cotton","build","TEST_VAR2","value2"],["5","high-cotton","build","TEST_VAR3","value3"],["7","high-cotton","build","TEST_VAR5","value5"]]}`

	testResult, err := processProjectVariables([]byte(all), false)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "projectInfo processing failed")
	}
	testResult, err = processProjectVariables([]byte(all), true)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess2 {
		checkEqual(t, string(testResult), allSuccess2, "projectInfo processing failed")
	}
}
