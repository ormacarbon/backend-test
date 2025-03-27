package repositories

import (
	"github.com/google/uuid"
	"github.com/joaooliveira247/backend-test/src/models"
	"gorm.io/gorm"
)

type CompetitionsRepository struct {
	db *gorm.DB
}

func NewCompetiotionsRepository(db *gorm.DB) CompetitionsRepository {
	return CompetitionsRepository{db}
}

func (repository *CompetitionsRepository) Create(
	competiotion *models.Competitions,
) (uuid.UUID, error) {
	result := repository.db.Create(&competiotion)

	if err := result.Error; err != nil {
		return uuid.UUID{}, err
	}

	return competiotion.ID, nil
}

func (repository *CompetitionsRepository) GetActiveCompetition() (models.Competitions, error) {
	var competition models.Competitions

	if err := repository.db.First(&competition, "status = true").Error; err != nil {
		return models.Competitions{}, err
	}

	return competition, nil
}
