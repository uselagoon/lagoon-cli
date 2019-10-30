package api

import (
	//"bytes"
	"context"
	//"errors"
	//"fmt"
	//"io/ioutil"
	"net/http"
	"regexp"
	//"strings"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/machinebox/graphql"
)

// Client struct
type Client interface {
	// Assorted
	RunQuery(*graphql.Request, interface{}) (interface{}, error)
	Request(CustomRequest) (interface{}, error)
	SanitizeGroupName(string) string
	SanitizeProjectName(string) string
	// Users
	AddUser(User) (interface{}, error)
	UpdateUser(UpdateUser) (interface{}, error)
	DeleteUser(User) (interface{}, error)
	GetUserBySSHKey(SSHKeyValue) (interface{}, error)
	AddSSHKey(AddSSHKey) (interface{}, error)
	DeleteSSHKey(DeleteSSHKey) (interface{}, error)
	// Tasks
	UpdateTask(UpdateTask) (interface{}, error)
	// Backups
	AddBackup(AddBackup) (interface{}, error)
	DeleteBackup(DeleteBackup) (interface{}, error)
	UpdateRestore(UpdateRestore) (interface{}, error)
	GetAllEnvironmentBackups() (interface{}, error)
	GetEnvironmentBackups(EnvironmentBackups) (interface{}, error)
	// Groups
	AddGroup(AddGroup) (interface{}, error)
	AddGroupWithParent(AddGroup) (interface{}, error)
	UpdateGroup(UpdateGroup) (interface{}, error)
	DeleteGroup(AddGroup) (interface{}, error)
	AddUserToGroup(AddUserToGroup) (interface{}, error)
	AddGroupToProject(ProjectToGroup) (interface{}, error)
	RemoveGroupFromProject(ProjectToGroup) (interface{}, error)
	RemoveUserFromGroup(UserGroup) (interface{}, error)
	// Environments
	GetEnvironmentByName(EnvironmentByName) (interface{}, error)
	AddOrUpdateEnvironment(AddUpdateEnvironment) (interface{}, error)
	UpdateEnvironment(UpdateEnvironment) (interface{}, error)
	DeleteEnvironment(DeleteEnvironment) (interface{}, error)
	SetEnvironmentServices(SetEnvironmentServices) (interface{}, error)
	// Projects
	GetOpenShiftInfoForProject(Project) (interface{}, error)
	AddProject(ProjectPatch, string) (interface{}, error)
	UpdateProject(UpdateProject, string) (interface{}, error)
	DeleteProject(Project) (interface{}, error)
	GetProductionEnvironmentForProject(Project) (interface{}, error)
	GetEnvironmentByOpenshiftProjectName(Environment) (interface{}, error)
	GetProjectsByGitURL(Project) (interface{}, error)
	GetProjectByName(Project, string) (interface{}, error)
	GetAllProjects(string) (interface{}, error)
	GetRocketChatInfoForProject(Project) (interface{}, error)
	GetSlackinfoForProject(Project) (interface{}, error)
	GetActiveSystemForProject(Project, string) (interface{}, error)
	GetEnvironmentsForProject(Project) (interface{}, error)
	GetDeploymentByRemoteID(Deployment) (interface{}, error)
	AddDeployment(Deployment) (interface{}, error)
	UpdateDeployment(UpdateDeployment) (interface{}, error)
}

// Interface struct
type Interface struct {
	tokenSigningKey string
	token           string
	jwtAudience     string
	graphQLEndpoint string
	netClient       *http.Client
	graphqlClient   *graphql.Client
}

// CustomRequest .
type CustomRequest struct {
	Query     string
	Variables map[string]interface{}
}

var netClient = &http.Client{
	Timeout: time.Second * time.Duration(10),
}

// New creates an interface against an endpoint with the api using jwt signing key and audience
func New(tokenSigningKey string, jwtAudience string, graphQLEndpoint string) (Client, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := graphql.NewClient(graphQLEndpoint)

	tokenString, tokenErr := getJWTToken(jwtAudience, tokenSigningKey)
	return &Interface{
		graphQLEndpoint: graphQLEndpoint,
		token:           tokenString,
		netClient:       netClient,
		graphqlClient:   client,
	}, tokenErr
}

// NewWithToken  creates an interface against an endpoint with the api using a jwt token
func NewWithToken(tokenString string, graphQLEndpoint string) (Client, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := graphql.NewClient(graphQLEndpoint)
	return &Interface{
		token:           tokenString,
		graphQLEndpoint: graphQLEndpoint,
		netClient:       netClient,
		graphqlClient:   client,
	}, nil
}

func getJWTToken(tokenAudience string, tokenSigningKey string) (string, error) {
	// generate a token with our claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":  int32(time.Now().Unix()),
		"role": "admin",
		"iss":  "lagoon-commons",
		"aud":  "" + tokenAudience + "",
		"sub":  "lagoon-commons",
	})
	// sign and get the complete encoded token as a string using the secret
	tokenString, tokenErr := token.SignedString([]byte(tokenSigningKey))
	return tokenString, tokenErr
}

// RunQuery run a graphql query against an api endpoint, return the result
func (api *Interface) RunQuery(graphQLQuery *graphql.Request, returnType interface{}) (interface{}, error) {
	graphQLQuery.Header.Set("Cache-Control", "no-cache")
	graphQLQuery.Header.Add("Authorization", "Bearer "+api.token)
	// define a Context for the request
	ctx := context.Background()
	// run it and capture the response
	err := api.graphqlClient.Run(ctx, graphQLQuery, &returnType)
	return returnType, err
}

// SanitizeGroupName .
func (api *Interface) SanitizeGroupName(name string) string {
	return sanitizeName(name)
}

// SanitizeProjectName .
func (api *Interface) SanitizeProjectName(name string) string {
	return sanitizeName(name)
}

func sanitizeName(name string) string {
	var re = regexp.MustCompile(`[^a-zA-Z0-9-]`)
	sanitizedName := re.ReplaceAllString(name, `$1-$2`)
	return sanitizedName
}

// Request .
func (api *Interface) Request(request CustomRequest) (interface{}, error) {
	req := graphql.NewRequest(request.Query)
	for varName, varValue := range request.Variables {
		req.Var(string(varName), varValue)
	}
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

func generateVars(request *graphql.Request, jsonData interface{}) {
	jsonString, _ := json.Marshal(jsonData)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &dat); err != nil {
		panic(err)
	}
	for varName, varValue := range dat {
		request.Var(string(varName), varValue)
	}
}
