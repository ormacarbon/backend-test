package repositories

import (
	"github.com/jpeccia/go-backend-test/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByReferralCode(code string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *userRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.db.Where("email= ?", email).First(&user).Error
	return &user, err
}

func (u *userRepository) FindUserByReferralCode(code string) (*models.User, error) {
	var user models.User
	err := u.db.Where("referral_code = ?", code).First(&user).Error
	return &user, err
}

func (u *userRepository) UpdateUser(user *models.User) error {
	return u.db.Save(user).Error
}
