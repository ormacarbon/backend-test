package utils

import (
	"net"
	"regexp"
	"strings"
)

func EmailValidator(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	if !re.MatchString(email) {
		return false
	}

	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	mxRecords, err := net.LookupMX(parts[1])
	return err == nil && len(mxRecords) > 0
}

func IsPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^\+55\d{11}$`)

	return re.MatchString(phone)
}
