package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/backend-test/src/models"
	"github.com/joaooliveira247/backend-test/src/repositories"
	"github.com/joaooliveira247/backend-test/src/utils"
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

func (ctrl *UsersController) CreateUser(ctx *gin.Context) {
	var user models.Users

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid request body", "details": err.Error()},
		)
		return
	}

	if err := user.Validate(); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"error": "invalid request body", "details": err.Error()},
		)
		return
	}

	code, err := utils.GenerateAffiliateCode()

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error":   "error when try create affiliate code",
				"details": err.Error(),
			},
		)
		return
	}

	user.AffiliateCode = code

	affiliateCode, err := ctrl.userRepository.Create(&user)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "database error", "details": err.Error()},
		)
		return
	}

	comp, _ := ctrl.competitionRepository.GetActiveCompetition()

	if !comp.IsEmpty() {
		ctrl.pointRepository.AddPoint(
			&models.Points{UserID: user.ID, CompetitionID: comp.ID},
		)
	}

	ctx.JSON(http.StatusCreated, gin.H{"affiliateCode": affiliateCode})
}