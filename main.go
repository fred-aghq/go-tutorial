package main

import (
	"fmt"
	"strings"
)

const maxTickets uint = 50
const conferenceName string = "Go Conference"

var remainingTickets uint = maxTickets

var bookings []string

func printHello() {
	fmt.Printf("Welcome to the %s booking system.\n", conferenceName)
	fmt.Println("Use this system to book attendance.")
	fmt.Println()
	fmt.Printf("There are currently %d tickets available.\n\n", remainingTickets)
}

func bookTickets() bool {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	// isValidName := validateName(firstName, lastName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Printf("Thanks %s. How many tickets would you like?:\n", firstName)

	fmt.Scan(&userTickets)

	for userTickets > remainingTickets {
		fmt.Println("That exceeds the number of remaining tickets.")
		fmt.Printf("Please enter a valid amount. There are %d tickets remaining:\n", remainingTickets)
		fmt.Scan(&userTickets)
	}

	return processBooking(userTickets, firstName, lastName, email)
}

func processBooking(userTickets uint, firstName string, lastName string, email string) bool {
	if validateName(firstName, lastName) && validateEmail(email) {
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("User %s booked %d tickets.\n", email, userTickets)
		return true
	}

	return false
}

func validateName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func main() {
	firstNames := []string{}

	for remainingTickets > 0 {
		printHello()

		if !bookTickets() {
			fmt.Println("Sorry, your information was invalid. Please try again.")
			continue
		}

		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstName := names[0]
			firstNames = append(firstNames, firstName)
		}

		fmt.Printf("There are %d tickets remaining.\n", remainingTickets)
		fmt.Println()
		fmt.Printf("All bookings: \n\n%v\n\n", firstNames)
	}

	fmt.Println("Unfortunately there are no tickets remaining.")
}
