package controllers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joaooliveira247/backend-test/src/models"
	"github.com/joaooliveira247/backend-test/src/repositories"
	"github.com/joaooliveira247/backend-test/src/services"
	"gorm.io/gorm"
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
				services.SendEmail(
					user.Email,
					"Carbon Offset Competition",
					fmt.Sprintf(
						"+1 point using your affiliate code: %s",
						code,
					),
				)
			}
		}
		ctx.JSON(http.StatusOK, compCheck)
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (ctrl *CompetitionsController) CloseCompetition(ctx *gin.Context) {
	compID, err := uuid.Parse(ctx.Query("ID"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format", "details": err.Error(),
		})
		return
	}

	if err := ctrl.CompetitionRepository.CloseCompetition(compID); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(
				http.StatusNotFound,
				gin.H{
					"error":   "competition not found",
					"details": err.Error(),
				},
			)
			return
		}

		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "closed competition error", "details": err.Error()},
		)
		return
	}

	winners, err := ctrl.CompetitionRepository.GetCompetitionReport(compID)

	if err != nil {
		log.Printf("error pick winners: %v", err)
	}

	for i, report := range winners {
		services.SendEmail(
			report.Email,
			"Carbon Offset Competition Ended",
			fmt.Sprintf(
				"Congratulations Carbon Offset Competition: %s ended, your position was %d",
				compID,
				i + 1,
			),
		)
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (ctrl *CompetitionsController) GetCompetitionReport(ctx *gin.Context) {
	compID, err := uuid.Parse(ctx.Query("ID"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id format", "details": err.Error(),
		})
		return
	}

	comp, _ := ctrl.CompetitionRepository.GetCompetitionByID(compID)

	if comp.Status {
		ctx.JSON(http.StatusConflict, gin.H{"error": "competition not closed"})
		return
	}

	reports, err := ctrl.CompetitionRepository.GetCompetitionReport(compID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "gen competition report", "details": err.Error(),
		})
		return
	}

	if reports == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "competiton not found"})
		return
	}

	ctx.JSON(http.StatusOK, reports)
}
