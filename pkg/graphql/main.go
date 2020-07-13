package graphql

import (
	"github.com/amazeeio/lagoon-cli/pkg/api"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

// LagoonAPI .
func LagoonAPI(debug bool) (api.Client, error) {
	lagoon := viper.GetString("current")
	lagoonAPI, err := api.NewWithToken(
		viper.GetString("lagoons."+lagoon+".token"),
		viper.GetString("lagoons."+lagoon+".graphql"),
	)
	lagoonAPI.Debug(debug)
	if err != nil {
		return nil, err
	}
	return lagoonAPI, nil
}

func getGraphQLToken(lagoon string) string {
	return viper.GetString("lagoons." + lagoon + ".token")
}

func hasValidToken(lagoon string) bool {
	return getGraphQLToken(lagoon) != ""
}

// VerifyTokenExpiry verfies if the current token is valid or not
func VerifyTokenExpiry(lagoon string) bool {
	var p jwt.Parser
	token, _, err := p.ParseUnverified(
		viper.GetString("lagoons."+lagoon+".token"), &jwt.StandardClaims{})
	if err != nil {
		return false
	}
	if token.Claims.Valid() != nil {
		return false
	}
	return true
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

// ProjectEnvVars .
var ProjectEnvVars = `fragment Project on Project {
	id
	name
	envVariables {
		id
		name
		scope
	}
}`

// ProjectEnvVarsRevealed .
var ProjectEnvVarsRevealed = `fragment Project on Project {
	id
	name
	envVariables {
		id
		name
		scope
		value
	}
}`

// EnvironmentEnvVars .
var EnvironmentEnvVars = `fragment Environment on Environment {
	openshiftProjectName
	name
	envVariables {
		id
		name
		scope
	}
}`

// EnvironmentEnvVarsRevealed .
var EnvironmentEnvVarsRevealed = `fragment Environment on Environment {
	openshiftProjectName
	name
	envVariables {
		id
		name
		scope
		value
	}
}`

// AllProjectsFragment .
var AllProjectsFragment = `fragment Project on Project {
	id
	gitUrl
	name,
	developmentEnvironmentsLimit,
	productionEnvironment,
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

// ProjectNameID .
var ProjectNameID = `fragment Project on Project {
	id
	name
}`

// EnvironmentByNameFragment .
var EnvironmentByNameFragment = `fragment Environment on Environment {
	id
	name
	route
	routes
	deployType
	environmentType
	openshiftProjectName
	updated
	created
	deleted
	monitoringUrls
	deployTitle
	deployBaseRef
	deployHeadRef
	autoIdle
}`
