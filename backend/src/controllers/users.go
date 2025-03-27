package controllers

import (
	"github.com/joaooliveira247/backend-test/src/repositories"
)

type UsersController struct {
	userRepository        repositories.UsersRepository
	competitionRepository repositories.CompetitionsRepository
	pointRepository       repositories.PointsRepository
}
