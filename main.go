package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = make([]map[string]string, 1)

func main() {

	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInput()
		isInputValid := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isInputValid {
			fmt.Println("Invalid Input. Please try again")
			continue
		}

		bookTickets(firstName, lastName, email, userTickets)

		fmt.Printf("The first names of bookings are: %v\n", firstNames())
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
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func bookTickets(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
) {
	remainingTickets = remainingTickets - userTickets

	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticketsCount"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
		firstName, lastName, userTickets, email,
	)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
