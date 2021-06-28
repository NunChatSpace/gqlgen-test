package dataloader

import (
	"context"
	"fmt"
	"time"

	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/states"
	"gorm.io/gorm"
)

type stateCtxKey struct{}
type statesByCountryCtxKey struct{}

// GetStateLoader gives a state loader from context
func GetStateLoader(ctx context.Context) *model.StateLoader {
	return ctx.Value(&stateCtxKey{}).(*model.StateLoader)
}

// GetStatesByCountryLoader gives a states by country loader
func GetStatesByCountryLoader(ctx context.Context) *model.StatesByCountryLoader {
	return ctx.Value(&statesByCountryCtxKey{}).(*model.StatesByCountryLoader)
}

func buildStatesByCountryLoader(db *gorm.DB) *model.StatesByCountryLoader {
	fmt.Println("buildStatesByCountryLoader")
	return model.NewStatesByCountryLoader(model.StatesByCountryLoaderConfig{
		Fetch: func(keys []int) ([][]*model.State, []error) {
			return states.FindByCountryIDs(db, keys)
		},
		Wait:     10 * time.Millisecond,
		MaxBatch: 100,
	})
}

func buildStateLoader(db *gorm.DB) *model.StateLoader {
	fmt.Println("buildStateLoader")
	return model.NewStateLoader(model.StateLoaderConfig{
		Fetch: func(keys []int) ([]*model.State, []error) {
			return states.FindByIDs(db, keys)
		},
		Wait:     10 * time.Millisecond,
		MaxBatch: 100,
	})
}
