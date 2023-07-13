package projects

import (
	"bytes"
	"encoding/json"
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

func TestAllProjects(t *testing.T) {
	var allProjects = `[
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/project1.git","id":1,"name":"credentialstest-project1"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/project2.git","id":2,"name":"credentialstest-project2"},
	{"developmentEnvironmentsLimit":5,"environments":[
		{"environmentType":"production","route":null}],
		"gitUrl":"ssh://git@192.168.99.1:2222/git/github.git","id":3,"name":"ci-github"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/gitlab.git","id":4,"name":"ci-gitlab"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/nginx.git","id":11,"name":"ci-nginx"},
	{"developmentEnvironmentsLimit":5,"environments":[
		{"environmentType":"production","route":null}],
		"gitUrl":"ssh://git@192.168.99.1:2222/git/features.git","id":12,"name":"ci-features"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"git@github.com:uselagoon/lagoon.git","id":13,"name":"lagoon"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/features-subfolder.git","id":17,"name":"ci-features-subfolder"},
	{"developmentEnvironmentsLimit":5,"environments":[
		{"environmentType":"production","route":"http://highcotton.org"},
		{"environmentType":"development","route":"https://varnish-highcotton-org-staging.us.amazee.io"},
		{"environmentType":"development","route":"https://varnish-highcotton-org-development.us.amazee.io"},
		{"environmentType":"development","route":""},
		{"environmentType":"development","route":null}],
		"gitUrl":"test","id":18,"name":"high-cotton"},
	{"developmentEnvironmentsLimit":5,"environments":[],"gitUrl":"ssh://git@192.168.99.1:2222/git/api.git","id":21,"name":"ci-api"}
]`
	var allProjectsSuccess = `{"header":["ID","ProjectName","GitURL","ProductionEnvironment","DevEnvironments"],"data":[["1","credentialstest-project1","ssh://git@192.168.99.1:2222/git/project1.git","","0/5"],["2","credentialstest-project2","ssh://git@192.168.99.1:2222/git/project2.git","","0/5"],["3","ci-github","ssh://git@192.168.99.1:2222/git/github.git","","0/5"],["4","ci-gitlab","ssh://git@192.168.99.1:2222/git/gitlab.git","","0/5"],["11","ci-nginx","ssh://git@192.168.99.1:2222/git/nginx.git","","0/5"],["12","ci-features","ssh://git@192.168.99.1:2222/git/features.git","","0/5"],["13","lagoon","git@github.com:uselagoon/lagoon.git","","0/5"],["17","ci-features-subfolder","ssh://git@192.168.99.1:2222/git/features-subfolder.git","","0/5"],["18","high-cotton","test","","4/5"],["21","ci-api","ssh://git@192.168.99.1:2222/git/api.git","","0/5"]]}`

	returnResult, err := processAllProjects([]byte(allProjects))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != allProjectsSuccess {
		checkEqual(t, string(returnResult), allProjectsSuccess, "allProject processing failed")
	}
}

func TestProjectInfo(t *testing.T) {
	var projectInfo = `
	{"autoIdle":1,"branches":"false","developmentEnvironmentsLimit":5,"factsUI":0,"problemsUI":0,"environments":[
		{"deployType":"branch","environmentType":"production","id":3,"name":"Master","openshiftProjectName":"high-cotton-master","route":"http://highcotton.org"},
		{"deployType":"branch","environmentType":"development","id":4,"name":"Staging","openshiftProjectName":"high-cotton-staging","route":"https://varnish-highcotton-org-staging.us.amazee.io"},
		{"deployType":"branch","environmentType":"development","id":5,"name":"Development","openshiftProjectName":"high-cotton-development","route":"https://varnish-highcotton-org-development.us.amazee.io"},
		{"deployType":"pullrequest","environmentType":"development","id":6,"name":"PR-175","openshiftProjectName":"high-cotton-pr-175","route":""},
		{"deployType":"branch","environmentType":"development","id":10,"name":"high-cotton","openshiftProjectName":"high-cotton-high-cotton","route":null}],
		"gitUrl":"test","id":18,"name":"high-cotton","productionEnvironment":"doopdd","routerPattern":"${environment}-${project}.lagoon.example.com","pullrequests":"false","storageCalc":1,"subfolder":null}`
	var projectInfoSuccess = `{"header":["ID","ProjectName","GitURL","Branches","PullRequests","ProductionRoute","DevEnvironments","DevEnvLimit","ProductionEnv","RouterPattern","AutoIdle","FactsUI","ProblemsUI"],"data":[["18","high-cotton","test","false","false","http://highcotton.org","4/5","5","doopdd","${environment}-${project}.lagoon.example.com","1","0","0"]]}`

	returnResult, err := processProjectInfo([]byte(projectInfo))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != projectInfoSuccess {
		checkEqual(t, string(returnResult), projectInfoSuccess, "projectInfo processing failed")
	}
}

func TestProjectUpdate(t *testing.T) {
	var updateProject = `
	{"autoIdle":1,"branches":"true","developmentEnvironmentsLimit":5,"environments":[
		{"deployType":"branch","environmentType":"production","id":3,"name":"Master","openshiftProjectName":"high-cotton-master","route":"http://highcotton.org"},
		{"deployType":"branch","environmentType":"development","id":4,"name":"Staging","openshiftProjectName":"high-cotton-staging","route":"https://varnish-highcotton-org-staging.us.amazee.io"},
		{"deployType":"branch","environmentType":"development","id":5,"name":"Development","openshiftProjectName":"high-cotton-development","route":"https://varnish-highcotton-org-development.us.amazee.io"},
		{"deployType":"pullrequest","environmentType":"development","id":6,"name":"PR-175","openshiftProjectName":"high-cotton-pr-175","route":""},
		{"deployType":"branch","environmentType":"development","id":10,"name":"high-cotton","openshiftProjectName":"high-cotton-high-cotton","route":null}],
		"gitUrl":"test","id":18,"name":"high-cotton","productionEnvironment":"Master","routerPattern":"${environment}-${project}.lagoon.example.com","pullrequests":"true","storageCalc":1,"subfolder":null}`

	var jsonPatch = `{"branches":"false"}`
	var updateProjectsSuccess = `{"id":18,"patch":{"branches":"false"}}`

	var jsonPatchFail = `{"branchesd":"false", "developmentEnvironmentsLimit": 10}`
	var updateProjectsFail = `{"id":18,"patch":{"developmentEnvironmentsLimit":10}}`

	returnResult, err := processProjectUpdate([]byte(updateProject), jsonPatch)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	updateResults, err := json.Marshal(returnResult)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(updateResults) != updateProjectsSuccess {
		checkEqual(t, string(updateResults), updateProjectsSuccess, "projectInfo processing failed")
	}
	returnResult, err = processProjectUpdate([]byte(updateProject), jsonPatchFail)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	updateResults, err = json.Marshal(returnResult)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(updateResults) == updateProjectsFail {
		checkEqual(t, string(updateResults), updateProjectsFail, "projectInfo processing succeeded, but should fail if json patch is broken")
	}
}
