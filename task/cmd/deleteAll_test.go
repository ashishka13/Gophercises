package cmd

import (
	"gophercises/task/db"
	"testing"

	"github.com/spf13/cobra"
)

func TestDelAll(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command
	delCmd.Run(cmd, []string{""})
	delCmd.Run(cmd, []string{""})
	db.MyCloseDb()
}
