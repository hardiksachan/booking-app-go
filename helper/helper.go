package helper

import "strings"

func ValidateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint,
) bool {
	isInputValid := true
	isInputValid = isInputValid && len(firstName) >= 2
	isInputValid = isInputValid && len(lastName) >= 2
	isInputValid = isInputValid && strings.Contains(email, "@")
	isInputValid = isInputValid && userTickets > 0 && userTickets <= remainingTickets

	return isInputValid
}
