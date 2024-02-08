// Package client implements the interfaces required by the parent lagoon
// package.
package client

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/hashicorp/go-version"
	"github.com/machinebox/graphql"
)

//go:embed _lgraphql/*
var lgraphql embed.FS

// Client implements the lagoon package interfaces for the Lagoon GraphQL API.
type Client struct {
	userAgent  string
	token      string
	apiVersion string
	cliVersion string
	client     *graphql.Client
}

// New creates a new Client for the given endpoint.
func New(endpoint, token, apiVersion string, cliVersion string, debug bool) *Client {
	if debug {
		return &Client{
			apiVersion: apiVersion,
			cliVersion: cliVersion,
			token:      token,
			client: graphql.NewClient(endpoint,
				// enable debug logging to stderr
				func(c *graphql.Client) {
					l := log.New(os.Stderr, "graphql", 0)
					c.Log = func(s string) {
						l.Println(s)
					}
				}),
		}
	}
	return &Client{
		apiVersion: apiVersion,
		cliVersion: cliVersion,
		token:      token,
		client:     graphql.NewClient(endpoint),
	}
}

// newRequest constructs a graphql request.
// assetName is the name of the graphql query template in _graphql/.
// varStruct is converted to a map of variables for the template.
func (c *Client) newRequest(
	assetName string, varStruct interface{}) (*graphql.Request, error) {

	q, err := lgraphql.ReadFile(assetName)
	if err != nil {
		return nil, fmt.Errorf("couldn't get asset: %w", err)
	}

	return c.doRequest(string(q), varStruct)
}

// newRequest constructs a graphql request which varies based on the version provided.
// assetName is the name of the graphql query template in _graphql/.
// varStruct is converted to a map of variables for the template.
func (c *Client) newVersionedRequest(
	assetName string, varStruct interface{}) (*graphql.Request, error) {

	q, err := lgraphql.ReadFile(assetName)
	if err != nil {
		return nil, fmt.Errorf("couldn't get asset: %w", err)
	}

	t, err := template.New("query").
		Funcs(template.FuncMap{
			// apiVerGreaterThanOrEqual returns true if a is greater than or equal
			// to b, and false otherwise. a and b should both be valid Semantic
			// Versions.
			"apiVerGreaterThanOrEqual": func(a, b string) (bool, error) {
				aVer, err := version.NewSemver(a)
				if err != nil {
					return false, err
				}
				bVer, err := version.NewSemver(b)
				if err != nil {
					return false, err
				}
				return aVer.GreaterThanOrEqual(bVer), nil
			},
		}).
		Parse(string(q))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse template: %w", err)
	}

	queryBuilder := strings.Builder{}
	if err = t.Execute(&queryBuilder, c.apiVersion); err != nil {
		return nil, fmt.Errorf("couldn't execute template: %w", err)
	}

	return c.doRequest(queryBuilder.String(), varStruct)
}

func (c *Client) doRequest(query string, varStruct interface{}) (*graphql.Request, error) {
	vars, err := structToVarMap(varStruct)
	if err != nil {
		return nil, fmt.Errorf("couldn't convert struct to map: %w", err)
	}

	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}

	headers := map[string]string{
		"User-Agent":    fmt.Sprintf("lagoon-cli version: %s", c.cliVersion),
		"Authorization": fmt.Sprintf("Bearer %s", c.token),
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}

// structToVarMap encodes the given struct to a map. The idea is that by
// round-tripping through Marshal/Unmarshal, omitempty is applied to the
// zero-valued fields.
func structToVarMap(
	varStruct interface{}) (vars map[string]interface{}, err error) {
	data, err := json.Marshal(varStruct)
	if err != nil {
		return vars, err
	}
	return vars, json.Unmarshal(data, &vars)
}
