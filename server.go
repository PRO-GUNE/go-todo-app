package main

import (
	"go-todo-app/controller"
	"go-todo-app/entity"
	"go-todo-app/initializer"
	"go-todo-app/service"

	"strconv"

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
	if err := initializer.LoadEnvVariables(); err != nil {
		panic("Failed to load environment variables: " + err.Error())
	}

	// Connect to the database
	if err := initializer.ConnectToDB(); err != nil {
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
	router.GET("/tasks", func(ctx *gin.Context) {
		tasks, err := taskController.GetAllTasks()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, tasks)
	})
	router.GET("/tasks/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid task ID"})
			return
		}
		task, err := taskController.GetTaskByID(id)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Task not found"})
			return
		}
		ctx.JSON(200, task)
	})
	router.POST("/tasks", func(ctx *gin.Context) {
		var task entity.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid task data"})
			return
		}
		createdTask, err := taskController.CreateTask(task)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, createdTask)
	})
	router.PUT("/tasks/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid task ID"})
			return
		}
		var task entity.Task
		if err := ctx.ShouldBindJSON(&task); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid task data"})
			return
		}
		task.ID = id
		updatedTask, err := taskController.UpdateTask(id, task)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, updatedTask)
	})
	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid task ID"})
			return
		}
		err = taskController.DeleteTask(id)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(204, nil) // No content
	})

	// Start the server on port 8080
	router.Run()
}
