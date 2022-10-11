run-server:
	go run app/main.go graphql

graphql:
	go run github.com/99designs/gqlgen generate
