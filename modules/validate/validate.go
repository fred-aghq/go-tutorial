package validate

import (
	"booking/modules/input"
	"strings"
)

func validateName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func validateTicketNumbers(userTickets uint, remainingTickets uint) bool {
	return userTickets <= remainingTickets
}

func ValidateUserInput(userData input.UserData, remainingTickets uint) (bool, bool, bool) {
	emailValid := validateEmail(userData.Email)
	nameValid := validateName(userData.FirstName, userData.LastName)
	ticketsValid := validateTicketNumbers(userData.NumberOfTickets, remainingTickets)

	return nameValid, emailValid, ticketsValid
}
