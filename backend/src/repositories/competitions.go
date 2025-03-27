package repositories

import (
	"gorm.io/gorm"
)

type CompetitionsRepository struct {
	db *gorm.DB
}
