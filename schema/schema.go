package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/situmorangbastian/grip"
)

type Schema struct {
	resolver grip.Resolver
}

func (s Schema) Query() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"examples": &graphql.Field{
				Type:        ExampleResult,
				Description: "Fetch Example Entity",
				Args:        graphql.FieldConfigArgument{},
				Resolve:     s.resolver.ExampleResolver,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

func (s Schema) Mutation() *graphql.Object {
	objectConfig := graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"storeExample": &graphql.Field{
				Type:        graphql.NewNonNull(Example),
				Description: "Store Example",
				Args: graphql.FieldConfigArgument{
					"example": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(json),
					},
				},
				Resolve: s.resolver.ExampleResolver,
			},
		},
	}

	return graphql.NewObject(objectConfig)
}

// Initiator is a function type
type Initiator func(r *Schema) *Schema

// WithResolver is a function which can be called when the schema needs a resolver
func (i Initiator) WithResolver(resolver grip.Resolver) Initiator {
	return func(r *Schema) *Schema {
		i(r).resolver = resolver
		return r
	}
}

// Build is function which need to be called when constructing grapql schema
func (i Initiator) Build() *Schema {
	return i(&Schema{})
}

// NewSchema create a new grapql schema
func NewSchema() Initiator {
	return func(r *Schema) *Schema {
		return r
	}
}
