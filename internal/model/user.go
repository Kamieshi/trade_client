package model

// User model
type User struct {
	ID      string `json:"id,omitempty" readonly:"true"`
	Name    string `json:"name,omitempty"`
	Balance int64  `json:"balance"`
}
