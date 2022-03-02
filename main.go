package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isInputValid := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isInputValid {
			fmt.Println("Invalid Input. Please try again")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The first names of bookings are: %v", firstNames())

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
			firstName, lastName, userTickets, email,
		)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	}
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

func firstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
	return firstNames
}

func validateUserInput(
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
	isInputValid = isInputValid && userTickets > 0 && userTickets > remainingTickets

	return isInputValid
}

func getUserInput() (
	string,
	string,
	string,
	uint,
) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
