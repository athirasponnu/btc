package usecases

import (
	"backend-code-base-template/internal/entities"
	"backend-code-base-template/internal/repo"
	"database/sql"
)

type AccessUseCases struct {
	repo repo.AccessRepoImply
}
type AccessUseCaseImply interface {
	GetRoles(int, int) ([]entities.Roles, error)
	GetRolesById(int) (entities.Roles, error)
	DeleteRoles(int) (sql.Result, error)
	UpdateRoles(entities.Role, int) (sql.Result, error)
	InsertRoles(entities.Role) (int, error)
}

// NewUserUseCases
func NewRoleUseCases(access repo.AccessRepoImply) AccessUseCaseImply {
	return &AccessUseCases{
		repo: access,
	}
}
