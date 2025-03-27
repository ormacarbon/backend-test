package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaooliveira247/backend-test/src/models"
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

func (ctrl *CompetitionsController) Create(ctx *gin.Context) {
	compCheck, _ := ctrl.CompetitionRepository.GetActiveCompetition()

	if compCheck.IsEmpty() {
		id, err := ctrl.CompetitionRepository.Create(
			&models.Competitions{Status: true},
		)

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "create competition", "details": err.Error()},
			)
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"id": id})
		return
	}

	ctx.JSON(
		http.StatusConflict,
		gin.H{
			"error":   "competition already activated",
			"details": compCheck.ID,
		},
	)
}

func (ctrl *CompetitionsController) GetCompetition(ctx *gin.Context) {
	// main request for main screen
	compCheck, _ := ctrl.CompetitionRepository.GetActiveCompetition()

	if !compCheck.IsEmpty() {
		if code := ctx.Query("affiliateCode"); code != "" {
			user, _ := ctrl.UserRepository.GetUserByAffiliateCode(code)

			if !user.IsEmpty() {
				ctrl.PointRepository.AddPoint(
					&models.Points{
						UserID:        user.ID,
						CompetitionID: compCheck.ID,
					},
				)
			}
		}
		ctx.JSON(http.StatusOK, compCheck)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}
