package schema

// place to put general schemas that may not be specific to any other type

// LagoonVersion is for the lagoon API version.
type LagoonVersion struct {
	LagoonVersion string `json:"lagoonVersion,omitempty"`
}

// LagoonSchema is the supported schema for a lagoon
type LagoonSchema struct {
	QueryType struct {
		Name string `json:"name"`
	} `json:"queryType"`
	MutationType struct {
		Name string `json:"name"`
	} `json:"mutationType"`
	SubscriptionType struct {
		Name string `json:"name"`
	} `json:"subscriptionType"`
	Types []struct {
		Kind        string `json:"kind"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Fields      []struct {
			Name        string        `json:"name"`
			Description string        `json:"description"`
			Args        []interface{} `json:"args"`
			Type        struct {
				Kind   string      `json:"kind"`
				Name   string      `json:"name"`
				OfType interface{} `json:"ofType"`
			} `json:"type"`
			IsDeprecated      bool        `json:"isDeprecated"`
			DeprecationReason interface{} `json:"deprecationReason"`
		} `json:"fields"`
		InputFields   interface{}   `json:"inputFields"`
		Interfaces    []interface{} `json:"interfaces"`
		EnumValues    interface{}   `json:"enumValues"`
		PossibleTypes interface{}   `json:"possibleTypes"`
	} `json:"types"`
	Directives []struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Locations   []string `json:"locations"`
		Args        []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Type        struct {
				Kind   string      `json:"kind"`
				Name   interface{} `json:"name"`
				OfType struct {
					Kind   string      `json:"kind"`
					Name   string      `json:"name"`
					OfType interface{} `json:"ofType"`
				} `json:"ofType"`
			} `json:"type"`
			DefaultValue interface{} `json:"defaultValue"`
		} `json:"args"`
	} `json:"directives"`
}

/* This query can get the entire schema of a lagoon, but it queries for a lot of information
and shouldn't be run very often, we use a small version of it for
`internal/lagoon/client/_lgraphql/lagoonSchema.graphql`

````
query IntrospectionQuery {
  __schema {
    queryType {
      name
    }
    mutationType {
      name
    }
    subscriptionType {
      name
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}

fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}

fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}

fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
````
*/
