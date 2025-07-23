package controller

import (
	"go-todo-app/entity"
	"go-todo-app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetAllTasks(ctx *gin.Context)
	GetTaskByID(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type taskController struct {
	service service.TaskService
}

// New creates a new instance of TaskController - Constructor function
func New(service service.TaskService) TaskController {
	return &taskController{
		service: service,
	}
}

// Implementing TaskController Interface methods
func (c *taskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, tasks)
}

func (c *taskController) GetTaskByID(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	task, err := c.service.GetTaskByID(id)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(200, task)
}

func (c *taskController) CreateTask(ctx *gin.Context) {
	var task entity.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task data"})
		return
	}
	createdTask, err := c.service.CreateTask(task)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, createdTask)
}

func (c *taskController) UpdateTask(ctx *gin.Context) {
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
	updatedTask, err := c.service.UpdateTask(id, task)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, updatedTask)
}

func (c *taskController) DeleteTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	err = c.service.DeleteTask(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(204, nil) // No content
}
