package main

import (
	"go-todo-app/initializers"
	"go-todo-app/model"
)

func init() {
	// load environment variables
	if err := initializers.LoadEnvVariables(); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}

	// connect to the database
	if err := initializers.ConnectToDB(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func main() {
	initializers.DB.AutoMigrate(&model.Task{})
}
