package service

import (
	"go-todo-app/entity"
)

type TaskService interface {
	GetAllTasks() ([]entity.Task, error)
	GetTaskByID(id int) (entity.Task, error)
	CreateTask(task entity.Task) (entity.Task, error)
	UpdateTask(id int, task entity.Task) (entity.Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	tasks []entity.Task
}

// New creates a new instance of TaskService - Constructor function
func New() TaskService {
	return &taskService{
		tasks: []entity.Task{},
	}
}

// Implementing TaskService Interface methods
func (s *taskService) GetAllTasks() ([]entity.Task, error) {
	return s.tasks, nil
}

func (s *taskService) GetTaskByID(id int) (entity.Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return entity.Task{}, nil // Return an empty Task if not found
}

func (s *taskService) CreateTask(task entity.Task) (entity.Task, error) {
	task.ID = task.GetTaskID(s.tasks) // Assign a new ID
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s *taskService) UpdateTask(id int, task entity.Task) (entity.Task, error) {
	for i, t := range s.tasks {
		if t.ID == id {
			task.ID = id // Ensure the ID remains the same
			s.tasks[i] = task
			return task, nil
		}
	}
	return entity.Task{}, nil // Return an empty Task if not found
}

func (s *taskService) DeleteTask(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return nil // Return nil if the task was not found
}
