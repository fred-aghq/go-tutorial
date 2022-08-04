package main

import (
	"booking/modules/input"
	"booking/modules/validate"
	"fmt"
	"strconv"
)

const maxTickets uint = 50

var remainingTickets uint = maxTickets

const conferenceName string = "Go Conference"

var bookings []map[string]string = make([]map[string]string, 0)

func greetUsers() {
	fmt.Printf("Welcome to the %s booking system.\n", conferenceName)
	fmt.Println("Use this system to book attendance.")
	fmt.Println()
	fmt.Printf("There are currently %d tickets available.\n\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thanks, %s! You have successfully booked %d tickets", firstName, userTickets)
}

func main() {
	for remainingTickets > 0 {
		greetUsers()

		firstName, lastName, email, userTickets := input.GetUserInput()

		validName, validEmail, validTickets := validate.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

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

		bookTickets(firstName, lastName, email, userTickets)

		fmt.Printf("\n\nFirst names: %v\n", getFirstNames())
		fmt.Printf("Bookings:\n%v\n", bookings)

		fmt.Printf("There are %d tickets remaining.\n", remainingTickets)
		fmt.Println()
	}

	fmt.Println("Unfortunately there are no tickets remaining.")
}
