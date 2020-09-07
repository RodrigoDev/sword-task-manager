package model

// Task defines the structure for the tasks
type Task struct {
	UserID     int64  `json:"userID"`
	AssignedTo int64  `json:"assignedTo"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CreatedAt  string `json:"createdAt,omitempty"`
	DoneAt     string `json:"doneAt,omitempty"`
}
