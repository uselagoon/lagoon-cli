//go:generate mockgen -source=import.go -destination=../mock/mock_importer.go -package=mock

package lagoon_test

import (
	"context"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/internal/mock"
	"github.com/uselagoon/lagoon-cli/internal/schema"
	"github.com/uselagoon/lagoon-cli/pkg/api"
)

// importCalls stores arrays of expected import calls associated with a given
// config file
type importCalls struct {
	NewProjectID                        uint
	NewEnvironmentID                    uint
	AddGroupInputs                      []schema.AddGroupInput
	AddUserInputs                       []schema.AddUserInput
	AddSSHKeyInputs                     []schema.AddSSHKeyInput
	UserGroupRoleInputs                 []schema.UserGroupRoleInput
	AddNotificationSlackInputs          []schema.AddNotificationSlackInput
	AddNotificationRocketChatInputs     []schema.AddNotificationRocketChatInput
	AddNotificationEmailInputs          []schema.AddNotificationEmailInput
	AddNotificationMicrosoftTeamsInputs []schema.AddNotificationMicrosoftTeamsInput
	AddProjectInputs                    []schema.AddProjectInput
	EnvVariableInputs                   []schema.EnvVariableInput
	AddEnvironmentInputs                []schema.AddEnvironmentInput
	ProjectGroupsInputs                 []schema.ProjectGroupsInput
	AddNotificationToProjectInputs      []schema.AddNotificationToProjectInput
}

func TestImport(t *testing.T) {
	var testCases = map[string]struct {
		input  string
		expect *importCalls
	}{
		"exhaustive": {input: "testdata/exhaustive.import.yaml", expect: &importCalls{
			NewProjectID:     99,
			NewEnvironmentID: 88,
			AddGroupInputs: []schema.AddGroupInput{
				{Name: "abc"},
			},
			AddUserInputs: []schema.AddUserInput{
				{Email: "foo@example.com", FirstName: "foofirst", LastName: "foolast"},
				{Email: "bar@example.com", FirstName: "barfirst", LastName: "barlast"},
				{Email: "projectuser@example.com", FirstName: "projectuserfirst",
					LastName: "projectuserlast"},
			},
			AddSSHKeyInputs: []schema.AddSSHKeyInput{
				{
					UserEmail: "foo@example.com",
					SSHKey: schema.SSHKey{
						Name:     "foo-example",
						KeyValue: "AAAAC3NzaC1lZDI1NTE5AAAAIPKqJ+OLYLCLJlUTF8SWVOwdUrCFfPVcNMF4Rr+rfXY3",
						KeyType:  api.SSHEd25519,
					},
				}, {
					UserEmail: "bar@example.com",
					SSHKey: schema.SSHKey{
						Name:     "bar-example",
						KeyValue: "AAAAC3NzaC1lZDI1NTE5AAAAIPKqJ+OLYLCLJlUTF8SWVOwdUrCFfPVcNMF4Rr+rfXY3",
						KeyType:  api.SSHEd25519,
					},
				}, {
					UserEmail: "projectuser@example.com",
					SSHKey: schema.SSHKey{
						Name:     "projectuser",
						KeyValue: "AAAAC3NzaC1lZDI1NTE5AAAAIPKqJ+OLYLCLJlUTF8SWVOwdUrCFfPVcNMF4Rr+rfXY3",
						KeyType:  api.SSHEd25519,
					},
				},
			},
			UserGroupRoleInputs: []schema.UserGroupRoleInput{
				{
					UserEmail: "foo@example.com",
					GroupName: "abc",
					GroupRole: api.OwnerRole,
				}, {
					UserEmail: "bar@example.com",
					GroupName: "abc",
					GroupRole: api.OwnerRole,
				}, {
					UserEmail: "projectuser@example.com",
					GroupName: "project-bananas",
					GroupRole: api.MaintainerRole,
				},
			},
			AddNotificationSlackInputs: []schema.AddNotificationSlackInput{
				{
					Name:    "example-slack",
					Webhook: "https://hooks.slack.example.com/services/xxx/yyy",
					Channel: "build-notifications",
				},
			},
			AddNotificationRocketChatInputs: []schema.AddNotificationRocketChatInput{
				{
					Name:    "example-rocketchat",
					Webhook: "https://hooks.rocketchat.example.com/services/xxx/yyy",
					Channel: "build-notifications",
				},
			},
			AddNotificationEmailInputs: []schema.AddNotificationEmailInput{
				{
					Name:         "example-email",
					EmailAddress: "example@example.com",
				},
			},
			AddNotificationMicrosoftTeamsInputs: []schema.AddNotificationMicrosoftTeamsInput{
				{
					Name:    "example-msteams",
					Webhook: "https://hooks.msteams.example.com/services/xxx/yyy",
				},
			},
			AddProjectInputs: []schema.AddProjectInput{
				{
					Name:                         "bananas",
					GitURL:                       "git@github.amazee.io:foo-bar/bananas-au.git",
					Openshift:                    2,
					Branches:                     "^(master|develop|production)$",
					ProductionEnvironment:        "production",
					AutoIdle:                     1,
					StorageCalc:                  1,
					DevelopmentEnvironmentsLimit: 10,
				},
			},
			EnvVariableInputs: []schema.EnvVariableInput{
				{
					EnvKeyValue: schema.EnvKeyValue{
						Name:  "ENABLE_REDIS",
						Scope: schema.GlobalVar,
						Value: "1",
					},
					Type: schema.ProjectVar,
					// NewProjectID
					TypeID: 99,
				}, {
					EnvKeyValue: schema.EnvKeyValue{
						Name:  "ENABLE_REDIS",
						Scope: schema.BuildVar,
						Value: "1",
					},
					Type: schema.EnvironmentVar,
					// NewEnvironmentID
					TypeID: 88,
				},
			},
			AddEnvironmentInputs: []schema.AddEnvironmentInput{
				{
					Name:                 "develop",
					OpenshiftProjectName: "bananas-develop",
					ProjectID:            99, // NewProjectID
				}, {
					Name:                 "master",
					OpenshiftProjectName: "bananas-master",
					ProjectID:            99, // NewProjectID
				}, {
					Name:                 "sandbox",
					OpenshiftProjectName: "bananas-sandbox",
					ProjectID:            99, // NewProjectID
				},
			},
			ProjectGroupsInputs: []schema.ProjectGroupsInput{
				{
					Project: schema.ProjectInput{Name: "bananas"},
					Groups:  []schema.GroupInput{{Name: "abc"}},
				},
			},
			AddNotificationToProjectInputs: []schema.AddNotificationToProjectInput{
				{
					Project:          "bananas",
					NotificationType: api.SlackNotification,
					NotificationName: "example-slack",
				},
			},
		}},
	}
	for name, tc := range testCases {
		t.Run(name, func(tt *testing.T) {
			ctx := context.Background()
			// set up the mock importer
			ctrl := gomock.NewController(tt)
			defer ctrl.Finish()
			importer := mock.NewMockImporter(ctrl)
			// use the provided importCalls to set the expectations
			for i := range tc.expect.AddGroupInputs {
				importer.EXPECT().AddGroup(ctx, &tc.expect.AddGroupInputs[i], nil)
			}
			for i := range tc.expect.AddUserInputs {
				importer.EXPECT().AddUser(ctx, &tc.expect.AddUserInputs[i], nil)
			}
			for i := range tc.expect.AddSSHKeyInputs {
				importer.EXPECT().AddSSHKey(ctx, &tc.expect.AddSSHKeyInputs[i], nil)
			}
			for i := range tc.expect.UserGroupRoleInputs {
				importer.EXPECT().AddUserToGroup(
					ctx, &tc.expect.UserGroupRoleInputs[i], nil)
			}
			for i := range tc.expect.AddNotificationSlackInputs {
				importer.EXPECT().AddNotificationSlack(
					ctx, &tc.expect.AddNotificationSlackInputs[i], nil)
			}
			for i := range tc.expect.AddNotificationRocketChatInputs {
				importer.EXPECT().AddNotificationRocketChat(
					ctx, &tc.expect.AddNotificationRocketChatInputs[i], nil)
			}
			for i := range tc.expect.AddNotificationEmailInputs {
				importer.EXPECT().AddNotificationEmail(
					ctx, &tc.expect.AddNotificationEmailInputs[i], nil)
			}
			for i := range tc.expect.AddNotificationMicrosoftTeamsInputs {
				importer.EXPECT().AddNotificationMicrosoftTeams(
					ctx, &tc.expect.AddNotificationMicrosoftTeamsInputs[i], nil)
			}
			for i := range tc.expect.AddProjectInputs {
				importer.EXPECT().AddProject(
					ctx, &tc.expect.AddProjectInputs[i], &schema.Project{}).Do(
					func(_ context.Context,
						_ *schema.AddProjectInput, p *schema.Project) {
						// set the ProjectID as the env variables calls require it
						p.ID = tc.expect.NewProjectID
					})
			}
			for i := range tc.expect.EnvVariableInputs {
				importer.EXPECT().AddEnvVariable(
					ctx, &tc.expect.EnvVariableInputs[i], nil)
			}
			for i := range tc.expect.AddEnvironmentInputs {
				importer.EXPECT().AddOrUpdateEnvironment(
					ctx, &tc.expect.AddEnvironmentInputs[i], &schema.Environment{}).Do(
					func(_ context.Context,
						_ *schema.AddEnvironmentInput, e *schema.Environment) {
						// set the EnvironmentID as the env variables calls require it
						e.ID = tc.expect.NewEnvironmentID
					})
			}
			for i := range tc.expect.ProjectGroupsInputs {
				importer.EXPECT().AddGroupsToProject(
					ctx, &tc.expect.ProjectGroupsInputs[i], nil)
			}
			for i := range tc.expect.AddNotificationToProjectInputs {
				importer.EXPECT().AddNotificationToProject(
					ctx, &tc.expect.AddNotificationToProjectInputs[i], nil)
			}
			// open the test yaml
			file, err := os.Open(tc.input)
			if err != nil {
				tt.Fatal(err)
			}
			// run the import
			if err := lagoon.Import(ctx, importer, file, true, 2); err != nil {
				tt.Fatal(err)
			}
		})
	}
}
