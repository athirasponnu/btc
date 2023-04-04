package usecases

import "backend-code-base-template/internal/entities"

// Get roles
func (access *AccessUseCases) GetRoles(limit int, offset int) ([]entities.Roles, error) {
	return access.repo.GetRoles(limit, offset)
}
