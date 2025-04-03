package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByShareCode(shareCode string) (*models.User, error)
	UpdatePoints(user *models.User, points int) error
	GetTopUsers(limit int) ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByShareCode(shareCode string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("share_code = ?", shareCode).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdatePoints(user *models.User, points int) error {
	return r.db.Model(user).Update("points", points).Error
}

func (r *userRepository) GetTopUsers(limit int) ([]models.User, error) {
	var users []models.User
	err := r.db.Order("points desc").Limit(limit).Select("name, points").Find(&users).Error
	return users, err
}
