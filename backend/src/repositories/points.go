package repositories

import (
	"github.com/joaooliveira247/backend-test/src/models"
	"gorm.io/gorm"
)

type PointsRepository struct {
	db *gorm.DB
}

func NewPointsRepository(db *gorm.DB) PointsRepository {
	return PointsRepository{db}
}

func (repository *PointsRepository) AddPoint(point *models.Points) error {
	if err := repository.db.Create(point).Error; err != nil {
		return err
	}

	return nil
}
