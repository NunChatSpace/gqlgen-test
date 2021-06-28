package dataloader_handler

import (
	"net/http"

	"github.com/NunChatSpace/gqlgen-test/internal/dataloader"
)

func Handler() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return dataloaderHandler(h)
	}
}

func dataloaderHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := dataloader.SetupLoaders(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
