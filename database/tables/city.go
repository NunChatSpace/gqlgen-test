package tables

import (
	"fmt"

	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name      string
	State     State
	StateID   uint
	Country   Country
	CountryID uint
}

var city City

func GetCityTable() TableInterface {
	return city
}

func (c City) NewTable(DB *gorm.DB) (err error) {
	fmt.Println("New City")
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

	DB.Migrator().CreateConstraint(&c, "State")
	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&c, "fk_state_id")
	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&c, "Country")
	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&c, "fk_country_id")
	if err != nil {
		return err
	}

	return nil
}

func (c City) DropTable(DB *gorm.DB) (err error) {
	if DB.Migrator().HasTable(&c) {
		err = DB.Migrator().DropTable(&c)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c City) FindByIDs(DB *gorm.DB, ids []int) ([]City, error) {
	var cities []City
	tx := DB.Find(&cities, ids)

	return cities, tx.Error
}

func (c City) FindByStateIDs(DB *gorm.DB, ids []int) (result []City, err error) {
	var cities []City
	tx := DB.Find(&cities)

	return result, tx.Error
}
