package states

import (
	"time"

	"github.com/NunChatSpace/gqlgen-test/database/tables"
	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"gorm.io/gorm"
)

func GetList(db *gorm.DB) (result []*model.State, err error) {
	var stateModel []tables.State
	tx := db.Find(&stateModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, row := range stateModel {
		result = append(result, mapState(row))
	}

	return result, nil
}

func FindByIDs(db *gorm.DB, ids []int) (result []*model.State, err []error) {
	var states []tables.State
	tx := db.Find(&states)
	if tx.Error != nil {
		return nil, []error{tx.Error}
	}

	for _, row := range states {
		result = append(result, mapState(row))
	}

	return result, err
}

func FindByCountryIDs(db *gorm.DB, ids []int) (result [][]*model.State, err []error) {
	var states []tables.State
	tx := db.Find(&states, "country_id IN ?", ids)

	return result, []error{tx.Error}
}

func mapState(state tables.State) *model.State {
	return &model.State{
		ID:        int(state.ID),
		Name:      state.Name,
		CreatedAt: state.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: state.UpdatedAt.UTC().Format(time.RFC3339),
		Country: &model.Country{
			ID: int(state.CountryID),
		},
	}
}
