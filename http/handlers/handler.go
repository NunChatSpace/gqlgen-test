package handlers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/NunChatSpace/gqlgen-test/graph"
	"github.com/NunChatSpace/gqlgen-test/graph/generated"
	"github.com/NunChatSpace/gqlgen-test/http/handlers/dataloader_handler"
	"github.com/justinas/alice"
)

func BuildRootHandler() http.Handler {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{},
	}))

	dataloaderHandler := dataloader_handler.Handler()
	return alice.New(
		dataloaderHandler,
	).Then(srv)
}
