package main

import (
	userRepo "gss-backend/pkg/repositories/user"
	userReferralRepo "gss-backend/pkg/repositories/user_referral"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	UserRepository userRepo.IUserRepository
	UserReferralRepository userReferralRepo.IUserReferralRepository
}

func NewRepositoryContainer(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		UserRepository: userRepo.NewPostgresUserRepository(db),
		UserReferralRepository: userReferralRepo.NewPostgresUserReferralRepository(db),
	}
}
