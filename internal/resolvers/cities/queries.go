package cities

import (
	"time"

	"github.com/NunChatSpace/gqlgen-test/database/tables"
	"github.com/NunChatSpace/gqlgen-test/graph/model"
	"gorm.io/gorm"
)

func GetList(db *gorm.DB) (result []*model.City, err error) {
	var cityModel []tables.City
	tx := db.Find(&cityModel)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// res2B, _ := json.Marshal(cityModel[0:2])
	// fmt.Println(string(res2B))

	for _, row := range cityModel {
		mapped := mapCity(row)
		result = append(result, mapped)
	}

	return result, nil
}

func FindByIDs(db *gorm.DB, ids []int) (result []*model.City, err []error) {
	var cityTable []tables.City
	tx := db.Find(&cityTable)
	if tx.Error != nil {
		return nil, []error{tx.Error}
	}

	for _, row := range cityTable {
		result = append(result, mapCity(row))
	}

	return result, nil
}

func FindByStateIDs(db *gorm.DB, ids []int) (result [][]*model.City, err []error) {
	var cities []tables.City
	tx := db.Find(&cities, "state_id IN ?", ids)

	return result, []error{tx.Error}
}

// func cityInOrder(keys []int, items []*model.City) []*model.City {
// 	mp := make(map[int]*model.City, len(items))
// 	for _, item := range items {
// 		mp[item.ID] = item
// 	}
// 	result := make([]*model.City, len(keys))
// 	for i, id := range keys {
// 		result[i] = mp[id]
// 	}

// 	return result
// }

func mapCity(city tables.City) *model.City {
	return &model.City{
		ID:        int(city.ID),
		Name:      city.Name,
		CreatedAt: city.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: city.UpdatedAt.UTC().Format(time.RFC3339),
		State: &model.State{
			ID: int(city.StateID),
		},
	}
}
