package cmd

import (
	"fmt"

	"gophercises/task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong: ", err.Error())
			return
		}
		if len(tasks) == 0 {
			fmt.Println("you have no tasks to complete")
			return
		}
		fmt.Print("You have the following tasks")
		for i, task := range tasks {
			fmt.Printf("\n%d. %s", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
