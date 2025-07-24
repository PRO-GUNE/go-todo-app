package main

import (
	"go-todo-app/controller"
	"go-todo-app/initializers"
	"go-todo-app/middleware"

	"github.com/gin-gonic/gin"
)

var (
	taskController controller.TaskController
	userController controller.UserController
)

func init() {
	// Load environment variables
	if err := initializers.LoadEnvVariables(); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}

	// Connect to the database
	if err := initializers.ConnectToDB(); err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Sync the database schema
	// initializers.SyncDatabase()

	// Initialize the task service and controller
	taskController = controller.NewTaskController()
	// Initialize the user controller
	userController = controller.NewUserController()
}

func main() {
	router := gin.Default()

	// Define home route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the Go Todo App!",
		})
	})

	// Define routes for user operations
	router.POST("/signup", userController.Signup)
	router.POST("/login", userController.Login)

	// Define routes for task operations
	router.GET("/tasks", taskController.GetAllTasks)
	router.GET("/tasks/:id", taskController.GetTaskByID)
	router.POST("/tasks", taskController.CreateTask)
	router.PUT("/tasks/:id", taskController.UpdateTask)
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	// Authentication routes
	authGroup := router.Group("/auth", middleware.RequireAuth())

	// Define routes for user profile operations
	authGroup.GET("/profile", userController.GetUserProfile)
	authGroup.DELETE("/profile", userController.DeleteUser)

	// Start the server on port 8080
	router.Run()
}
