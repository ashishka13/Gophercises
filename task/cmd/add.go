package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {

		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("something went wrong: ", err.Error())
			return
		}
		fmt.Printf("added %s to list \n", task)
		// for i, arg := range args {
		// 	fmt.Println(i, ":", arg)
		// } // this is other way of doing above printing
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
