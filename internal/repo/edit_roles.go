package repo

import (
	"backend-code-base-template/internal/entities"
	"database/sql"
	"errors"
	"fmt"
)

func (access *AccessRepo) UpdateRoles(updateDetails entities.Role, id int) (sql.Result, error) {
	var EditRoleQ = `UPDATE roles SET `
	if updateDetails.Name == "" && updateDetails.Description == "" {
		return nil, errors.New("enter the feild to update")
	}
	if updateDetails.Name != "" {
		EditRoleQ = fmt.Sprintf(" %s role_name = '%s' ", EditRoleQ, updateDetails.Name)
	}
	if updateDetails.Description != "" {
		EditRoleQ = fmt.Sprintf(" %s ,description = '%s' ", EditRoleQ, updateDetails.Description)
	}
	EditRoleQ = fmt.Sprintf("%s WHERE id = %d", EditRoleQ, id)
	result, err := access.db.Exec(EditRoleQ)
	return result, err
}
