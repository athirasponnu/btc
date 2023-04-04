package repo

import (
	"backend-code-base-template/internal/entities"
)

func (access *AccessRepo) InsertRoles(details entities.Role) (int, error) {
	var id int
	NewRoleQ := `INSERT INTO roles 
				(role_name,description) 
				VALUES ( $1 , $2 ) RETURNING id`
	rows := access.db.QueryRow(NewRoleQ, details.Name, details.Description)
	err := rows.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, err
}
