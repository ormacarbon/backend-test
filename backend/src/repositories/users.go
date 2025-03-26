package repositories

import "gorm.io/gorm"

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return UsersRepository{db}
}
