package api

import (
	"encoding/json"
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["addBackup"])
	if err != nil {
		return []byte(""), err
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["deleteBackup"])
	if err != nil {
		return []byte(""), err
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["updateRestore"])
	if err != nil {
		return []byte(""), err
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["allEnvironments"])
	if err != nil {
		return []byte(""), err
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
	returnType, err := api.RunQuery(req, Data{})
	if err != nil {
		return []byte(""), err
	}
	reMappedResult := returnType.(map[string]interface{})
	jsonBytes, err := json.Marshal(reMappedResult["environmentByOpenshiftProjectName"])
	if err != nil {
		return []byte(""), err
	}
	return jsonBytes, nil
}
