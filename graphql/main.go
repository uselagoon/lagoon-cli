package graphql

import (
	"context"
	"fmt"
	"github.com/amazeeio/lagoon-cli/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/machinebox/graphql"
	"github.com/spf13/viper"
)

// HasValidToken .
func HasValidToken() bool {
	return getGraphQLToken() != ""
}

// LagoonAPI .
func LagoonAPI() (api.Client, error) {
	lagoon := viper.GetString("current")
	lagoonAPI, err := api.NewWithToken(
		viper.GetString("lagoons."+lagoon+".token"),
		viper.GetString("lagoons."+lagoon+".graphql"),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return lagoonAPI, nil
}

func getGraphQLToken() string {
	lagoon := viper.GetString("current")
	return viper.GetString("lagoons." + lagoon + ".token")
}

// GraphQLClient returns a new GraphQL client.
func GraphQLClient() *graphql.Client {
	lagoon := viper.GetString("current")
	return graphql.NewClient(viper.GetString("lagoons." + lagoon + ".graphql"))
}

// GraphQLRequest performs a GraphQL request.
func GraphQLRequest(q string, resp interface{}) error {
	client := GraphQLClient()
	req := graphql.NewRequest(q)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", getGraphQLToken()))
	ctx := context.Background()
	return client.Run(ctx, req, &resp)
}

// VerifyTokenExpiry verfies if the current token is valid or not
func VerifyTokenExpiry() bool {
	if HasValidToken() {
		var p jwt.Parser
		token, _, err := p.ParseUnverified(getGraphQLToken(), &jwt.StandardClaims{})
		if err = token.Claims.Valid(); err != nil {
			//handle invalid token
			return false
		}
		return true
	}
	return false
}

// DefaultFragment is blank to use what is defined in api
var DefaultFragment = ""

// ProjectByNameFragment .
var ProjectByNameFragment = `fragment Project on Project {
	id
	name
	gitUrl
	subfolder
	branches
	pullrequests
	productionEnvironment
	environments {
		id
		name
		openshiftProjectName
		environmentType
		deployType
		route
	}
	autoIdle
	storageCalc
	developmentEnvironmentsLimit
}`

// ProjectEnvVars .
var ProjectEnvVars = `fragment Project on Project {
	id
	name
	envVariables {
		id
		name
	}
}`

// ProjectAndEnvironmentEnvVars .
var ProjectAndEnvironmentEnvVars = `fragment Project on Project {
	id
	name
	envVariables {
		id
		name
		scope
	}
	environments {
		openshiftProjectName
		name
		envVariables {
			id
			name
			scope
		}
	}
}`

// ProjectAndEnvironmentEnvVarsRevealed .
var ProjectAndEnvironmentEnvVarsRevealed = `fragment Project on Project {
	id
	name
	envVariables {
		id
		name
		scope
		value
	}
	environments {
		openshiftProjectName
		name
		envVariables {
			id
			name
			scope
			value
		}
	}
}`

// AllProjectsFragment .
var AllProjectsFragment = `fragment Project on Project {
	id
	gitUrl
	name,
	developmentEnvironmentsLimit,
	environments {
		environmentType,
		route
	}
}`

// RocketChatFragment .
var RocketChatFragment = `fragment Notification on NotificationRocketChat {
	id
	name
	webhook
	channel
}`

// SlackFragment .
var SlackFragment = `fragment Notification on NotificationSlack {
	id
	name
	webhook
	channel
}`

// EnvironmentVariablesFragment .
var EnvironmentVariablesFragment = `fragment Environment on Environment {
	id
	name
	environmentType
	openshiftProjectName
	envVariables {
		id
		name
	}
}`
