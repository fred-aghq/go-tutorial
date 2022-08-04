package main

import (
	"booking/modules/input"
	"booking/modules/validate"
	"fmt"
	"time"
)

const maxTickets uint = 50

var remainingTickets uint = maxTickets

const conferenceName string = "Go Conference"

var bookings []input.UserData = make([]input.UserData, 0)

func greetUsers() {
	fmt.Println("##################################################################")
	fmt.Printf("Welcome to the %s booking system.\n", conferenceName)
	fmt.Println("Use this system to book attendance.")
	fmt.Println()
	fmt.Printf("There are currently %d tickets available.\n\n", remainingTickets)
	fmt.Println("##################################################################")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}

func bookTickets(userData input.UserData) {
	remainingTickets = remainingTickets - userData.NumberOfTickets
	bookings = append(bookings, userData)

	fmt.Printf("Thanks, %s! You have successfully booked %d tickets\n\n", userData.FirstName, userData.NumberOfTickets)
}

func sendTicket(userData input.UserData) {
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%d tickets for %s %s", userData.NumberOfTickets, userData.FirstName, userData.LastName)
	fmt.Println()
	fmt.Println("########################################################################")
	fmt.Printf("Sending ticket:\n %s\nto email address %s\n", ticket, userData.Email)
	fmt.Println("########################################################################")
	fmt.Println()
}

func main() {
	for remainingTickets > 0 {
		greetUsers()

		userData := input.GetUserInput()

		validName, validEmail, validTickets := validate.ValidateUserInput(userData, remainingTickets)

		switch false {
		case validEmail:
			println("Your email appears to be invalid.")
			continue
		case validTickets:
			println("Number of tickets selected is invalid")
			continue
		case validName:
			println("Your name does not appear to have been typed properly")
			continue
		}

		bookTickets(userData)
		sendTicket(userData)

		fmt.Printf("\n\nFirst names: %v\n", getFirstNames())
		fmt.Printf("Bookings:\n%v\n", bookings)

		fmt.Printf("There are %d tickets remaining.\n", remainingTickets)
		fmt.Println()
	}

	fmt.Println("Unfortunately there are no tickets remaining.")
}
