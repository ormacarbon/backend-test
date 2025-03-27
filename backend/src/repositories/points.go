package repositories

import "gorm.io/gorm"

type PointsRepository struct {
	db *gorm.DB
}

func NewPointsRepository(db *gorm.DB) PointsRepository {
	return PointsRepository{db}
}
