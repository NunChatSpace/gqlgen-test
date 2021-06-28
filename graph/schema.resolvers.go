package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NunChatSpace/gqlgen-test/database"
	"github.com/NunChatSpace/gqlgen-test/graph/generated"
	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"github.com/NunChatSpace/gqlgen-test/internal/dataloader"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/cities"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/countries"
	"github.com/NunChatSpace/gqlgen-test/internal/resolvers/states"
)

func (r *cityResolver) State(ctx context.Context, obj *model.City) (*model.State, error) {
	panic(fmt.Errorf("not implemented, cityResolver State"))
}

func (r *countryResolver) States(ctx context.Context, obj *model.Country) ([]*model.State, error) {
	panic(fmt.Errorf("not implemented, countryResolver States"))
}

func (r *queryResolver) City(ctx context.Context, id int) (*model.City, error) {
	// panic(fmt.Errorf("not implemented, queryResolver City"))
	return dataloader.GetCityLoader(ctx).Load(id)
}

func (r *queryResolver) Cities(ctx context.Context) ([]*model.City, error) {
	return cities.GetList(database.GetDB())
}

func (r *queryResolver) State(ctx context.Context, id int) (*model.State, error) {
	panic(fmt.Errorf("not implemented, queryResolver State"))
}

func (r *queryResolver) States(ctx context.Context) ([]*model.State, error) {
	return states.GetList(database.GetDB())
}

func (r *queryResolver) Country(ctx context.Context, id int) (*model.Country, error) {
	panic(fmt.Errorf("not implemented, queryResolver Country"))
}

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	return countries.GetList(database.GetDB())
}

func (r *stateResolver) Country(ctx context.Context, obj *model.State) (*model.Country, error) {
	panic(fmt.Errorf("not implemented, stateResolver Country"))
}

func (r *stateResolver) Cities(ctx context.Context, obj *model.State) ([]*model.City, error) {
	panic(fmt.Errorf("not implemented, stateResolver Cities"))
}

// City returns generated.CityResolver implementation.
func (r *Resolver) City() generated.CityResolver { return &cityResolver{r} }

// Country returns generated.CountryResolver implementation.
func (r *Resolver) Country() generated.CountryResolver { return &countryResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// State returns generated.StateResolver implementation.
func (r *Resolver) State() generated.StateResolver { return &stateResolver{r} }

type cityResolver struct{ *Resolver }
type countryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type stateResolver struct{ *Resolver }
