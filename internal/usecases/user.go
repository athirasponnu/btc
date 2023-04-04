package usecases

import (
	"backend-code-base-template/internal/entities"
	"backend-code-base-template/internal/repo"
)

type UserUseCases struct {
	repo repo.UserRepoImply
}

type UserUseCaseImply interface {
	GetUsers() []entities.User
}

// NewUserUseCases
func NewUserUseCases(userRepo repo.UserRepoImply) UserUseCaseImply {
	return &UserUseCases{
		repo: userRepo,
	}
}

// GetUsers
func (user *UserUseCases) GetUsers() []entities.User {
	return user.repo.GetUsers()
}
