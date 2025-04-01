package main

import (
	emailService "gss-backend/pkg/services/email"
	userService "gss-backend/pkg/services/user"
	userReferralService "gss-backend/pkg/services/user_referral"
)
	

type ServiceContainer struct {
	EmailService emailService.IEmailService
	UserService userService.IUserService
	UserReferralService userReferralService.IUserReferralService
}

func NewServiceContainer(repoContainer *RepositoryContainer, emailConfig emailService.EmailConfig) *ServiceContainer {
	emailService := emailService.NewEmailService(emailConfig)
	userService := userService.NewUserService(repoContainer.UserRepository, repoContainer.UserReferralRepository, emailService)
	userReferralService := userReferralService.NewUserReferralService(repoContainer.UserRepository, repoContainer.UserReferralRepository, emailService)

	return &ServiceContainer{
		EmailService: emailService,
		UserService: userService,
		UserReferralService: userReferralService,
	}
}