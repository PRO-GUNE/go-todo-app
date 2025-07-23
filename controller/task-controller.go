package controller

import (
	"go-todo-app/entity"
	"go-todo-app/service"
)

type TaskController interface {
	GetAllTasks() ([]entity.Task, error)
	GetTaskByID(id int) (entity.Task, error)
	CreateTask(task entity.Task) (entity.Task, error)
	UpdateTask(id int, task entity.Task) (entity.Task, error)
	DeleteTask(id int) error
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
func (c *taskController) GetAllTasks() ([]entity.Task, error) {
	return c.service.GetAllTasks()
}

func (c *taskController) GetTaskByID(id int) (entity.Task, error) {
	return c.service.GetTaskByID(id)
}

func (c *taskController) CreateTask(task entity.Task) (entity.Task, error) {
	return c.service.CreateTask(task)
}

func (c *taskController) UpdateTask(id int, task entity.Task) (entity.Task, error) {
	return c.service.UpdateTask(id, task)
}

func (c *taskController) DeleteTask(id int) error {
	return c.service.DeleteTask(id)
}
