package controller

import (
	"go-todo-app/initializers"
	"go-todo-app/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController interface {
	GetAllTasks(ctx *gin.Context)
	GetTaskByID(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
}

type taskController struct {
	db *gorm.DB
}

// New creates a new instance of TaskController - Constructor function
func New() TaskController {
	return &taskController{
		db: initializers.DB,
	}
}

// Implementing TaskController Interface methods
func (c *taskController) GetAllTasks(ctx *gin.Context) {
	var tasks []model.Task
	if err := c.db.Find(&tasks).Error; err != nil {
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
	var task model.Task
	if err := c.db.First(&task, id).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(200, task)
}

func (c *taskController) CreateTask(ctx *gin.Context) {
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task data"})
		return
	}
	if err := c.db.Create(&task).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, task)
}

func (c *taskController) UpdateTask(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	var task model.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task data"})
		return
	}
	task.ID = uint(id)
	if err := c.db.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, task)
}

func (c *taskController) DeleteTask(ctx *gin.Context) {

	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}
	if err := c.db.Delete(&model.Task{}, id).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(204, nil) // No content
}
