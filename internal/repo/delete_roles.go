package repo

import (
	"database/sql"
)

func (access *AccessRepo) DeleteRoles(id int) (sql.Result, error) {
	DeleteRoleQ := `DELETE FROM roles WHERE id = $1`
	result, err := access.db.Exec(DeleteRoleQ, id)
	return result, err
}
