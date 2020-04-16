package helpers

import (
	"context"
	"strings"

	"github.com/amazeeio/lagoon-cli/internal/lagoon"
	"github.com/amazeeio/lagoon-cli/internal/lagoon/client"
	"github.com/amazeeio/lagoon-cli/internal/schema"
)

// GetLagoonAPIVersion get the lagoon API version
func GetLagoonAPIVersion(lc *client.Client) (string, error) {
	// always start at v1.0.0
	lagoonVersion := "v1.0.0"

	// if we have the api version available, just use it
	lagoonAPIVersion, err := lagoon.GetLagoonAPIVersion(context.TODO(), lc)
	if err != nil {
		if !strings.Contains(err.Error(), `Cannot query field "lagoonVersion" on type "Query"`) {
			return lagoonVersion, err
		}
	}
	if lagoonAPIVersion.LagoonVersion != "" {
		lagoonVersion = lagoonAPIVersion.LagoonVersion
		if IsValidSemver(lagoonVersion) {
			// if we get a valid version from the api, drop off here
			return lagoonVersion, nil
		}
	}
	// otherwise lets try to determine from schema changes
	lagoonSchema, err := lagoon.GetLagoonSchema(context.TODO(), lc)
	if err != nil {
		return lagoonVersion, err
	}
	return DetermineLagoonVersion(lagoonVersion, *lagoonSchema)
}

// DetermineLagoonVersion will determine the version of the API based on the schema
func DetermineLagoonVersion(lagoonVersion string, lagoonSchema schema.LagoonSchema) (string, error) {
	var err error
	for _, schemaType := range lagoonSchema.Types {
		if schemaType.Name == "Mutation" {
			for _, field := range schemaType.Fields {
				if field.Name == "switchActiveStandby" {
					lagoonVersion, err = GreaterThanOrEqualVersion(lagoonVersion, "v1.4.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
			}
		}
		if schemaType.Name == "Query" {
			for _, field := range schemaType.Fields {
				if field.Name == "allGroups" {
					lagoonVersion, err = GreaterThanOrEqualVersion(lagoonVersion, "v1.1.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
				if field.Name == "me" {
					lagoonVersion, err = GreaterThanOrEqualVersion(lagoonVersion, "v1.3.0")
					if err != nil {
						return lagoonVersion, err
					}
				}
			}
		}
		if schemaType.Name == "NotificationMicrosoftTeams" {
			lagoonVersion, err = GreaterThanOrEqualVersion(lagoonVersion, "v1.2.0")
			if err != nil {
				return lagoonVersion, err
			}
		}
	}
	return lagoonVersion, nil
}
