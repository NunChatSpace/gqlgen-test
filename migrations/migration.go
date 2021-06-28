package migrations

import (
	"path"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func NewMigrate(db *gorm.DB) (*migrate.Migrate, error) {
	database, err := db.DB()
	if err != nil {
		return nil, err
	}

	dbDriver, err := postgres.WithInstance(database, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	_, filename, _, _ := runtime.Caller(0)
	migrationPath := path.Join(path.Dir(filename), "")
	migrator, err := migrate.NewWithDatabaseInstance("file:///"+migrationPath, "postgres", dbDriver)
	if err != nil {
		return nil, err
	}

	return migrator, nil
}
