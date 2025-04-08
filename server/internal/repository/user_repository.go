package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByShareCode(shareCode string) (*models.User, error)
	UpdatePoints(user *models.User, points int) error
	GetLeaderboard(filters Filters) ([]models.User, int64, error)
}

type Filters struct {
	Sort   string
	Search string
	Page   int
	Limit  int
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

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
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

func (r *userRepository) GetLeaderboard(filters Filters) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := r.db.Model(&models.User{})

	if filters.Search != "" {
		query = query.Where("name ILIKE ? OR email ILIKE ?", "%"+filters.Search+"%", "%"+filters.Search+"%")
	}

	query.Count(&total)

	switch filters.Sort {
	case "name":
		query = query.Order("name asc")
	case "email":
		query = query.Order("email asc")
	default:
		query = query.Order("points desc")
	}

	offset := (filters.Page - 1) * filters.Limit
	err := query.Offset(offset).Limit(filters.Limit).Find(&users).Error

	return users, total, err
}
