package adapters

import (
	"0AlexZhong0/goblog/internal/users/domain/user"
	"context"
	"sync"
)

type MemoryUserRepository struct {
	users map[string]*user.User

	lock    *sync.RWMutex
	factory user.Factory
}

func (m *MemoryUserRepository) GetUser(ctx context.Context, userId string) (*user.User, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	result, exists := m.users[userId]
	if !exists {
		return nil, user.UserNotFound
	}

	return result, nil
}

func (m *MemoryUserRepository) CreateUser(ctx context.Context, in *user.NewUserInput) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	for id := range m.users {
		if id == in.Id {
			return user.UserAlreadyExists
		}
	}

	newUser, err := m.factory.NewUser(in)
	if err != nil {
		return err
	}

	m.users[in.Id] = newUser
	return nil
}

func (m *MemoryUserRepository) UpdateUser(ctx context.Context, in *user.UpdateUserInput) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	u, exists := m.users[in.Id]

	if !exists {
		return user.UserNotFound
	}

	u.SetAvatar(in.Avatar)

	if err := u.SetLastName(in.LastName); err != nil {
		return err
	}

	if err := u.SetFirstName(in.FirstName); err != nil {
		return err
	}

	m.users[in.Id] = u

	return nil
}

func (m *MemoryUserRepository) DeleteUser(ctx context.Context, userId string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	_, exists := m.users[userId]
	if !exists {
		return user.UserNotFound
	}

	delete(m.users, userId)
	return nil
}

func NewMemoryUserRepository(f user.Factory) *MemoryUserRepository {
	return &MemoryUserRepository{
		factory: f,
		lock:    &sync.RWMutex{},
		users:   map[string]*user.User{},
	}
}
