package schema

import "encoding/json"

// ProjectByNameResponse is a wrapper for unmarshalling raw graphql responses.
// This is exported for tests only.
type ProjectByNameResponse struct {
	Response *struct {
		Project *Project `json:"projectByName"`
	} `json:"data"`
}

// UnmarshalProjectConfig takes raw JSON and returns the unmarshalled
// ProjectConfig. This is exported for tests only.
func UnmarshalProjectByNameResponse(
	data []byte, r *ProjectByNameResponse) error {
	return json.Unmarshal(data, &r)
}
