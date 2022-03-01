package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")

	for remainingTickets > 0 {
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

		isInputValid := true
		isInputValid = isInputValid && len(firstName) >= 2
		isInputValid = isInputValid && len(lastName) >= 2
		isInputValid = isInputValid && strings.Contains(email, "@")
		isInputValid = isInputValid && userTickets > 0 && userTickets > remainingTickets

		if !isInputValid {
			fmt.Println("Invalid Input. Please try again")
			continue
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
			firstName, lastName, userTickets, email,
		)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Fields(booking)[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	}
}
