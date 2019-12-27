package cmd

import (
	"gophercises/task/db"
	"testing"

	"github.com/spf13/cobra"
)

func TestList(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command                    //created variable of cobra type
	addCmd.Run(cmd, []string{"ashishasihah"}) //provided above variable to addcmd.Run
	listCmd.Run(cmd, []string{""})
	db.MyCloseDb()
}

func TestListSomethingWrong(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command                    //created variable of cobra type
	addCmd.Run(cmd, []string{"ashishasihah"}) //provided above variable to addcmd.Run
	db.MyCloseDb()
	listCmd.Run(cmd, []string{""})
}

func TestListTasksZero(t *testing.T) {
	db.Init(dbPath)
	var cmd *cobra.Command        //created variable of cobra type
	delCmd.Run(cmd, []string{""}) //provided above variable to addcmd.Run
	listCmd.Run(cmd, []string{""})
	db.MyCloseDb()
}
