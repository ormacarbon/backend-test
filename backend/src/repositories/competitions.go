package repositories

import (
	"gorm.io/gorm"
)

type CompetitionsRepository struct {
	db *gorm.DB
}

func NewCompetiotionsRepository(db *gorm.DB) CompetitionsRepository {
	return CompetitionsRepository{db}
}
