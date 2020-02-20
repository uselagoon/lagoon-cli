package projects

import (
	"testing"
)

func TestListProjectRocketChats(t *testing.T) {
	var allRocketChats = `{"rocketchats":[{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]}`
	var allRocketChatsSuccess = `{"header":["NID","NotificationName","Channel","Webhook"],"data":[["1","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`

	returnResult, err := processProjectRocketChats([]byte(allRocketChats))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != allRocketChatsSuccess {
		checkEqual(t, string(returnResult), allRocketChatsSuccess, "projectInfo processing failed")
	}
}

func TestListAllRocketChats(t *testing.T) {
	var allRocketChats = `[
		{"id":1,"name":"credentialstest-project1","notifications":[]},
		{"id":2,"name":"credentialstest-project2","notifications":[]},
		{"id":3,"name":"ci-github","notifications":[
			{},{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":4,"name":"ci-gitlab","notifications":[
			{},{},{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":5,"name":"ci-bitbucket","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":6,"name":"ci-rest","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":7,"name":"ci-node","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":8,"name":"ci-multiproject1","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":9,"name":"ci-multiproject2","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":10,"name":"ci-drupal","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":11,"name":"ci-nginx","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":12,"name":"ci-features","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":13,"name":"lagoon","notifications":[
			{"channel":"lagoon-kickstart","id":3,"name":"amazeeio--lagoon-kickstart","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":14,"name":"ci-elasticsearch","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":15,"name":"ci-drupal-galera","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":16,"name":"ci-drupal-postgres","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":17,"name":"ci-features-subfolder","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":18,"name":"high-cotton","notifications":[{},{}]},
		{"id":19,"name":"ci-env-limit","notifications":[]},
		{"id":20,"name":"ci-solr","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":21,"name":"ci-api","notifications":[
			{"channel":"lagoon-local-ci","id":1,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]},
		{"id":22,"name":"credentialstest-project3","notifications":[]}]`
	var allRocketChatsSuccess = `{"header":["NID","Project","NotificationName","Channel","Webhook"],"data":[["1","ci-github","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-gitlab","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-bitbucket","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-rest","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-node","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-multiproject1","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-multiproject2","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-drupal","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-nginx","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-features","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["3","lagoon","amazeeio--lagoon-kickstart","lagoon-kickstart","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-elasticsearch","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-drupal-galera","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-drupal-postgres","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-features-subfolder","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-solr","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"],["1","ci-api","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.rocket.chat/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`

	returnResult, err := processAllSlacks([]byte(allRocketChats))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != allRocketChatsSuccess {
		checkEqual(t, string(returnResult), allRocketChatsSuccess, "projectInfo processing failed")
	}
}

func TestListProjectSlacks(t *testing.T) {
	var allSlacks = `{"slacks":[{"channel":"lagoon-local-ci","id":30,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.slack.fake/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"}]}`
	var allSlacksSuccess = `{"header":["NID","NotificationName","Channel","Webhook"],"data":[["30","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.slack.fake/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`

	returnResult, err := processProjectSlacks([]byte(allSlacks))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != allSlacksSuccess {
		checkEqual(t, string(returnResult), allSlacksSuccess, "projectInfo processing failed")
	}
}

func TestListAllSlacks(t *testing.T) {
	var allSlacks = `[
		{"id":1,"name":"credentialstest-project1","notifications":[]},
		{"id":2,"name":"credentialstest-project2","notifications":[]},
		{"id":3,"name":"ci-github","notifications":[
			{"channel":"lagoon-local-ci","id":30,"name":"amazeeio--lagoon-local-ci","webhook":"https://amazeeio.slack.fake/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"},
			{}]},
		{"id":4,"name":"ci-gitlab","notifications":[]},
		{"id":5,"name":"ci-bitbucket","notifications":[{}]},
		{"id":6,"name":"ci-rest","notifications":[{}]},
		{"id":7,"name":"ci-node","notifications":[{}]},
		{"id":8,"name":"ci-multiproject1","notifications":[{}]},
		{"id":9,"name":"ci-multiproject2","notifications":[{}]},
		{"id":10,"name":"ci-drupal","notifications":[{}]},
		{"id":11,"name":"ci-nginx","notifications":[{}]},
		{"id":12,"name":"ci-features","notifications":[{}]},
		{"id":13,"name":"lagoon","notifications":[{}]},
		{"id":14,"name":"ci-elasticsearch","notifications":[{}]},
		{"id":15,"name":"ci-drupal-galera","notifications":[{}]},
		{"id":16,"name":"ci-drupal-postgres","notifications":[{}]},
		{"id":17,"name":"ci-features-subfolder","notifications":[{}]},
		{"id":18,"name":"high-cotton","notifications":[]},
		{"id":19,"name":"ci-env-limit","notifications":[]},
		{"id":20,"name":"ci-solr","notifications":[{}]},
		{"id":21,"name":"ci-api","notifications":[{}]},
		{"id":22,"name":"credentialstest-project3","notifications":[]}]`
	var allSlacksSuccess = `{"header":["NID","Project","NotificationName","Channel","Webhook"],"data":[["30","ci-github","amazeeio--lagoon-local-ci","lagoon-local-ci","https://amazeeio.slack.fake/hooks/ikF5XMohDZK7KpsZf/c9BFBt2ch8oMMuycoERJQMSLTPo8nmZhg2Hf2ny68ZpuD4Kn"]]}`

	returnResult, err := processAllSlacks([]byte(allSlacks))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(returnResult) != allSlacksSuccess {
		checkEqual(t, string(returnResult), allSlacksSuccess, "projectInfo processing failed")
	}
}
