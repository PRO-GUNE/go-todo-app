package initializers

import (
	"go-todo-app/model"
)

func SyncDatabase() {
	DB.AutoMigrate(&model.Task{})
	DB.AutoMigrate(&model.User{})
	// Add any other models that need to be migrated here

}
