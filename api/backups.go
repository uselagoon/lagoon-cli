package api

import (
	"encoding/json"
	"errors"

	"github.com/machinebox/graphql"
)

// AddBackup .
func (api *Interface) AddBackup(backup AddBackup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($id: Int, $environment: Int!, $source: String!, $backupId: String!, $created: String!) {
		addBackup(input: {
			id: $id
			environment: $environment
			source: $source
			backupId: $backupId
			created: $created
		}) {
		  	...Backup
		}
	}` + backupFragment)
	generateVars(req, backup)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addBackup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// DeleteBackup .
func (api *Interface) DeleteBackup(backup DeleteBackup) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($backupId: String!) {
		deleteBackup(input: {
		  	backupId: $backupId
		})
	}`)
	generateVars(req, backup)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteBackup"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// UpdateRestore .
func (api *Interface) UpdateRestore(update UpdateRestore) ([]byte, error) {
	req := graphql.NewRequest(`
	mutation ($backupId: String!, $patch: UpdateRestorePatchInput!) {
		updateRestore(input: {
			backupId: $backupId
			patch: $patch
		}) {
		  	...Restore
		}
	}` + restoreFragment)
	generateVars(req, update)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateRestore"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// GetAllEnvironmentBackups .
func (api *Interface) GetAllEnvironmentBackups() ([]byte, error) {
	req := graphql.NewRequest(`
	query {
		allEnvironments {
			id
			name
			openshiftProjectName
			project {
				name
			}
			backups {
				...Backup
			}
		}
	}` + backupFragment)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["allEnvironments"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}

// GetEnvironmentBackups .
func (api *Interface) GetEnvironmentBackups(backups EnvironmentBackups) ([]byte, error) {
	req := graphql.NewRequest(`
	query environmentByOpenshiftProjectName($openshiftProjectName: String!) {
		environmentByOpenshiftProjectName(openshiftProjectName: $openshiftProjectName) {
			id
			name
			openshiftProjectName
			project {
				name
			}
			backups {
				id
				backupId
				source
				created
			}
		}
	}`)
	generateVars(req, backups)
	if api.debug {
		debugRequest(req)
	}
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["environmentByOpenshiftProjectName"])
	if err != nil {
		return []byte(""), err
	}
	if api.debug {
		debugResponse(jsonBytes)
	}
	if string(jsonBytes) == "null" {
		return []byte(""), errors.New("graphql: returned null")
	}
	return jsonBytes, nil
}
