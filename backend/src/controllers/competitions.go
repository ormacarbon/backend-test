package controllers

import (
	"github.com/joaooliveira247/backend-test/src/repositories"
)

type CompetitionsController struct {
	UserRepository        repositories.UsersRepository
	CompetitionRepository repositories.CompetitionsRepository
	PointRepository       repositories.PointsRepository
}

func NewCompetitionsController(
	userRepo repositories.UsersRepository,
	compRepo repositories.CompetitionsRepository,
	pointRepo repositories.PointsRepository,
) *CompetitionsController {
	return &CompetitionsController{userRepo, compRepo, pointRepo}
}
