package users

import (
	"testing"
)

func TestListUsers(t *testing.T) {
	var userList = `[{"id":"21ab7da7-4dc7-4745-92ef-a9faf663b8a4","members":[],"name":"High Cotton Billing Group"},{"id":"07b3d263-3a3e-4d0e-ac3f-56eab1e80df9","members":[{"role":"OWNER","user":{"email":"ci-customer-user-ed25519@example.com","firstName":null,"id":"23781ccd-8e35-4206-a7cf-97153311ba91","lastName":null}},{"role":"OWNER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"ci-group"},{"id":"f6786ff8-4eea-476a-8462-25931c0b2724","members":[{"role":"OWNER","user":{"email":"credentialtestbothgroupaccess_user@example.com","firstName":null,"id":"678707fd-0d01-458d-981f-acae396624bb","lastName":null}}],"name":"credentialtest-group1"},{"id":"8349ffb3-d940-445e-a610-4bb6f3ba8a0f","members":[{"role":"OWNER","user":{"email":"credentialtestbothgroupaccess_user@example.com","firstName":null,"id":"678707fd-0d01-458d-981f-acae396624bb","lastName":null}}],"name":"credentialtest-group2"},{"id":"ec16cf3e-47a9-4aed-825a-5ac5b39acb3d","members":[],"name":"kickstart-group"},{"id":"94311670-e817-4335-b440-25a92e1ac83f","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-api"},{"id":"2d3968f0-36f0-4082-9f02-1dbf86ce410e","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-bitbucket"},{"id":"6ede38ac-b54d-43c2-a4ad-a949eb05edc6","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-drupal"},{"id":"64bd6e32-b3f2-48d5-8a24-bb553d397891","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-drupal-galera"},{"id":"6ef3300b-0e5a-4acc-9b83-4848b295dba4","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-drupal-postgres"},{"id":"aa5d4e8a-c330-4208-bf9a-b9cb3fda7373","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-elasticsearch"},{"id":"a21de1e1-fb3e-4a8c-adcb-f05b6a5b3d98","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-env-limit"},{"id":"60f0103a-a4ea-4dad-b5c7-da786ffa9db8","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-features"},{"id":"68e35dd0-5b0f-4284-8559-065e87f5ce06","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-features-subfolder"},{"id":"c5ca68de-fb5a-4b5e-9005-dd737132acad","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-github"},{"id":"fb1e5d7b-0326-4c81-a1d3-20bbc7de2dc3","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-gitlab"},{"id":"cc5cbaf6-0fc6-40e7-9e07-ff974a65979d","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-multiproject1"},{"id":"54c83c15-1359-44f9-b72f-40401f4f9ae9","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-multiproject2"},{"id":"8bcc7d64-76cb-4ed3-ba14-ce28db8655c8","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-nginx"},{"id":"a9e5c8bb-e201-408f-b7b9-e3dacd648461","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-node"},{"id":"190fc14f-76e4-453f-aefe-473dc4318892","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-rest"},{"id":"ba8abffe-2cea-4315-b788-081a51017290","members":[{"role":"MAINTAINER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null}}],"name":"project-ci-solr"},{"id":"13d15964-4bc3-4e05-aa91-b4ec8a07909a","members":[{"role":"MAINTAINER","user":{"email":"default-user@credentialstest-project1","firstName":null,"id":"f483d9e9-0f1d-4700-bdeb-80b62f89451f","lastName":null}}],"name":"project-credentialstest-project1"},{"id":"fbca7521-aab5-4f60-98e9-8e299439b3e5","members":[{"role":"MAINTAINER","user":{"email":"default-user@credentialstest-project2","firstName":null,"id":"250c1384-c41c-475e-ac7f-5c36d8defc10","lastName":null}}],"name":"project-credentialstest-project2"},{"id":"061fa7d7-5f5c-4ff5-8492-cc99f2250b37","members":[{"role":"MAINTAINER","user":{"email":"default-user@credentialstest-project3","firstName":null,"id":"48947347-ed80-42d4-b070-6c80a3463ec3","lastName":null}}],"name":"project-credentialstest-project3"},{"id":"41e54119-65ba-463f-b2bf-7b5e6575bd27","members":[{"role":"MAINTAINER","user":{"email":"default-user@high-cotton","firstName":null,"id":"ca1ed845-6200-4456-912a-6c3dc162448a","lastName":null}}],"name":"project-high-cotton"},{"id":"cc24ba9b-a39b-48dc-9602-f945d0c2ec69","members":[{"role":"MAINTAINER","user":{"email":"default-user@lagoon","firstName":null,"id":"840671ce-fe85-4d60-ba0a-bb63502955a8","lastName":null}}],"name":"project-lagoon"},{"id":"8140f065-a606-4827-a454-2c9bc4513975","members":[],"name":"ui-customer"}]`
	var allSuccess = `{"header":["ID","Name","FirstName","LastName","Group","Role"],"data":[["23781ccd-8e35-4206-a7cf-97153311ba91","ci-customer-user-ed25519@example.com","-","-","ci-group","OWNER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","ci-group","OWNER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-api","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-bitbucket","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-drupal","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-drupal-galera","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-drupal-postgres","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-elasticsearch","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-env-limit","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-features","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-features-subfolder","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-github","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-gitlab","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-multiproject1","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-multiproject2","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-nginx","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-node","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-rest","MAINTAINER"],["906391b3-b3b2-43cb-82e7-fa0c07007fbc","ci-customer-user-rsa@example.com","-","-","project-ci-solr","MAINTAINER"],["678707fd-0d01-458d-981f-acae396624bb","credentialtestbothgroupaccess_user@example.com","-","-","credentialtest-group1","OWNER"],["678707fd-0d01-458d-981f-acae396624bb","credentialtestbothgroupaccess_user@example.com","-","-","credentialtest-group2","OWNER"],["f483d9e9-0f1d-4700-bdeb-80b62f89451f","default-user@credentialstest-project1","-","-","project-credentialstest-project1","MAINTAINER"],["250c1384-c41c-475e-ac7f-5c36d8defc10","default-user@credentialstest-project2","-","-","project-credentialstest-project2","MAINTAINER"],["48947347-ed80-42d4-b070-6c80a3463ec3","default-user@credentialstest-project3","-","-","project-credentialstest-project3","MAINTAINER"],["ca1ed845-6200-4456-912a-6c3dc162448a","default-user@high-cotton","-","-","project-high-cotton","MAINTAINER"],["840671ce-fe85-4d60-ba0a-bb63502955a8","default-user@lagoon","-","-","project-lagoon","MAINTAINER"]]}`

	testResult, err := processUserList([]byte(userList))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "list user processing failed")
	}
}
func TestListUserKeys(t *testing.T) {
	var userList = `[{"id":"07b3d263-3a3e-4d0e-ac3f-56eab1e80df9","members":[{"role":"MAINTAINER","user":{"email":"c@c.com","firstName":"bob","id":"a08f166e-64f0-4461-9daa-e62b6f414faf","lastName":"bob","sshKeys":[{"keyType":"ssh-rsa","keyValue":"AAAAB3NzaC1yc2EAAAADAQABAAACAQC++bRFdPP6d3kdXv1eImtfSgHhumcsy4IhAYId23v85nmcnTMqA5ahCoOzChPuxKWVsTGaU3xh+PMkQAO/HAkyUYIK8VlIMP/w9+VYraOYHwCZBZqiKwJH6XjpX24qTzdNYU8WdC+6OdOV+0SrEdtduxJV/TnVkoe+Ga+y7013mxCbw5y9LIPs/eBXjwuN/lASxaZGpzAP5FipKC/HOC+oS96gaNVQgmWl3Lm+faGpq2V1afEE9A0RXYhvQ6qG9qvcmFGtqLyhu6m7i9tjNiTlXalsFeox0pu8cnFvKZZZnFwi1EI4ngBUUg/hTFWmr13F2TslEJrnCqm8efs6o40l3tsdD64Jr1q9LUvqKWrwrv4B2vM2O3iccaE5Ll7fNH4pRDTZpFRL92MoX+TBpPjLxFYhz5zOJUKiFMsERkHJB/28a2BPU2etThwkIy5EwOrwHl/Q2KMxrddwwfd9FAZnmHXoUA0OtcZtgyrBDECOzuleGSyhcbXynQBwiEJ1RRQrbeWBTberQi4v+rDKgauxxVyfxH4yQfkoAFwt+QjQPUWv/kKCS+8LEJv1QlEd0lNh2A+TO/ugmsdshO+PzYKUtm4wMiqo0XfzyJGJFvYKbUbPNSZ7iGPRKkwQot8UQgAY/jwXn6z1sm8VmwDLirP1IUIHdt8pGarTffufPd9Rww==","name":"deploy@nhmrc"}]}},{"role":"OWNER","user":{"email":"ci-customer-user-ed25519@example.com","firstName":null,"id":"23781ccd-8e35-4206-a7cf-97153311ba91","lastName":null,"sshKeys":[{"keyType":"ssh-ed25519","keyValue":"AAAAC3NzaC1lZDI1NTE5AAAAIMdEs1h19jv2UrbtKcqPDatUxT9lPYcbGlEAbInsY8Ka","name":"ci-customer-sshkey-ed25519"}]}},{"role":"OWNER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null,"sshKeys":[{"keyType":"ssh-rsa","keyValue":"AAAAB3NzaC1yc2EAAAADAQABAAACAQDEZlms5XsiyWjmnnUyhpt93VgHypse9Bl8kNkmZJTiM3Ex/wZAfwogzqd2LrTEiIOWSH1HnQazR+Cc9oHCmMyNxRrLkS/MEl0yZ38Q+GDfn37h/llCIZNVoHlSgYkqD0MQrhfGL5AulDUKIle93dA6qdCUlnZZjDPiR0vEXR36xGuX7QYAhK30aD2SrrBruTtFGvj87IP/0OEOvUZe8dcU9G/pCoqrTzgKqJRpqs/s5xtkqLkTIyR/SzzplO21A+pCKNax6csDDq3snS8zfx6iM8MwVfh8nvBW9seax1zBvZjHAPSTsjzmZXm4z32/ujAn/RhIkZw3ZgRKrxzryttGnWJJ8OFyF31JTJgwWWuPdH53G15PC83ZbmEgSV3win51RZRVppN4uQUuaqZWG9wwk2a6P5aen1RLCSLpTkd2mAEk9PlgmJrf8vITkiU9pF9n68ENCoo556qSdxW2pxnjrzKVPSqmqO1Xg5K4LOX4/9N4n4qkLEOiqnzzJClhFif3O28RW86RPxERGdPT81UI0oDAcU5euQr8Emz+Hd+PY1115UIld3CIHib5PYL9Ee0bFUKiWpR/acSe1fHB64mCoHP7hjFepGsq7inkvg2651wUDKBshGltpNkMj6+aZedNc0/rKYyjl80nT8g8QECgOSRzpmYp0zli2HpFoLOiWw==","name":"ci-customer-sshkey-rsa"}]}}],"name":"ci-group"}]`
	var allSuccess = `{"header":["Email","Name","Type","Value"],"data":[["c@c.com","deploy@nhmrc","ssh-rsa","AAAAB3NzaC1yc2EAAAADAQABAAACAQC++bRFdPP6d3kdXv1eImtfSgHhumcsy4IhAYId23v85nmcnTMqA5ahCoOzChPuxKWVsTGaU3xh+PMkQAO/HAkyUYIK8VlIMP/w9+VYraOYHwCZBZqiKwJH6XjpX24qTzdNYU8WdC+6OdOV+0SrEdtduxJV/TnVkoe+Ga+y7013mxCbw5y9LIPs/eBXjwuN/lASxaZGpzAP5FipKC/HOC+oS96gaNVQgmWl3Lm+faGpq2V1afEE9A0RXYhvQ6qG9qvcmFGtqLyhu6m7i9tjNiTlXalsFeox0pu8cnFvKZZZnFwi1EI4ngBUUg/hTFWmr13F2TslEJrnCqm8efs6o40l3tsdD64Jr1q9LUvqKWrwrv4B2vM2O3iccaE5Ll7fNH4pRDTZpFRL92MoX+TBpPjLxFYhz5zOJUKiFMsERkHJB/28a2BPU2etThwkIy5EwOrwHl/Q2KMxrddwwfd9FAZnmHXoUA0OtcZtgyrBDECOzuleGSyhcbXynQBwiEJ1RRQrbeWBTberQi4v+rDKgauxxVyfxH4yQfkoAFwt+QjQPUWv/kKCS+8LEJv1QlEd0lNh2A+TO/ugmsdshO+PzYKUtm4wMiqo0XfzyJGJFvYKbUbPNSZ7iGPRKkwQot8UQgAY/jwXn6z1sm8VmwDLirP1IUIHdt8pGarTffufPd9Rww=="],["ci-customer-user-ed25519@example.com","ci-customer-sshkey-ed25519","ssh-ed25519","AAAAC3NzaC1lZDI1NTE5AAAAIMdEs1h19jv2UrbtKcqPDatUxT9lPYcbGlEAbInsY8Ka"],["ci-customer-user-rsa@example.com","ci-customer-sshkey-rsa","ssh-rsa","AAAAB3NzaC1yc2EAAAADAQABAAACAQDEZlms5XsiyWjmnnUyhpt93VgHypse9Bl8kNkmZJTiM3Ex/wZAfwogzqd2LrTEiIOWSH1HnQazR+Cc9oHCmMyNxRrLkS/MEl0yZ38Q+GDfn37h/llCIZNVoHlSgYkqD0MQrhfGL5AulDUKIle93dA6qdCUlnZZjDPiR0vEXR36xGuX7QYAhK30aD2SrrBruTtFGvj87IP/0OEOvUZe8dcU9G/pCoqrTzgKqJRpqs/s5xtkqLkTIyR/SzzplO21A+pCKNax6csDDq3snS8zfx6iM8MwVfh8nvBW9seax1zBvZjHAPSTsjzmZXm4z32/ujAn/RhIkZw3ZgRKrxzryttGnWJJ8OFyF31JTJgwWWuPdH53G15PC83ZbmEgSV3win51RZRVppN4uQUuaqZWG9wwk2a6P5aen1RLCSLpTkd2mAEk9PlgmJrf8vITkiU9pF9n68ENCoo556qSdxW2pxnjrzKVPSqmqO1Xg5K4LOX4/9N4n4qkLEOiqnzzJClhFif3O28RW86RPxERGdPT81UI0oDAcU5euQr8Emz+Hd+PY1115UIld3CIHib5PYL9Ee0bFUKiWpR/acSe1fHB64mCoHP7hjFepGsq7inkvg2651wUDKBshGltpNkMj6+aZedNc0/rKYyjl80nT8g8QECgOSRzpmYp0zli2HpFoLOiWw=="]]}`

	processedList, err := processReturnedUserKeysList([]byte(userList))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	testResult, err := processAllUserKeysList(processedList)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "list sshkeys processing failed")
	}
}
func TestListSpecificUserKeys(t *testing.T) {
	var testUser = `c@c.com`
	var userList = `[{"id":"07b3d263-3a3e-4d0e-ac3f-56eab1e80df9","members":[{"role":"MAINTAINER","user":{"email":"c@c.com","firstName":"bob","id":"a08f166e-64f0-4461-9daa-e62b6f414faf","lastName":"bob","sshKeys":[{"keyType":"ssh-rsa","keyValue":"AAAAB3NzaC1yc2EAAAADAQABAAACAQC++bRFdPP6d3kdXv1eImtfSgHhumcsy4IhAYId23v85nmcnTMqA5ahCoOzChPuxKWVsTGaU3xh+PMkQAO/HAkyUYIK8VlIMP/w9+VYraOYHwCZBZqiKwJH6XjpX24qTzdNYU8WdC+6OdOV+0SrEdtduxJV/TnVkoe+Ga+y7013mxCbw5y9LIPs/eBXjwuN/lASxaZGpzAP5FipKC/HOC+oS96gaNVQgmWl3Lm+faGpq2V1afEE9A0RXYhvQ6qG9qvcmFGtqLyhu6m7i9tjNiTlXalsFeox0pu8cnFvKZZZnFwi1EI4ngBUUg/hTFWmr13F2TslEJrnCqm8efs6o40l3tsdD64Jr1q9LUvqKWrwrv4B2vM2O3iccaE5Ll7fNH4pRDTZpFRL92MoX+TBpPjLxFYhz5zOJUKiFMsERkHJB/28a2BPU2etThwkIy5EwOrwHl/Q2KMxrddwwfd9FAZnmHXoUA0OtcZtgyrBDECOzuleGSyhcbXynQBwiEJ1RRQrbeWBTberQi4v+rDKgauxxVyfxH4yQfkoAFwt+QjQPUWv/kKCS+8LEJv1QlEd0lNh2A+TO/ugmsdshO+PzYKUtm4wMiqo0XfzyJGJFvYKbUbPNSZ7iGPRKkwQot8UQgAY/jwXn6z1sm8VmwDLirP1IUIHdt8pGarTffufPd9Rww==","name":"deploy@nhmrc"}]}},{"role":"OWNER","user":{"email":"ci-customer-user-ed25519@example.com","firstName":null,"id":"23781ccd-8e35-4206-a7cf-97153311ba91","lastName":null,"sshKeys":[{"keyType":"ssh-ed25519","keyValue":"AAAAC3NzaC1lZDI1NTE5AAAAIMdEs1h19jv2UrbtKcqPDatUxT9lPYcbGlEAbInsY8Ka","name":"ci-customer-sshkey-ed25519"}]}},{"role":"OWNER","user":{"email":"ci-customer-user-rsa@example.com","firstName":null,"id":"906391b3-b3b2-43cb-82e7-fa0c07007fbc","lastName":null,"sshKeys":[{"keyType":"ssh-rsa","keyValue":"AAAAB3NzaC1yc2EAAAADAQABAAACAQDEZlms5XsiyWjmnnUyhpt93VgHypse9Bl8kNkmZJTiM3Ex/wZAfwogzqd2LrTEiIOWSH1HnQazR+Cc9oHCmMyNxRrLkS/MEl0yZ38Q+GDfn37h/llCIZNVoHlSgYkqD0MQrhfGL5AulDUKIle93dA6qdCUlnZZjDPiR0vEXR36xGuX7QYAhK30aD2SrrBruTtFGvj87IP/0OEOvUZe8dcU9G/pCoqrTzgKqJRpqs/s5xtkqLkTIyR/SzzplO21A+pCKNax6csDDq3snS8zfx6iM8MwVfh8nvBW9seax1zBvZjHAPSTsjzmZXm4z32/ujAn/RhIkZw3ZgRKrxzryttGnWJJ8OFyF31JTJgwWWuPdH53G15PC83ZbmEgSV3win51RZRVppN4uQUuaqZWG9wwk2a6P5aen1RLCSLpTkd2mAEk9PlgmJrf8vITkiU9pF9n68ENCoo556qSdxW2pxnjrzKVPSqmqO1Xg5K4LOX4/9N4n4qkLEOiqnzzJClhFif3O28RW86RPxERGdPT81UI0oDAcU5euQr8Emz+Hd+PY1115UIld3CIHib5PYL9Ee0bFUKiWpR/acSe1fHB64mCoHP7hjFepGsq7inkvg2651wUDKBshGltpNkMj6+aZedNc0/rKYyjl80nT8g8QECgOSRzpmYp0zli2HpFoLOiWw==","name":"ci-customer-sshkey-rsa"}]}}],"name":"ci-group"}]`
	var allSuccess = `{"header":["Email","Name","Type","Value"],"data":[["c@c.com","deploy@nhmrc","ssh-rsa","AAAAB3NzaC1yc2EAAAADAQABAAACAQC++bRFdPP6d3kdXv1eImtfSgHhumcsy4IhAYId23v85nmcnTMqA5ahCoOzChPuxKWVsTGaU3xh+PMkQAO/HAkyUYIK8VlIMP/w9+VYraOYHwCZBZqiKwJH6XjpX24qTzdNYU8WdC+6OdOV+0SrEdtduxJV/TnVkoe+Ga+y7013mxCbw5y9LIPs/eBXjwuN/lASxaZGpzAP5FipKC/HOC+oS96gaNVQgmWl3Lm+faGpq2V1afEE9A0RXYhvQ6qG9qvcmFGtqLyhu6m7i9tjNiTlXalsFeox0pu8cnFvKZZZnFwi1EI4ngBUUg/hTFWmr13F2TslEJrnCqm8efs6o40l3tsdD64Jr1q9LUvqKWrwrv4B2vM2O3iccaE5Ll7fNH4pRDTZpFRL92MoX+TBpPjLxFYhz5zOJUKiFMsERkHJB/28a2BPU2etThwkIy5EwOrwHl/Q2KMxrddwwfd9FAZnmHXoUA0OtcZtgyrBDECOzuleGSyhcbXynQBwiEJ1RRQrbeWBTberQi4v+rDKgauxxVyfxH4yQfkoAFwt+QjQPUWv/kKCS+8LEJv1QlEd0lNh2A+TO/ugmsdshO+PzYKUtm4wMiqo0XfzyJGJFvYKbUbPNSZ7iGPRKkwQot8UQgAY/jwXn6z1sm8VmwDLirP1IUIHdt8pGarTffufPd9Rww=="]]}`

	processedList, err := processReturnedUserKeysList([]byte(userList))
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	testResult, err := processUserKeysList(processedList, testUser)
	if err != nil {
		t.Error("Should not fail if processing succeeded", err)
	}
	if string(testResult) != allSuccess {
		checkEqual(t, string(testResult), allSuccess, "list specific user sshkeys processing failed")
	}
}
