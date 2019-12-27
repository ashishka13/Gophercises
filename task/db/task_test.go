package db

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

var home, _ = homedir.Dir()
var dbPath = filepath.Join(home, "tasks.db")

func TestInit(t *testing.T){
	dbPath1 := ""
	err := Init(dbPath1)

	if err == nil{
		t.Error("database path not found  :",err)
	}
}

func TestCreateTask(t *testing.T) {
	Init(dbPath)
	task := "test one"

	id, _ := CreateTask(task)
	assert.NotEqual(t, id, nil)

	MyCloseDb()
}

func TestCreateTaskWithErr(t *testing.T) {
	Init(dbPath)
	task := "test one"
	MyCloseDb()
	_, err := CreateTask(task)
	assert.NotEqual(t, err, nil)
}

func TestAllTasks(t *testing.T) {
	Init(dbPath)

	tasks1, _ := AllTasks()
	before := len(tasks1)

	task1 := "test one"
	task2 := "test two"

	CreateTask(task1)
	CreateTask(task2)

	tasks, err := AllTasks()

	if err != nil {
		t.Error(len(tasks))
	}
	after := len(tasks)

	if before == after {
		t.Error("ssdvsv")
	}
	MyCloseDb()
}

func TestAllTasksErr(t *testing.T) {
	Init(dbPath)

	task1 := "test one"
	task2 := "test two"

	CreateTask(task1)
	CreateTask(task2)

	MyCloseDb()
	tasks, err := AllTasks()
	if err == nil {
		t.Error(len(tasks))
	}
}
func TestDeleteTask(t *testing.T) {
	Init(dbPath)
	task := "test one del"

	CreateTask(task)
	MyCloseDb()
	err := DeleteTask(5678)
	if err == nil {
		t.Error("failed to mark as complete : ", err)
	}
}

func TestDeleteTaskNoErr(t *testing.T) {
	Init(dbPath)
	task := "test one del"

	CreateTask(task)
	err := DeleteTask(1)
	if err != nil {
		t.Error("failed to mark as complete : ", err)
	}
	MyCloseDb()
}
