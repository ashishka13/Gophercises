package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercises/task/db"
	"strconv"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) { //cmd is of cobra command type
		fmt.Println("do called")
		var ids []int              //ids array is generated to accept all the do numbers from user
		for _, arg := range args { //iterate over them
			id, err := strconv.Atoi(arg) //equivalent to ParseInt(s, 10, 0), converted to type int.
			//strconv converts strings into basic datatypes
			if err != nil {
				fmt.Println("failed to parse the argument", arg)
			} else {
				ids = append(ids, id) //append user do list to ids array
			}
		}
		tasks, err := db.AllTasks() //call allTasks() which will
		if err != nil {             //return tasks array
			fmt.Print("something went wrong", err)
			return
		}
		for _, id := range ids { //iterate over final id list
			if id <= 0 || id > len(tasks) { //check if the id to be processed is in the valid range
				fmt.Println("invalid task number: ", id)
				continue
			}
			task := tasks[id-1]            //the array counting starts from zero
			err := db.DeleteTask(task.Key) //call delete task with to mark the task as do
			if err != nil {                //above only Key is given and not the value
				fmt.Printf("failed to mark %d as complete %s : ", id, err)
			} else { //if no error then delete keys is successful
				fmt.Printf("marked as complete %d : ", id)
			}
		}
		fmt.Print(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
