package repo

import (
	"backend-code-base-template/internal/entities"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepoImply interface {
	GetUsers() []entities.User
}

// NewUserRepo
func NewUserRepo(db *sql.DB) UserRepoImply {
	return &UserRepo{db: db}
}

// GetUsers
func (user *UserRepo) GetUsers() []entities.User {
	return []entities.User{}
}
