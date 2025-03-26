package repositories

import "gorm.io/gorm"

type UsersRepository struct {
	db *gorm.DB
}
