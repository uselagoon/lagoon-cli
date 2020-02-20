package environments

import (
	"testing"
)

func TestListEnvironmentTassks(t *testing.T) {
	var all = `{"tasks":[
		{"completed":null,"created":"2019-11-10 22:04:58","id":31,"name":"Drush cache-clear","remoteId":null,"started":null,"status":"failed"},
		{"completed":"2018-10-07 23:13:41","created":"2018-10-07 23:02:41","id":5,"name":"Drupal Archive","remoteId":null,"started":"2018-10-07 23:03:41","status":"succeeded"},
		{"completed":"2018-10-07 23:05:41","created":"2018-10-07 23:02:41","id":6,"name":"Drupal Archive","remoteId":null,"started":"2018-10-07 23:03:41","status":"failed"},
		{"completed":null,"created":"2018-10-07 23:02:41","id":7,"name":"Site Status","remoteId":null,"started":"2018-10-07 23:03:41","status":"active"},
		{"completed":null,"created":"2018-10-07 23:02:41","id":4,"name":"Site Status","remoteId":null,"started":"2018-10-07 23:03:41","status":"active"},
		{"completed":null,"created":"2018-10-07 23:02:41","id":16,"name":"Site Status","remoteId":"600abdae-003f-11ea-bcee-02d6ad974cf2","started":"2018-10-07 23:03:41","status":"active"}
	]}`
	var allSuccess = `{"header":["ID","RemoteID","Name","Status","Created","Started","Completed","Service"],"data":[["31","-","Drush_cache-clear","failed","2019-11-10 22:04:58","-","-","-"],["5","-","Drupal_Archive","succeeded","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:13:41","-"],["6","-","Drupal_Archive","failed","2018-10-07 23:02:41","2018-10-07 23:03:41","2018-10-07 23:05:41","-"],["7","-","Site_Status","active","2018-10-07 23:02:41","2018-10-07 23:03:41","-","-"],["4","-","Site_Status","active","2018-10-07 23:02:41","2018-10-07 23:03:41","-","-"],["16","600abdae-003f-11ea-bcee-02d6ad974cf2","Site_Status","active","2018-10-07 23:02:41","2018-10-07 23:03:41","-","-"]]}`

	testResult, err := processEnvironmentTasks([]byte(all))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "projectInfo processing failed")
	}
}
