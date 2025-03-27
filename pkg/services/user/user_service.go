package services

import (
	"gss-backend/pkg/models"
	pointsRepo "gss-backend/pkg/repositories/points"
	userRepo "gss-backend/pkg/repositories/user"
)


func NewUserService(userRepo userRepo.IUserRepository, pointsRepo pointsRepo.IPointsRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		pointsRepo: pointsRepo,
	}
}

// Register a new user and create a new points record for the user
func (s *UserService) Create(user *models.User) (*models.User, error) {
	// Create a new user
	user, err := s.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	// Create a new points record for the user
	_, err = s.pointsRepo.Create(user.ID)

	if err != nil {
		return nil, err
	}

	return user, nil
	
}

// Find all users (used for developing purposes)
func (s *UserService) FindAll() (*[]models.User, error) {
	return s.userRepo.FindAll()
}

// Find a user by their ID (used for developing purposes)
func (s *UserService) FindByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}





