package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/spf13/viper"

	"github.com/situmorangbastian/grip/graph"
)

func main() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read env: ", err.Error())
	}

	port := viper.GetString("PORT")
	if port == "" {
		log.Fatal("port not set")
	}

	srv := handler.New(graph.NewExecutableSchema(graph.NewResolver()))
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.Use(extension.Introspection{})

	app := fiber.New()
	if viper.GetString("ENV") == "development" {
		app.Get("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	}
	app.Post("/query", adaptor.HTTPHandler(srv))

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.Shutdown(); err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
	}
}
