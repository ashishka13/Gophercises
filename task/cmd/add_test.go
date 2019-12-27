package cmd

import (
	"gophercises/task/db"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var home, _ = homedir.Dir()
var dbPath = filepath.Join(home, "tasks.db")

func TestAddCmd(t *testing.T) {
	db.Init(dbPath)

	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})

	db.MyCloseDb()
}

func TestAddCmdErrs(t *testing.T) {
	db.Init(dbPath)
	db.MyCloseDb()
	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})
}
