package parser

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

func TestAllProjects(t *testing.T) {
	var allProjects = `[
			{
			  "name": "credentialstest-project1",
			  "autoIdle": 1,
			  "branches": "true",
			  "pullrequests": "true",
			  "privateKey": "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW\nQyNTUxOQAAACCW9R+b8pbIPMvOMntZvlDw+gOGkvSDiedZyGBVFBasjAAAAIh3Gavhdxmr\n4QAAAAtzc2gtZWQyNTUxOQAAACCW9R+b8pbIPMvOMntZvlDw+gOGkvSDiedZyGBVFBasjA\nAAAEDC96enxOrHKiZ9DyRscnVlXxPI6FjEISu1O5J2nf8N/5b1H5vylsg8y84ye1m+UPD6\nA4aS9IOJ51nIYFUUFqyMAAAAAAECAwQF\n-----END OPENSSH PRIVATE KEY-----\n",
			  "productionEnvironment": "master",
			  "activeSystemsDeploy": "lagoon_openshiftBuildDeploy",
			  "activeSystemsTask": "lagoon_openshiftJob",
			  "activeSystemsRemove": "lagoon_openshiftRemove",
			  "activeSystemsPromote": "lagoon_openshiftBuildDeploy",
			  "storageCalc": 1,
			  "openshiftProjectPattern": null,
			  "developmentEnvironmentsLimit": 5,
			  "gitUrl": "ssh://git@192.168.42.1:2222/git/project1.git",
			  "groups": [
				{
				  "name": "credentialtest-group1",
				  "members": [
					{
					  "user": {
						"email": "credentialtestbothgroupaccess_user@example.com",
						"sshKeys": [
						  {
							"name": "credentialtestbothgroupaccess",
							"keyType": "ssh-rsa",
							"keyValue": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDHDRxmCLzLCLdo8M8hK+v5Zr5p4q0XQYwCm9xxWWU8ItkyP4LB90fyh8qWFJRQTYrGNe4usDL5kPyHXhUevdZt75jcjWqlWypbtNY4nFfi+HQ9dnR/f7RRkoBWbha3j8mqZdRHXo9xrXMc82wXaN9a9dMbvEmZPkTpi75g+C0KXBrfjJjDR6Lvr7zWoB2nPMmgy1UGLl5dKxJOg7vHYCpAI582knd0nep3t3bRdlxe7l/CxMthZmoJxz6dfoFotDyBjivVYqEiybtjkNBQnRf2xYQ7m6837hu3EkSVKdgbonWASFMenpuFacDE2S2dVftwW3QjzIOtdQERGZuMsi0p"
						  },
						  {
							"name": "credentialtestsubgroupaccess",
							"keyType": "ssh-rsa",
							"keyValue": "AAAAB3NzaC1yc2EAAAADAQABAAABAQDSeDyxIwMWOjeAq7hF2qirbtoD33JZA2RxodsnnesrmWvrRlANn/fYZSR9zc8SMUMA6s6gR6eOLP+Nrz6bz9xIPFDhHpU7yJqWvb8EM2EPpgSppGmBIBDIPN+5YSxiX5fTuLn2dGHyIdgRvl1yXC09Pa16i/gweMIskP7nanFUfVcgUlTNm0nS6F+MzqklRRO7Tw1zOuleEnXohpzGpUJV5xgQmX4CoqiKYuyXssgBdwsrB9oYdGl8i7xi1w7Xlop6FhbaY/vQdy5f9xfoY2jiO7Big6FDeUGccpjL+xZ0MlKfiTl/OIHcsbshQZ1+D9eSmeaaQ8r9cujgTX2Y7umt"
						  }
						],
						"firstName": null,
						"lastName": null
					  },
					  "role": "OWNER"
					}
				  ]
				},
				{
				  "name": "project-credentialstest-project1",
				  "members": [
					{
					  "user": {
						"email": "default-user@credentialstest-project1",
						"sshKeys": [
						  {
							"name": "auto-add via api",
							"keyType": "ssh-ed25519",
							"keyValue": "AAAAC3NzaC1lZDI1NTE5AAAAIJb1H5vylsg8y84ye1m+UPD6A4aS9IOJ51nIYFUUFqyM"
						  }
						],
						"firstName": null,
						"lastName": null
					  },
					  "role": "MAINTAINER"
					}
				  ]
				}
			  ],
			  "notifications": [],
			  "openshift": {
				"id": 1
			  },
			  "envVariables": [],
			  "environments": []
			}
			]`
	allProjectsSuccess := `groups:
- name: credentialtest-group1
- name: project-credentialstest-project1
notifications: {}
projects:
- groups:
  - credentialtest-group1
  - project-credentialstest-project1
  notifications: {}
  project:
    activeSystemsDeploy: lagoon_openshiftBuildDeploy
    activeSystemsPromote: lagoon_openshiftBuildDeploy
    activeSystemsRemove: lagoon_openshiftRemove
    activeSystemsTask: lagoon_openshiftJob
    autoIdle: 1
    branches: "true"
    developmentEnvironmentsLimit: 5
    gitUrl: ssh://git@192.168.42.1:2222/git/project1.git
    name: credentialstest-project1
    openshift: 0
    privateKey: |
      -----BEGIN OPENSSH PRIVATE KEY-----
      b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzc2gtZW
      QyNTUxOQAAACCW9R+b8pbIPMvOMntZvlDw+gOGkvSDiedZyGBVFBasjAAAAIh3Gavhdxmr
      4QAAAAtzc2gtZWQyNTUxOQAAACCW9R+b8pbIPMvOMntZvlDw+gOGkvSDiedZyGBVFBasjA
      AAAEDC96enxOrHKiZ9DyRscnVlXxPI6FjEISu1O5J2nf8N/5b1H5vylsg8y84ye1m+UPD6
      A4aS9IOJ51nIYFUUFqyMAAAAAAECAwQF
      -----END OPENSSH PRIVATE KEY-----
    productionEnvironment: master
    pullrequests: "true"
    storageCalc: 1
users:
- groups:
  - name: credentialtest-group1
    role: OWNER
  user:
    email: credentialtestbothgroupaccess_user@example.com
    sshkeys:
    - keyname: credentialtestbothgroupaccess
      sshkey: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDHDRxmCLzLCLdo8M8hK+v5Zr5p4q0XQYwCm9xxWWU8ItkyP4LB90fyh8qWFJRQTYrGNe4usDL5kPyHXhUevdZt75jcjWqlWypbtNY4nFfi+HQ9dnR/f7RRkoBWbha3j8mqZdRHXo9xrXMc82wXaN9a9dMbvEmZPkTpi75g+C0KXBrfjJjDR6Lvr7zWoB2nPMmgy1UGLl5dKxJOg7vHYCpAI582knd0nep3t3bRdlxe7l/CxMthZmoJxz6dfoFotDyBjivVYqEiybtjkNBQnRf2xYQ7m6837hu3EkSVKdgbonWASFMenpuFacDE2S2dVftwW3QjzIOtdQERGZuMsi0p
    - keyname: credentialtestsubgroupaccess
      sshkey: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDSeDyxIwMWOjeAq7hF2qirbtoD33JZA2RxodsnnesrmWvrRlANn/fYZSR9zc8SMUMA6s6gR6eOLP+Nrz6bz9xIPFDhHpU7yJqWvb8EM2EPpgSppGmBIBDIPN+5YSxiX5fTuLn2dGHyIdgRvl1yXC09Pa16i/gweMIskP7nanFUfVcgUlTNm0nS6F+MzqklRRO7Tw1zOuleEnXohpzGpUJV5xgQmX4CoqiKYuyXssgBdwsrB9oYdGl8i7xi1w7Xlop6FhbaY/vQdy5f9xfoY2jiO7Big6FDeUGccpjL+xZ0MlKfiTl/OIHcsbshQZ1+D9eSmeaaQ8r9cujgTX2Y7umt
- groups:
  - name: project-credentialstest-project1
    role: MAINTAINER
  user:
    email: default-user@credentialstest-project1
    sshkeys:
    - keyname: auto-add via api
      sshkey: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIJb1H5vylsg8y84ye1m+UPD6A4aS9IOJ51nIYFUUFqyM
`

	skip := SkipExport{
		Users:         false,
		Groups:        false,
		Notifications: false,
		Slack:         false,
		RocketChat:    false,
	}
	returnResult := processParser([]byte(allProjects), skip)
	if string(returnResult) != allProjectsSuccess {
		checkEqual(t, string(returnResult), allProjectsSuccess, "allProject processing failed")
	}
}
