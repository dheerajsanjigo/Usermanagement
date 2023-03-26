package validators

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	// Define a regular expression to match against the email address
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// Test the regular expression against the email address
	return regex.MatchString(email)
}
