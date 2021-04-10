package grip

import "github.com/graphql-go/graphql"

type Resolver interface {
	Fetch(params graphql.ResolveParams) (interface{}, error)
	Store(params graphql.ResolveParams) (interface{}, error)
}
