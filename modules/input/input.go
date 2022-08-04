package input

import "fmt"

func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Printf("Thanks %s. How many tickets would you like?:\n", firstName)
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
