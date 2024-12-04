package helper

import (
	"fmt"
	"strings"
)

// demo package scope

func ValidateUserInput(firstName string, lastName string, emailAddress string, userTickets uint) bool {
	// validate names, names have to be > than 2 chars
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// checks if string contains an @
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0
	if !isValidName {
		fmt.Printf(("invalid first or last name\n"))
		return false
	} else if !isValidEmail {
		fmt.Printf(("invalid email\n"))
		return false
	} else if !isValidTicketNumber {
		fmt.Printf(("invalid ticket number\n"))
		return false
	}
	return true
}
