package model

// Task defines the structure for the tasks
type Task struct {
	ID         *int64  `json:"id,omitempty"`
	UserID     int64  `json:"user_id"`
	AssignedTo *int64  `json:"assigned_to,omitempty"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	CreatedAt  *string `json:"created_at,omitempty"`
	DoneAt     *string `json:"done_at,omitempty"`
}
