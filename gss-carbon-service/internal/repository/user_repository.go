package repository

import (
	"context"
	"errors"

	"github.com/icl00ud/backend-test/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByReferralToken(ctx context.Context, token string) (*model.User, error)
	GetReferrals(ctx context.Context, offset, limit int) ([]*model.User, error)
	UpdateUserPoints(ctx context.Context, id uint, points int) error
	GetTopUsers(ctx context.Context, limit int) ([]model.User, error)
	CleanTable(ctx context.Context) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByReferralToken(ctx context.Context, token string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Where("referral_token = ?", token).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetReferrals(ctx context.Context, offset, limit int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).
		Where("referred_by IS NOT NULL").
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserPoints(ctx context.Context, id uint, points int) error {
	result := r.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Update("points", gorm.Expr("points + ?", points))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *userRepository) GetTopUsers(ctx context.Context, limit int) ([]model.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Order("points desc").Limit(limit).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) CleanTable(ctx context.Context) error {
	tx := r.db.Session(&gorm.Session{AllowGlobalUpdate: true})
	return tx.WithContext(ctx).Delete(&model.User{}).Error
}
