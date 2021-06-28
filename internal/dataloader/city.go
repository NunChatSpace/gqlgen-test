package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/cities"
	"gorm.io/gorm"
)

type cityCtxKey struct{}
type citiesByStateCtxKey struct{}

func GetCityLoader(ctx context.Context) *model.CityLoader {
	res2B, _ := json.Marshal(ctx.Value(&cityCtxKey{}))
	fmt.Println(string(res2B))

	return ctx.Value(&cityCtxKey{}).(*model.CityLoader)
}

func buildCityLoader(db *gorm.DB) *model.CityLoader {
	fmt.Println("buildCityLoader")
	return model.NewCityLoader(model.CityLoaderConfig{
		Fetch: func(keys []int) ([]*model.City, []error) {
			return cities.FindByIDs(db, keys)
		},
		Wait:     10 * time.Millisecond,
		MaxBatch: 100,
	})
}

// GetCitiesByStateLoader gives a states by country loader
func GetCitiesByStateLoader(ctx context.Context) *model.CitiesByStateLoader {
	return ctx.Value(&citiesByStateCtxKey{}).(*model.CitiesByStateLoader)
}

func buildCitiesByStateLoader(db *gorm.DB) *model.CitiesByStateLoader {
	fmt.Println("buildCitiesByStateLoader")
	return model.NewCitiesByStateLoader(model.CitiesByStateLoaderConfig{
		Fetch: func(keys []int) ([][]*model.City, []error) {
			return cities.FindByStateIDs(db, keys)
		},
		Wait:     10 * time.Millisecond,
		MaxBatch: 100,
	})
}
