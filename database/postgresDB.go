package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func NewPostgresDB() (err error) {
	cfg := postgres.Config{
		DSN: "host=localhost user=postgres password=postgres dbname=gql_test port=15432 sslmode=disable TimeZone=Asia/Shanghai", //nolint:lll
	}

	DB, err = gorm.Open(postgres.New(cfg))
	if err != nil {
		return err
	}
	return nil
}
