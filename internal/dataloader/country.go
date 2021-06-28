package dataloader

import (
	"context"
	"fmt"
	"time"

	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/countries"
	"gorm.io/gorm"
)

type countryCtxKey struct{}

// GetCountryLoader gives a country loader from context
func GetCountryLoader(ctx context.Context) *model.CountryLoader {
	return ctx.Value(&countryCtxKey{}).(*model.CountryLoader)
}

func buildCountryLoader(db *gorm.DB) *model.CountryLoader {
	fmt.Println("buildCountryLoader")
	return model.NewCountryLoader(model.CountryLoaderConfig{
		Fetch: func(keys []int) ([]*model.Country, []error) {
			return countries.FindByIDs(db, keys)
		},
		Wait:     10 * time.Millisecond,
		MaxBatch: 100,
	})
}
