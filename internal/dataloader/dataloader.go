package dataloader

import (
	"context"

	"github.com/NunChatSpace/gqlgen-test/database"
)

func SetupLoaders(ctx context.Context) context.Context {
	db := database.GetDB()

	c := context.WithValue(ctx, &cityCtxKey{}, buildCityLoader(db))
	c = context.WithValue(c, &citiesByStateCtxKey{}, buildCitiesByStateLoader(db))

	c = context.WithValue(c, &stateCtxKey{}, buildStateLoader(db))
	c = context.WithValue(c, &statesByCountryCtxKey{}, buildStatesByCountryLoader(db))

	c = context.WithValue(c, &countryCtxKey{}, buildCountryLoader(db))

	return c
}
