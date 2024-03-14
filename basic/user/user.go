package users

// User represents a user entity with fields for ID, name, and email.
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
