package tables

import (
	"fmt"

	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name string
}

var country Country

func GetCountryTable() TableInterface {
	return country
}

func (c Country) NewTable(DB *gorm.DB) (err error) {
	fmt.Println("New Country")
	err = DB.AutoMigrate(&c)
	if err != nil {
		return err
	}
	if DB.Migrator().HasTable(&c) {
		err = DB.Migrator().DropTable(&c)
		if err != nil {
			return err
		}
	}

	err = DB.Migrator().CreateTable(&c)
	if err != nil {
		return err
	}
	return nil
}

func (c Country) DropTable(DB *gorm.DB) (err error) {
	if DB.Migrator().HasTable(&c) {
		err = DB.Migrator().DropTable(&c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db Country) FindByIDs(DB *gorm.DB, ids []int) ([]Country, error) {
	var countries []Country
	tx := DB.Find(&countries, ids)

	// res2B, _ := json.Marshal(countries[0])
	// fmt.Println(string(res2B))

	return countries, tx.Error
}
