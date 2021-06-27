package http

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/NunChatSpace/gqlgen-test/database"
	"github.com/NunChatSpace/gqlgen-test/graph"
	"github.com/NunChatSpace/gqlgen-test/graph/generated"
	"github.com/friendsofgo/graphiql"
	"github.com/gorilla/mux"
)

const defaultPort = "3001"

func ListentAndServe() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := database.NewPostgresDB()
	if err != nil {
		panic(err)
	}
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	router := mux.NewRouter()
	router.Handle("/graphql", graphiqlHandler).Methods("GET")
	router.Handle("/graphql", srv).Methods("POST")

	log.Printf("connect to http://localhost:%s/graphql for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
