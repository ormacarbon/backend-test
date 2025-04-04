package repository

import (
	"context"
	"errors"

	"github.com/icl00ud/backend-test/internal/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByReferralToken(ctx context.Context, token string) (*model.User, error)
	UpdateUserPoints(ctx context.Context, id uint, points int) error
	GetTopUsers(ctx context.Context, limit int) ([]model.User, error)
}

type userRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewUserRepository(db *gorm.DB, logger *zap.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger.Named("UserRepository"),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	sugar := r.logger.Sugar()
	sugar.Debugw("Creating user", "email", user.Email)

	err := r.db.WithContext(ctx).Create(user).Error

	if err != nil {
		sugar.Errorw("Failed to create user", "email", user.Email, "error", err)
	}
	return err
}

func (r *userRepository) GetUserByReferralToken(ctx context.Context, token string) (*model.User, error) {
	sugar := r.logger.Sugar()
	var user model.User
	sugar.Debugw("Getting user by referral token", "token", token)

	err := r.db.WithContext(ctx).Where("referral_token = ?", token).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sugar.Warnw("User not found by referral token", "token", token)
			return nil, errors.New("user not found")
		}

		sugar.Errorw("Failed to get user by referral token", "token", token, "error", err)
		return nil, err
	}

	sugar.Debugw("Found user by referral token", "token", token, "userID", user.ID)

	return &user, nil
}

func (r *userRepository) UpdateUserPoints(ctx context.Context, id uint, points int) error {
	sugar := r.logger.Sugar()
	sugar.Debugw("Updating user points", "userID", id, "pointsToAdd", points)
	result := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("points", gorm.Expr("points + ?", points))

	if result.Error != nil {
		sugar.Errorw("Failed to update user points", "userID", id, "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		sugar.Warnw("Update user points affected 0 rows", "userID", id)
		return errors.New("user not found or no update needed")
	}

	sugar.Debugw("Successfully updated user points", "userID", id, "rowsAffected", result.RowsAffected)
	return nil
}

func (r *userRepository) GetTopUsers(ctx context.Context, limit int) ([]model.User, error) {
	sugar := r.logger.Sugar()
	sugar.Debugw("Getting top users", "limit", limit)
	var users []model.User

	err := r.db.WithContext(ctx).Order("points desc").Limit(limit).Find(&users).Error
	if err != nil {
		sugar.Errorw("Failed to get top users", "limit", limit, "error", err)
		return nil, err
	}

	sugar.Debugw("Successfully retrieved top users", "limit", limit, "count", len(users))
	return users, nil
}
