package repo

import (
	"backend-code-base-template/internal/entities"
	"fmt"
	"log"
)

// GetRoles
func (access *AccessRepo) GetRoles(limit int, offset int) ([]entities.Roles, error) {
	var info []entities.Roles
	GetRolesQ := `SELECT id,role_name,description FROM roles LIMIT $1 OFFSET $2 `
	rows, err := access.db.Query(GetRolesQ, limit, offset)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a entities.Roles
		err = rows.Scan(&a.Id, &a.RoleName, &a.Description)
		if err != nil {
			log.Printf("error occurred while reading the database rows: %v", err)
			break
		}
		info = append(info, a)
	}
	return info, nil

}
