run-server:
	go run cmd/grip/main.go graphql

run-playground:
	go run cmd/grip/main.go graphiql

graphql:
	go run github.com/99designs/gqlgen generate
