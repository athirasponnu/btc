package usecases

import "database/sql"

func (access *AccessUseCases) DeleteRoles(id int) (sql.Result, error) {

	return access.repo.DeleteRoles(id)
}
