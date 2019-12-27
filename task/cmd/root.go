package cmd

import (
	"github.com/spf13/cobra"
)

//RootCmd use
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a cli task manager",
}
