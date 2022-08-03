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

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Printf("Thanks %s. How many tickets would you like?:\n", firstName)

	var userTickets uint
	fmt.Scan(&userTickets)

	return processBooking(userTickets, firstName, lastName, email)
}

func processBooking(userTickets uint, firstName string, lastName string, email string) bool {
	if userTickets < remainingTickets {
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("User %s booked %d tickets.\n", email, userTickets)
		return true
	}

	fmt.Println("That exceeds the number of remaining tickets")

	return false
}

func main() {
	firstNames := []string{}

	for {
		if remainingTickets == 0 {
			fmt.Println("Unfortunately there are no tickets remaining.")
			break
		}

		printHello()

		if !bookTickets() {
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
}
