package countries

import (
	"github.com/NunChatSpace/gqlgen-test/database/tables"
	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"gorm.io/gorm"
)

func GetList(db *gorm.DB) (result []*model.Country, err error) {
	var countryModel []tables.Country
	tx := db.Find(&countryModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	for _, row := range countryModel {
		result = append(result, mapCountry(row))
	}

	return result, nil
}

func FindByIDs(db *gorm.DB, ids []int) (result []*model.Country, err []error) {
	var countries []tables.Country
	tx := db.Find(&countries)
	if tx.Error != nil {
		return nil, []error{tx.Error}
	}

	for _, row := range countries {
		result = append(result, mapCountry(row))
	}

	return result, err
}

func mapCountry(country tables.Country) *model.Country {
	return &model.Country{
		ID:   int(country.ID),
		Name: country.Name,
	}
}
