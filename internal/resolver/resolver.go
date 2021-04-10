package resolver

import (
	"github.com/graphql-go/graphql"

	"github.com/situmorangbastian/grip"
)

type resolver struct {
}

func (r resolver) ExampleResolver(params graphql.ResolveParams) (interface{}, error) {
	return grip.ExampleEntity{
		ID: "example-id",
	}, nil
}

// Initiator is a function type
type Initiator func(r *resolver) *resolver

// Build is function which need to be called when constructing resolver
func (i Initiator) Build() grip.Resolver {
	return i(&resolver{})
}

// NewResolver as new resolver
func NewResolver() Initiator {
	return func(r *resolver) *resolver {
		return r
	}
}
