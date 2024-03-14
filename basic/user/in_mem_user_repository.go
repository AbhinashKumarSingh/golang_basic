package users

import (
	"fmt"
	"sync"
)

// InMemoryUserRepository is an in-memory implementation of UserRepository.
type InMemoryUserRepository struct {
	users map[int]User
	mu    sync.RWMutex
}

// NewInMemoryUserRepository creates a new instance of InMemoryUserRepository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]User),
	}
}

// Create inserts a new user into the in-memory data store.
func (repo *InMemoryUserRepository) Create(user User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.users[user.ID] = user
	return nil
}

// Read retrieves a user from the in-memory data store by ID.
func (repo *InMemoryUserRepository) Read(id int) (User, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	user, ok := repo.users[id]
	if !ok {
		return User{}, fmt.Errorf("user with ID %d not found", id)
	}
	return user, nil
}

// Update updates an existing user in the in-memory data store by ID.
func (repo *InMemoryUserRepository) Update(id int, user User) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, ok := repo.users[id]
	if !ok {
		return fmt.Errorf("user with ID %d not found", id)
	}
	repo.users[id] = user
	return nil
}

// Delete removes a user from the in-memory data store by ID.
func (repo *InMemoryUserRepository) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	_, ok := repo.users[id]
	if !ok {
		return fmt.Errorf("user with ID %d not found", id)
	}
	delete(repo.users, id)
	return nil
}
