package controllers

import (
	"github.com/joaooliveira247/backend-test/src/repositories"
)

type UsersController struct {
	userRepository        repositories.UsersRepository
	competitionRepository repositories.CompetitionsRepository
	pointRepository       repositories.PointsRepository
}

func NewUsersController(
	userRepo repositories.UsersRepository,
	compRepo repositories.CompetitionsRepository,
	pointRepo repositories.PointsRepository,
) *UsersController {
	return &UsersController{userRepo, compRepo, pointRepo}
}