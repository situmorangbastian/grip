package schema

import "github.com/graphql-go/graphql"

// Example holds example information with graphql object
var Example = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Example",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

// ExampleEdge holds example edge information with graphql object
var ExampleEdge = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ExampleEdge",
		Fields: graphql.Fields{
			"node": &graphql.Field{
				Type: Example,
			},
			"cursor": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// ExampleResult holds example result information with graphql object
var ExampleResult = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ExampleResult",
		Fields: graphql.Fields{
			"edges": &graphql.Field{
				Type: graphql.NewList(ExampleEdge),
			},
			"pageInfo": &graphql.Field{
				Type: pageInfo,
			},
			"totalCount": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
