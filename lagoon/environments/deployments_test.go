package environments

import (
	"testing"
)

func TestGetEnvironmentDeployments(t *testing.T) {
	var testData = `{
		"deployments":[
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":14,"name":"build-2","remoteId":null,"started":"2018-10-07 23:03:41","status":"failed"},
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":1,"name":"build-1","remoteId":null,"started":"2018-10-07 23:03:41","status":"complete"},
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":4,"name":"build-1","remoteId":null,"started":"2018-10-07 23:03:41","status":"complete"},
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":5,"name":"build-2","remoteId":null,"started":"2018-10-07 23:03:41","status":"failed"},
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":7,"name":"build-1","remoteId":null,"started":"2018-10-07 23:03:41","status":"complete"},
			{"completed":"2018-10-07 23:20:41","created":"2018-10-07 23:02:41","id":8,"name":"build-2","remoteId":null,"started":"2018-10-07 23:03:41","status":"failed"}
		]
	}`
	var testSuccess = `{"header":["RemoteID","Name","Status","Created","Started","Completed"],"data":[["","build-2","failed","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"],["","build-1","complete","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"],["","build-1","complete","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"],["","build-2","failed","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"],["","build-1","complete","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"],["","build-2","failed","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:20:41"]]}`

	testResult, err := processEnvironmentDeployments([]byte(testData))
	if err != nil {
		t.Error("Should not fail if processing succeeded, error was:", err)
	}
	if string(testResult) != testSuccess {
		checkEqual(t, string(testResult), testSuccess, "projectInfo processing failed")
	}
}
