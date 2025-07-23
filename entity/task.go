package entity

type Task struct {
	ID           int    `json:"id"`
	Description  string `json:"description"`
	Completed    bool   `json:"completed"`
	Priority     int    `json:"priority"`     // Higher value = higher priority
	Duration     int    `json:"duration"`     // Estimated time in minutes
	Dependencies []int  `json:"dependencies"` // IDs of todos that must be done first
}

func (s *Task) GetTaskID(tasks []Task) int {
	var maxID int
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1 // Return the next available ID
}
