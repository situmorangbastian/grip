package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/graphql-go/graphql/gqlerrors"
	graphqlHandler "github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var graphqlServerCMD = &cobra.Command{
	Use:   "graphql",
	Short: "start graphql server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.Server.ReadTimeout = time.Duration(viper.GetInt("http.server_read_timeout")) * time.Second
		e.Server.WriteTimeout = time.Duration(viper.GetInt("http.server_write_timeout")) * time.Second

		e.GET("/health", func(c echo.Context) error {
			return c.String(http.StatusOK, "ok")
		})

		handler := graphqlHandler.New(&graphqlHandler.Config{
			Schema:        initSchema(),
			Pretty:        true,
			FormatErrorFn: gqlerrors.FormatError,
		})

		e.POST("/graphql", echo.WrapHandler(handler))

		// Start server
		go func() {
			if err := e.Start(viper.GetString("graphql_server.address")); err != nil && err != http.ErrServerClosed {
				e.Logger.Fatal("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	},
}
