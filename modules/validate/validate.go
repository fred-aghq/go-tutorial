package validate

import "strings"

func validateName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func validateTicketNumbers(userTickets uint, remainingTickets uint) bool {
	return userTickets <= remainingTickets
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	emailValid := validateEmail(email)
	nameValid := validateName(firstName, lastName)
	ticketsValid := validateTicketNumbers(userTickets, remainingTickets)

	return nameValid, emailValid, ticketsValid
}
