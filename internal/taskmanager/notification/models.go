package notification

//User defines the structure for the user
type Notification struct {
	ID      *int64 `json:"id,omitempty"`
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

type Notifications []Notification
