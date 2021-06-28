package tables

import (
	"fmt"

	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name      string
	Country   Country
	CountryID uint
}

var state State

func GetStateTable() TableInterface {
	return state
}

func (s State) NewTable(DB *gorm.DB) (err error) {
	fmt.Println("New State")

	err = DB.AutoMigrate(&s)
	if err != nil {
		return err
	}

	if DB.Migrator().HasTable(&s) {
		err = DB.Migrator().DropTable(&s)
		if err != nil {
			return err
		}
	}

	err = DB.Migrator().CreateTable(&s)
	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&s, "Country")
	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&s, "fk_country_id")
	if err != nil {
		return err
	}

	return nil
}

func (s State) DropTable(DB *gorm.DB) (err error) {
	if DB.Migrator().HasTable(&s) {
		err = DB.Migrator().DropTable(&s)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s State) FindByIDs(DB *gorm.DB, ids []int) ([]State, error) {
	var states []State
	tx := DB.Find(&states, ids)
	return states, tx.Error
}

func (s State) FindByCountryIDs(DB *gorm.DB, ids []int) (result []State, err error) {
	var states []State
	tx := DB.Find(&states, "country_id IN ?", ids)

	return states, tx.Error
}
