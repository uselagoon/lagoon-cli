package api

import (
	"github.com/machinebox/graphql"
)

// AddBackup .
func (api *Interface) AddBackup(backup AddBackup) (interface{}, error) {
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
	return returnType, err
}

// DeleteBackup .
func (api *Interface) DeleteBackup(backup DeleteBackup) (interface{}, error) {
	req := graphql.NewRequest(`
	mutation ($backupId: String!) {
		deleteBackup(input: {
		  	backupId: $backupId
		})
	}`)
	generateVars(req, backup)
	returnType, err := api.RunQuery(req, Data{})
	return returnType, err
}

// UpdateRestore .
func (api *Interface) UpdateRestore(update UpdateRestore) (interface{}, error) {
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
	return returnType, err
}

// GetAllEnvironmentBackups .
func (api *Interface) GetAllEnvironmentBackups() (interface{}, error) {
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
	return returnType, err
}

// GetEnvironmentBackups .
func (api *Interface) GetEnvironmentBackups(backups EnvironmentBackups) (interface{}, error) {
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
	return returnType, err
}
