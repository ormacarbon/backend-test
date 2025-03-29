package repositories

import (
	"github.com/joaooliveira247/backend-test/src/models"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return UsersRepository{db}
}

func (repository *UsersRepository) Create(user *models.Users) (string, error) {
	result := repository.db.Create(&user)

	if err := result.Error; err != nil {
		return "", err
	}

	return user.AffiliateCode, nil
}

func (repository *UsersRepository) GetUserByAffiliateCode(
	code string,
) (models.Users, error) {
	var user models.Users

	if err := repository.db.First(&user, "affiliate_code = ?", code).Error; err != nil {
		return models.Users{}, err
	}

	return user, nil
}
