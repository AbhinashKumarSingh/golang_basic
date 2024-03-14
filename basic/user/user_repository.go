package users

// UserRepository defines the interface for user management operations.
type UserRepository interface {
	Create(user User) error
	Read(id int) (User, error)
	Update(id int, user User) error
	Delete(id int) error
}
