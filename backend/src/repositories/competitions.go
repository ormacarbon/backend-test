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

func (repository *CompetitionsRepository) GetCompetitionByID(competitionID uuid.UUID) (models.Competitions, error) {
	var competition models.Competitions

	if err := repository.db.First(&competition, "id = ?", competitionID).Error; err != nil {
		return models.Competitions{}, err
	}

	return competition, nil
}

func (repository *CompetitionsRepository) CloseCompetition(
	competitionID uuid.UUID,
) error {
	var competition models.Competitions

	if err := repository.db.First(&competition, "id = ?", competitionID).Error; err != nil {
		return err
	}

	competition.Status = false

	if err := repository.db.Save(&competition).Error; err != nil {
		return err
	}

	return nil
}

func (repository *CompetitionsRepository) GetCompetitionReport(
	competitionID uuid.UUID,
) ([]models.CompetitionReport, error) {
	rawQuery := `
SELECT 
    u.name, 
    u.email, 
    u.phone, 
    COUNT(*) AS points 
FROM points p
INNER JOIN users u ON p.user_id = u.id 
WHERE p.competition_id = ? 
GROUP BY u.name, u.email, u.phone 
ORDER BY points DESC 
LIMIT 10;
	`

	var reports []models.CompetitionReport

	if err := repository.db.Raw(rawQuery, competitionID).Scan(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}
