package main

import (
	"strings"
	"unicode"
)

type Result struct {
	Email    bool
	Password bool
}

func ValidateCredentials(email, password string) Result {
	atIndex := strings.Index(email, "@")
	dotIndex := strings.LastIndex(email, ".")

	emailValid := atIndex > 0 && (dotIndex > atIndex && dotIndex < len(email)-1)

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, c := range password {
		if unicode.IsLower(c) {
			hasLower = true
		}
		if unicode.IsUpper(c) {
			hasUpper = true
		}
		if unicode.IsNumber(c) {
			hasNumber = true
		}
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			hasSpecial = true
		}
		if hasUpper && hasLower && hasNumber && hasSpecial {
			break
		}
	}
	passwordValid := len(password) >= 8

	passwordValid = passwordValid && hasLower && hasUpper && hasNumber && hasSpecial

	return Result{
		Email:    emailValid,
		Password: passwordValid,
	}
}
