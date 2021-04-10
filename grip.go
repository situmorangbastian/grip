package grip

import "github.com/graphql-go/graphql"

type Resolver interface {
	ExampleResolver(params graphql.ResolveParams) (interface{}, error)
}
