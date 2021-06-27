package commands

import (
	"fmt"

	"github.com/NunChatSpace/gqlgen-test/database"
	"github.com/NunChatSpace/gqlgen-test/database/tables"
	"github.com/spf13/cobra"
)

var DbCmd = &cobra.Command{
	Use:   "db",
	Short: "Manipulate the database",
}

var DbMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	RunE:  dbMigrate,
}

func init() {
	rootCmd.AddCommand(DbCmd)
	DbCmd.AddCommand(DbMigrateCmd)
}

func dbMigrate(cmd *cobra.Command, args []string) (err error) {
	err = database.NewPostgresDB()
	if err != nil {
		return err
	}

	db := database.GetDB()
	err = tables.InitTables(db)
	if err != nil {
		return err
	}

	fmt.Println("Migrate Done")
	return nil
}
