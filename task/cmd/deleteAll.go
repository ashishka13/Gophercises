package cmd

import (
	"fmt"
	"gophercises/task/db"
	"strconv"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "dela",
	Short: "deletes all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, _ := db.AllTasks() //get all the tasks saved into tasks variable
		var dcmd *cobra.Command   //created variable of cobra type

		fmt.Print("\n tasks to delete ", len(tasks), "\n")

		var doArgs []string
		if len(tasks) == 0 {
			fmt.Print("no tasks to delete ")
			return
		}
		for i := range tasks {
			i = i + 1             //tasks array start from zero and 0 is invalid task
			ii := strconv.Itoa(i) //"do" taskes only srting arguments
			fmt.Print(ii)
			doArgs = append(doArgs, ii) //create our do array for passing to original do function
		}
		doCmd.Run(dcmd, doArgs) //pass our argument array to do function
	},
}

func init() {
	RootCmd.AddCommand(delCmd)
}
