package db

import (
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/google/uuid"
)

type UserModel struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name       string     `gorm:"column:name;not null"`
	Email      string     `gorm:"column:email;unique;not null"`
	Phone      string     `gorm:"column:phone;not null"`
	Password   string     `gorm:"column:password;not null"`
	InviteCode string     `gorm:"column:invite_code;unique;not null"`
	InvitedBy  *uuid.UUID `gorm:"column:invited_by"`
	Points     int        `gorm:"column:points;not null;default:1"`
}

func (model UserModel) ToDomain() entities.User {
	email, _ := object_values.NewEmail(model.Email)
	phone, _ := object_values.NewPhoneNumber(model.Phone)
	hashedPass := object_values.NewPasswordFromHash(model.Password)

	return entities.LoadUser(
		model.ID,
		model.Name,
		email,
		hashedPass,
		phone,
		model.InviteCode,
		model.InvitedBy,
		model.Points,
	)
}

func UserToModel(user entities.User) UserModel {
	return UserModel{
		ID:         user.ID(),
		Name:       user.Name(),
		Email:      user.Email().Value(),
		Phone:      user.Phone().Value(),
		Password:   user.Password().Hash(),
		InviteCode: user.InviteCode(),
		InvitedBy:  user.InvitedBy(),
		Points:     user.Points(),
	}
}

func (UserModel) TableName() string {
	return "users"
}
