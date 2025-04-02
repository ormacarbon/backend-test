package utils

import "github.com/google/uuid"

func GenerateReferralCode() string {
	return uuid.NewString()
}