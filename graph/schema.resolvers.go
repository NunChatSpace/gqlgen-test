package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/NunChatSpace/gqlgen-test/graph/generated"
	"github.com/NunChatSpace/gqlgen-test/graph/model"
)

func (r *cityResolver) State(ctx context.Context, obj *model.City) (*model.State, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *countryResolver) States(ctx context.Context, obj *model.Country) ([]*model.State, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) City(ctx context.Context, id int) (*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cities(ctx context.Context) ([]*model.City, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) State(ctx context.Context, id int) (*model.State, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) States(ctx context.Context) ([]*model.State, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Country(ctx context.Context, id int) (*model.Country, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Countries(ctx context.Context) ([]*model.Country, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stateResolver) Country(ctx context.Context, obj *model.State) (*model.Country, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stateResolver) Cities(ctx context.Context, obj *model.State) ([]*model.City, error) {
	panic(fmt.Errorf("not implemented"))
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
