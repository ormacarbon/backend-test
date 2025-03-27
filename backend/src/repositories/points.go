package repositories

import "gorm.io/gorm"


type PointsRepository struct {
	db *gorm.DB
}
