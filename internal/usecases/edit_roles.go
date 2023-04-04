package usecases

import (
	"backend-code-base-template/internal/entities"
	"database/sql"
)

func (access *AccessUseCases) UpdateRoles(updateDetails entities.Role, id int) (sql.Result, error) {
	return access.repo.UpdateRoles(updateDetails, id)
}
