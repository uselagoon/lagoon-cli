package environments

import (
	"github.com/amazeeio/lagoon-cli/api"
	"github.com/amazeeio/lagoon-cli/graphql"
)

// DeployEnvironmentBranch .
func DeployEnvironmentBranch(projectName string, branchName string) ([]byte, error) {
	lagoonAPI, err := graphql.LagoonAPI()
	if err != nil {
		return []byte(""), err
	}

	customRequest := api.CustomRequest{
		Query: `mutation ($project: String!, $branch: String!){
			deployEnvironmentBranch(
				input: {
					project:{name: $project}
					branchName: $branch
				}
			)
		}`,
		Variables: map[string]interface{}{
			"project": projectName,
			"branch":  branchName,
		},
		MappedResult: "deployEnvironmentBranch",
	}
	returnResult, err := lagoonAPI.Request(customRequest)
	return returnResult, err
}
