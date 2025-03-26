package repositories

import (
	"gss-backend/pkg/models"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *PostgresUserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *PostgresUserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	result := r.db.First(&user, id)
	return user, result.Error
}