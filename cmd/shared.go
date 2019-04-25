package cmd

import (
	context "context"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/spf13/viper"
)

type ProjectByName struct {
	ProjectByName Project `json:"projectByName"`
}
type WhatIsThere struct {
	AllProjects []Project `json:"allProjects"`
}
type Environments struct {
	Name            string `json:"name"`
	EnvironmentType string `json:"environmentType"`
	DeployType      string `json:"deployType"`
	Route           string `json:"route"`
}
type Project struct {
	ID                           int            `json:"id"`
	GitURL                       string         `json:"gitUrl"`
	Subfolder                    string         `json:"subfolder"`
	Name                         string         `json:"name"`
	Branches                     string         `json:"branches"`
	Pullrequests                 string         `json:"pullrequests"`
	ProductionEnvironment        string         `json:"productionEnvironment"`
	Environments                 []Environments `json:"environments"`
	AutoIdle                     int            `json:"autoIdle"`
	DevelopmentEnvironmentsLimit int            `json:"developmentEnvironmentsLimit"`
}

func getGraphQLToken() string {
	return viper.GetString("lagoon_token")
}
func ValidateToken() bool {
	return getGraphQLToken() != ""
}

func GraphQLClient() *graphql.Client {
	return graphql.NewClient(viper.GetString("lagoon_graphql"))
}
func GraphQLRequest(q string, resp interface{}) error {
	client := GraphQLClient()
	req := graphql.NewRequest(q)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", getGraphQLToken()))
	ctx := context.Background()
	return client.Run(ctx, req, &resp)
}
