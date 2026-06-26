package auth

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

func validateSignup(form UserSignupForm) map[string]string {
	fieldErrors := make(map[string]string)

	if isBlank(form.Name) {
		fieldErrors["name"] = "Name cannot be empty"
	}

	if isBlank(form.Email) {
		fieldErrors["email"] = "Email cannot be empty"
	} else if !validEmail(form.Email) {
		fieldErrors["email"] = "Invalid email address"
	}

	if isBlank(form.Password) {
		fieldErrors["password"] = "Password cannot be empty"
	} else if lessThanMinChars(form.Password, 8) {
		fieldErrors["password"] = "Password must be at least 8 characters"
	} else if exceedMaxBytes(form.Password, 72) {
		fieldErrors["password"] = "Password cannot be longer than 72 bytes"
	}

	return fieldErrors
}

func isBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}

func exceedMaxBytes(s string, length int) bool {
	return len(s) > length
}

func lessThanMinChars(s string, length int) bool {
	return utf8.RuneCountInString(s) < length
}

func validEmail(email string) bool {
	emailRx := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRx.MatchString(email)
}
