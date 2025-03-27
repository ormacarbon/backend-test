package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

// Concrete implementation of the IPointsRepository interface
func NewPostgresPointsRepository(db *gorm.DB) *PostgresPointsRepository {
	return &PostgresPointsRepository{db: db}
}

func (r *PostgresPointsRepository) Create(user_id uint) (*models.Points, error) {
	points := models.Points{UserId: user_id, Points: 1}
	result := r.db.Create(&points)
	return &points, result.Error
}

func (r *PostgresPointsRepository) Update(user_id uint) (*models.Points, error) {
	var points models.Points
	
	result := r.db.Where("user_id = ?", user_id).First(&points)

	if result.Error != nil {
		return nil, result.Error
	}

	points.Points += 1

	result = r.db.Save(&points)

	if result.Error != nil {
		return nil, result.Error
	}
	
	return &points, result.Error
}

func (r *PostgresPointsRepository) FindByUserID(user_id uint) (*models.Points, error) {
	var points models.Points
	result := r.db.Where("user_id = ?", user_id).First(&points)
	return &points, result.Error
}

func (r *PostgresPointsRepository) CreateLeaderboard() (*[]models.Points, error) {
	var points []models.Points
	result := r.db.Order("points desc").Find(&points).Limit(10)
	return &points, result.Error
}