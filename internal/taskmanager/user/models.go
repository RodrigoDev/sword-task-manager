package user

//User defines the structure for the user
type User struct {
	ID       *int64  `json:"id,omitempty"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Type     *string `json:"type,omitempty"`
}

type Users []User