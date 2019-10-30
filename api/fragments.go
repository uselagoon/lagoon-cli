package api

var deploymentFragment = `fragment Deployment on Deployment {
	id
	name
	status
	created
	started
	completed
	remoteId
	environment {
		name
	}
}`

var taskFragment = `fragment Task on Task {
	id
	name
	status
	created
	started
	completed
	remoteId
	environment {
		name
	}
}`

var projectFragment = `fragment Project on Project {
	id
	name
	gitUrl
	privateKey
}`

var restoreFragment = `fragment Restore on Restore {
    id
    status
    created
    restoreLocation
    backupId
}`

var backupFragment = `fragment Backup on Backup {
	id
	environment {
		id
	}
	backupId
	source
	created
}`

var groupFragment = `fragment Group on Group {
	id
	name
}`

var userFragment = `fragment User on User {
	id
	email
	firstName
	lastName
	gitlabId
	sshKeys {
		id
		name
	}
}`

var sshKeyFragment = `fragment SshKey on SshKey {
	id
	name
	keyValue
	keyType
}`

var notificationsRocketChatFragment = `fragment Notification on Notification {
	webhook
	channel
}`

var notificationsSlackFragment = `fragment Notification on Notification {
	webhook
	channel
}`
