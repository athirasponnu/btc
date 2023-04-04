package repo

import (
	"backend-code-base-template/internal/entities"
	"log"
)

func (access *AccessRepo) GetRolesById(id int) (entities.Roles, error) {
	var info entities.Roles
	GetRolesQ := `SELECT id,role_name,description FROM roles WHERE id= $1 `
	rows := access.db.QueryRow(GetRolesQ, id)
	err := rows.Scan(&info.Id, &info.RoleName, &info.Description)
	if err != nil {
		log.Printf("error occurred while reading the database rows: %v", err)
		return entities.Roles{}, err
	}

	return info, nil

}
