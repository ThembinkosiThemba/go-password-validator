# go-password-validator
A simple go script to validate your password!!!

### Get Started
```
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
```
author: Thembinkosi
