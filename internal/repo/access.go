package repo

import (
	"backend-code-base-template/internal/entities"
	"database/sql"
)

type AccessRepo struct {
	db *sql.DB
}

type AccessRepoImply interface {
	GetRoles(int, int) ([]entities.Roles, error)
	GetRolesById(int) (entities.Roles, error)
	DeleteRoles(int) (sql.Result, error)
	UpdateRoles(entities.Role, int) (sql.Result, error)
	InsertRoles(entities.Role) (int, error)
}

// NewUserRepo
func NewAccessRepo(db *sql.DB) AccessRepoImply {
	return &AccessRepo{db: db}
}
