package db

import (
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/google/uuid"
)

type UserModel struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string    `gorm:"column:name;not null"`
	Email    string    `gorm:"column:email;unique;not null"`
	Phone    string    `gorm:"column:phone;unique;not null"`
	Password string    `gorm:"column:password;not null"`
}

func (model UserModel) ToDomain() entities.User {
	email, _ := object_values.NewEmail(model.Email)
	phone, _ := object_values.NewPhoneNumber(model.Phone)
	hashedPass, _ := object_values.NewPassword(model.Password)

	user, _ := entities.NewUser(model.Name, email, hashedPass, phone)
	return user
}

func UserToModel(user entities.User) UserModel {
	return UserModel{
		ID:       user.ID(),
		Name:     user.Name(),
		Email:    user.Email().Value(),
		Phone:    user.Phone().Value(),
		Password: user.Password().Hash(),
	}
}

func (UserModel) TableName() string {
	return "users"
}
