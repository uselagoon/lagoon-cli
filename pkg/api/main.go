package api

import (
	"context"
	"fmt"
	"strings"

	"net/http"
	"regexp"

	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/logrusorgru/aurora"
	"github.com/machinebox/graphql"
)

// Client struct
type Client interface {
	// Assorted
	RunQuery(*graphql.Request, interface{}) (interface{}, error)
	Request(CustomRequest) ([]byte, error)
	SanitizeGroupName(string) string
	SanitizeProjectName(string) string
	Debug(bool)
	// Users
	AddUser(User) ([]byte, error)
	UpdateUser(UpdateUser) ([]byte, error)
	DeleteUser(User) ([]byte, error)
	GetUserBySSHKey(SSHKeyValue) ([]byte, error)
	AddSSHKey(AddSSHKey) ([]byte, error)
	DeleteSSHKey(DeleteSSHKey) ([]byte, error)
	// Tasks
	UpdateTask(UpdateTask) ([]byte, error)
	// Backups
	AddBackup(AddBackup) ([]byte, error)
	DeleteBackup(DeleteBackup) ([]byte, error)
	UpdateRestore(UpdateRestore) ([]byte, error)
	GetAllEnvironmentBackups() ([]byte, error)
	GetEnvironmentBackups(EnvironmentBackups) ([]byte, error)
	// Groups
	AddGroup(AddGroup) ([]byte, error)
	AddGroupWithParent(AddGroup) ([]byte, error)
	UpdateGroup(UpdateGroup) ([]byte, error)
	DeleteGroup(AddGroup) ([]byte, error)
	AddUserToGroup(AddUserToGroup) ([]byte, error)
	AddGroupToProject(ProjectToGroup) ([]byte, error)
	RemoveGroupFromProject(ProjectToGroup) ([]byte, error)
	RemoveUserFromGroup(UserGroup) ([]byte, error)
	// Environments
	GetEnvironmentByName(EnvironmentByName, string) ([]byte, error)
	AddOrUpdateEnvironment(AddUpdateEnvironment) ([]byte, error)
	UpdateEnvironment(UpdateEnvironment) ([]byte, error)
	DeleteEnvironment(DeleteEnvironment) ([]byte, error)
	SetEnvironmentServices(SetEnvironmentServices) ([]byte, error)
	// Projects
	GetOpenShiftInfoForProject(Project) ([]byte, error)
	AddProject(ProjectPatch, string) ([]byte, error)
	UpdateProject(UpdateProject, string) ([]byte, error)
	DeleteProject(Project) ([]byte, error)
	GetProductionEnvironmentForProject(Project) ([]byte, error)
	GetEnvironmentByOpenshiftProjectName(Environment) ([]byte, error)
	GetProjectsByGitURL(Project) ([]byte, error)
	GetProjectByName(Project, string) ([]byte, error)
	GetAllProjects(string) ([]byte, error)
	GetRocketChatInfoForProject(Project, string) ([]byte, error)
	GetSlackInfoForProject(Project, string) ([]byte, error)
	GetEnvironmentsForProject(Project) ([]byte, error)
	GetDeploymentByRemoteID(Deployment) ([]byte, error)
	AddDeployment(Deployment) ([]byte, error)
	UpdateDeployment(UpdateDeployment) ([]byte, error)
}

// Interface struct
type Interface struct {
	tokenSigningKey string
	token           string
	jwtAudience     string
	graphQLEndpoint string
	debug           bool
	netClient       *http.Client
	graphqlClient   *graphql.Client
}

// CustomRequest .
type CustomRequest struct {
	Query        string
	Variables    map[string]interface{}
	MappedResult string
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

// Debug .
func (api *Interface) Debug(debug bool) {
	api.debug = debug
}

func sanitizeName(name string) string {
	var re = regexp.MustCompile(`[^a-zA-Z0-9-]`)
	sanitizedName := re.ReplaceAllString(name, `$1-$2`)
	return sanitizedName
}

// Request .
func (api *Interface) Request(request CustomRequest) ([]byte, error) {
	req := graphql.NewRequest(request.Query)
	for varName, varValue := range request.Variables {
		req.Var(string(varName), varValue)
	}
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult[request.MappedResult])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	return jsonBytes, nil
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

// DebugData .
type DebugData struct {
	Query string     `json:"query"`
	Vars  []DebugVar `json:"variables"`
}

// DebugVar .
type DebugVar struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// debug the query and variables that are sent for the request
func debugRequest(req *graphql.Request) {
	data := DebugData{
		Query: req.Query(),
	}
	for n, v := range req.Vars() {
		data.Vars = append(data.Vars, DebugVar{Name: n, Value: v})
	}
	jsonData, _ := json.Marshal(data)
	fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Request"), strings.Replace(string(jsonData), "\\t", "", -1)))
}

func debugResponse(resp []byte) {
	fmt.Println(fmt.Sprintf("%s: %s", aurora.Yellow("Response"), string(resp)))
}
