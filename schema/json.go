package schema

import "github.com/graphql-go/graphql"

var json = graphql.NewScalar(
	graphql.ScalarConfig{
		Name: "JSON",
		ParseValue: func(value interface{}) interface{} {
			if value == nil {
				return make(map[string]interface{})
			}
			return value
		},
	},
)
