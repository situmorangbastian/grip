package cmd

import (
	"github.com/graphql-go/graphql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/situmorangbastian/grip/internal/resolver"
	"github.com/situmorangbastian/grip/schema"
)

var (
	// RootCMD is root command init
	RootCMD = &cobra.Command{
		Use:   "grip",
		Short: "grip is a boilerplate for graphql with go",
		Long:  "For more information, see https://github.com/situmorangbastian/grip",
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	RootCMD.AddCommand(graphqlServerCMD)
	RootCMD.AddCommand(graphiqlServerCMD)
}

func initConfig() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initSchema() *graphql.Schema {
	resolver := resolver.NewResolver().Build()

	gripschema := schema.NewSchema().WithResolver(resolver).Build()

	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    gripschema.Query(),
		Mutation: gripschema.Mutation(),
	})
	if err != nil {
		log.Fatal(err)
	}

	return &graphqlSchema
}
