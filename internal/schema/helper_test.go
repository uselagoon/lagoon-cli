package schema

import "encoding/json"

// GraphqlResponse is a wrapper for unmarshalling raw graphql responses.
// This is exported for tests only.
type GraphQLResponse struct {
	Response *struct {
		Project *Project `json:"projectByName"`
	} `json:"data"`
}

// UnmarshalProjectConfig takes raw JSON and returns the unmarshalled
// ProjectConfig. This is exported for tests only.
func UnmarshalProjectConfigData(data []byte, r *GraphQLResponse) error {
	return json.Unmarshal(data, &r)
}
