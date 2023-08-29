package main

import (
	"fmt"
	// "golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

func isUppercase(r rune) bool {
	return unicode.IsUpper(r)
}

func isLowercase(r rune) bool {
	return unicode.IsLower(r)
}

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isSpecialChar(r rune) bool {
	specialChars := "!@#$%^&*()_+[]{}|;:,.<>?~"
	return strings.ContainsRune(specialChars, r)
}

// func HashPassowrd(password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return string(bytes)
// }

func isPasswordValid(password string) bool {
	// Check the minimum length
	if len(password) < 8 {
		return false
	}

	// Check for uppercase and lowercase letters, digits, and special characters
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		if isUppercase(char) {
			hasUppercase = true
		}
		if isLowercase(char) {
			hasLowercase = true
		}
		if isDigit(char) {
			hasDigit = true
		}
		if isSpecialChar(char) {
			hasSpecialChar = true
		}
	}

	// Check if all criteria are met
	return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}

func main() {
	var password string

	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	if isPasswordValid(password) {
		fmt.Println("Password is valid.")
		fmt.Println(password)
	} else {
		fmt.Println("Password is not valid / password to weak!.")
	}
}
