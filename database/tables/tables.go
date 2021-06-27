package tables

import "gorm.io/gorm"

type TableInterface interface {
	NewTable(DB *gorm.DB) error
	DropTable(DB *gorm.DB) error
	InsertTableData(DB *gorm.DB) error
}

type Tables struct {
	TableList []TableInterface
}

var tables Tables

func InitTables(DB *gorm.DB) (err error) {
	tables.TableList = append(tables.TableList, GetCityTable())
	tables.TableList = append(tables.TableList, GetCountryTable())
	tables.TableList = append(tables.TableList, GetStateTable())

	for _, t := range tables.TableList {
		err = t.NewTable(DB)
		if err != nil {
			return err
		}
	}

	return nil
}
