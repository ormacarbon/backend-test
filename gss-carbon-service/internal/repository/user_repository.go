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
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByReferralToken(ctx context.Context, token string) (*model.User, error)
	GetReferrals(ctx context.Context, offset, limit int) ([]*model.User, error)
	UpdateUserPoints(ctx context.Context, id uint, points int) error
	GetTopUsers(ctx context.Context, limit int) ([]model.User, error)
	CleanTable() error
}

type userRepository struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserRepository(db *gorm.DB, logger *zap.SugaredLogger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger.Named("UserRepository"),
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	return err
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warnw("User not found by ID", "id", id)
			return nil, nil
		}

		r.logger.Errorw("Failed to get user by ID", "id", id, "error", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warnw("User not found by email", "email", email)
			return nil, nil
		}

		r.logger.Errorw("Failed to get user by email", "email", email, "error", err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByReferralToken(ctx context.Context, token string) (*model.User, error) {
	var user model.User
	r.logger.Debugw("Getting user by referral token", "token", token)

	err := r.db.WithContext(ctx).Where("referral_token = ?", token).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warnw("User not found by referral token", "token", token)
			return nil, errors.New("user not found")
		}

		r.logger.Errorw("Failed to get user by referral token", "token", token, "error", err)
		return nil, err
	}

	r.logger.Debugw("Found user by referral token", "token", token, "userID", user.ID)

	return &user, nil
}

func (r *userRepository) GetReferrals(ctx context.Context, offset, limit int) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.WithContext(ctx).
		Where("referred_by IS NOT NULL").
		Order("created_at desc").
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		r.logger.Errorw("Failed to get referrals", "error", err)
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserPoints(ctx context.Context, id uint, points int) error {
	r.logger.Debugw("Updating user points", "userID", id, "pointsToAdd", points)
	result := r.db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Update("points", gorm.Expr("points + ?", points))

	if result.Error != nil {
		r.logger.Errorw("Failed to update user points", "userID", id, "error", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		r.logger.Warnw("Update user points affected 0 rows", "userID", id)
		return errors.New("user not found or no update needed")
	}

	r.logger.Debugw("Successfully updated user points", "userID", id, "rowsAffected", result.RowsAffected)
	return nil
}

func (r *userRepository) GetTopUsers(ctx context.Context, limit int) ([]model.User, error) {
	var users []model.User

	err := r.db.WithContext(ctx).Order("points desc").Limit(limit).Find(&users).Error
	if err != nil {
		r.logger.Errorw("Failed to get top users", "limit", limit, "error", err)
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CleanTable() error {
	r.logger.Debug("Cleaning user table")

	tx := r.db.Session(&gorm.Session{AllowGlobalUpdate: true})

	result := tx.Delete(&model.User{})
	if err := result.Error; err != nil {
		r.logger.Errorw("Failed to clean user table", "error", err)
		return err
	}

	r.logger.Infow("User table cleaned successfully", "rowsAffected", result.RowsAffected)
	return nil
}
