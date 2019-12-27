package main

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"gophercises/task/cmd"
	"gophercises/task/db"
	"os"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
