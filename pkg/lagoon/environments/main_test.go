package environments

import (
	"bytes"
	"reflect"
	"testing"
)

func checkEqual(t *testing.T, got, want interface{}, msgs ...interface{}) {
	if !reflect.DeepEqual(got, want) {
		buf := bytes.Buffer{}
		buf.WriteString("got:\n[%v]\nwant:\n[%v]\n")
		for _, v := range msgs {
			buf.WriteString(v.(string))
		}
		t.Errorf(buf.String(), got, want)
	}
}

func TestGetEnvironmentByName(t *testing.T) {
	var all = `{"autoIdle":1,"created":"2019-10-29 04:26:11","deleted":"0000-00-00 00:00:00","deployBaseRef":"Master","deployHeadRef":null,"deployTitle":null,"deployType":"branch","environmentType":"production","id":3,"name":"Master","openshiftProjectName":"high-cotton-master","route":"http://highcotton.org","routes":"http://highcotton.org,https://varnish-highcotton-org-prod.us.amazee.io,https://nginx-highcotton-org-prod.us.amazee.io","updated":"2019-10-29 04:26:43"}`
	var allSuccess = `{"header":["ID","EnvironmentName","EnvironmentType","DeployType","Created","OpenshiftProjectName","Route","Routes","AutoIdle","DeployTitle","DeployBaseRef","DeployHeadRef"],"data":[["3","Master","production","branch","2019-10-29 04:26:11","high-cotton-master","http://highcotton.org","http://highcotton.org,https://varnish-highcotton-org-prod.us.amazee.io,https://nginx-highcotton-org-prod.us.amazee.io","1","-","Master","-"]]}`

	testResult, err := processEnvInfo([]byte(all))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "projectInfo processing failed")
	}
}
