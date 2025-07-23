package main

import (
	"go-todo-app/controller"
	"go-todo-app/initializers"
	"go-todo-app/service"

	"github.com/gin-gonic/gin"
)

var (
	taskService    service.TaskService
	taskController controller.TaskController
)

func init() {
	// Initialize the task service and controller
	taskService = service.New()
	taskController = controller.New(taskService)

	// Load environment variables
	if err := initializers.LoadEnvVariables(); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}

	// Connect to the database
	if err := initializers.ConnectToDB(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}
}

func main() {
	router := gin.Default()

	// Define home route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Go Todo App!",
		})
	})

	// Define routes for task operations
	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.POST("/tasks", taskController.CreateTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	// Start the server on port 8080
	router.Run()
}
