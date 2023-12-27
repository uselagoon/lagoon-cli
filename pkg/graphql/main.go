package graphql

import (
	"github.com/golang-jwt/jwt"
	"github.com/uselagoon/lagoon-cli/internal/lagoon"
	"github.com/uselagoon/lagoon-cli/pkg/api"
)

// LagoonAPI .
func LagoonAPI(lc *lagoon.Config, debug bool) (api.Client, error) {
	lagoon := lc.Current
	lagoonAPI, err := api.NewWithToken(
		lc.Lagoons[lagoon].Token,
		lc.Lagoons[lagoon].GraphQL,
	)
	lagoonAPI.Debug(debug)
	if err != nil {
		return nil, err
	}
	return lagoonAPI, nil
}

func hasValidToken(lc *lagoon.Config, lagoon string) bool {
	return lc.Lagoons[lagoon].Token != ""
}

// VerifyTokenExpiry verfies if the current token is valid or not
func VerifyTokenExpiry(lc *lagoon.Config, lagoon string) bool {
	var p jwt.Parser
	token, _, err := p.ParseUnverified(
		lc.Lagoons[lagoon].Token, &jwt.StandardClaims{})
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
	routerPattern
	branches
	pullrequests
	problemsUi
	factsUi
	productionEnvironment
	deployTargetConfigs{
		id
		deployTarget{
		  id
		  name
		  token
		}
		
	  }
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

// ProjectByNameMinimalFragment .
var ProjectByNameMinimalFragment = `fragment Project on Project {
	id
	name
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

// ProjectEnvironmentEnvVars .
var ProjectEnvironmentEnvVars = `fragment Project on Project {
	id
	name
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

// ProjectEnvironmentEnvVarsRevealed .
var ProjectEnvironmentEnvVarsRevealed = `fragment Project on Project {
	id
	name
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

// EnvironmentEnvVars .
var EnvironmentEnvVars = `fragment Environment on Environment {
	id
	name
	envVariables {
		id
		name
		scope
	}
}`

// EnvironmentEnvVarsRevealed .
var EnvironmentEnvVarsRevealed = `fragment Environment on Environment {
	id
	name
	envVariables {
		id
		name
		scope
		value
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
	deployTitle
	deployBaseRef
	deployHeadRef
	autoIdle
}`
