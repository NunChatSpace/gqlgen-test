package commands

import (
	"github.com/NunChatSpace/gqlgen-test/http"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run server",
	RunE:  serveExecute,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serveExecute(cmd *cobra.Command, args []string) (err error) {
	http.ListentAndServe()
	return nil
}
