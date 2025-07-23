package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Completed    bool    `json:"completed"`
	Priority     int32   `json:"priority"`                           // Higher value = higher priority
	Duration     int32   `json:"duration"`                           // Estimated time in minutes
	Dependencies []int32 `json:"dependencies" gorm:"type:integer[]"` // IDs of todos that must be done first
}
