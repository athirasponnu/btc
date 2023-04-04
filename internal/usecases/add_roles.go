package usecases

import (
	"backend-code-base-template/internal/entities"
)

func (access *AccessUseCases) InsertRoles(details entities.Role) (int, error) {

	return access.repo.InsertRoles(details)
}
