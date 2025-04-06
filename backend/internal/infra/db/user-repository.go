package db

import (
	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) output_ports.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(user entities.User) error {
	model := UserToModel(user)
	return r.db.Save(&model).Error
}

func (r *UserGormRepository) FindByEmail(email string) (*entities.User, error) {
	var model UserModel
	err := r.db.Where("email = ?", email).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	user := model.ToDomain()
	return &user, nil
}

func (r *UserGormRepository) FindByID(id string) (*entities.User, error) {
	var model UserModel
	err := r.db.Where("id = ?", id).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	user := model.ToDomain()
	return &user, nil
}

func (r *UserGormRepository) FindByInviteCode(inviteCode string) (*entities.User, error) {
	var model UserModel
	err := r.db.Where("invite_code = ?", inviteCode).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	user := model.ToDomain()
	return &user, nil
}

func (r *UserGormRepository) FindUsersOrderedByPoints(page int, limit int) ([]entities.User, error) {
	var models []UserModel
	err := r.db.Order("points desc").Offset(page * limit).Limit(limit).Find(&models).Error
	if err != nil {
		return nil, err
	}

	users := make([]entities.User, len(models))
	for i, model := range models {
		users[i] = model.ToDomain()
	}

	return users, nil
}

func (r *UserGormRepository) ResetAllScores() error {
	return r.db.Model(&UserModel{}).Update("points", 0).Error
}
