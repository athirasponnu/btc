package usecases

import (
	"backend-code-base-template/internal/repo"
)

type AccessUseCases struct {
	repo repo.AccessRepoImply
}
type AccessUseCaseImply interface {
}

// NewUserUseCases
func NewRoleUseCases(access repo.AccessRepoImply) AccessUseCaseImply {
	return &AccessUseCases{
		repo: access,
	}
}
