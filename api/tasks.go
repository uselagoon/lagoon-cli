package api

import (
	"github.com/machinebox/graphql"
)

// UpdateTask .
func (api *Interface) UpdateTask(task UpdateTask) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($id: Int!, $patch: UpdateTaskPatchInput!) {
		updateTask(input: {
			id: $id
			patch: $patch
		}) {
		  	...Task
		}
	}` + taskFragment)
	generateVars(req, task)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}
