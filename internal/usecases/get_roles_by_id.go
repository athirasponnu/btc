package usecases

import "backend-code-base-template/internal/entities"

func (access *AccessUseCases) GetRolesById(id int) (entities.Roles, error) {
	return access.repo.GetRolesById(id)
}
