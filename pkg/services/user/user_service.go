package services

import (
	"gss-backend/pkg/models"
	repositories "gss-backend/pkg/repositories/user"
)


func NewUserService(repository repositories.IUserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	return s.repository.Create(user)
}

func (s *UserService) FindAll() (*[]models.User, error) {
	return s.repository.FindAll()
}

func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.repository.FindByID(id)
}





