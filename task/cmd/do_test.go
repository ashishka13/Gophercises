package cmd

import (
	"gophercises/task/db"
	"strconv"
	"testing"

	"github.com/spf13/cobra"
)

func TestDo(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})
	doCmd.Run(cmd, []string{"1"})
	db.MyCloseDb()
}

func TestDoFailedParse(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})
	doCmd.Run(cmd, []string{"1 hy %& *$#"})
	db.MyCloseDb()
}
func TestDoSomethingWrong(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})
	db.MyCloseDb()
	doCmd.Run(cmd, []string{"1"})
}

func TestDoWithErr(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command
	addCmd.Run(cmd, []string{"task1"})
	myNegative := strconv.Itoa(-1)
	doCmd.Run(cmd, []string{myNegative})
	db.MyCloseDb()
}
