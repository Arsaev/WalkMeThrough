package entity

// User - struct that represents a user in the system
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

// Password - struct that stores the hashed password and the salt
type Password struct {
	UserID         string `json:"user_id"`
	HashedPassword string `json:"hashed_password"`
}
