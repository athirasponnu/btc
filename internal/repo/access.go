package repo

import (
	"database/sql"
)

type AccessRepo struct {
	db *sql.DB
}

type AccessRepoImply interface {
}

// NewUserRepo
func NewAccessRepo(db *sql.DB) AccessRepoImply {
	return &AccessRepo{db: db}
}
