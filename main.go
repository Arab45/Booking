package main

import (
	"Booking-App/helper"
	"Booking-App/premium"
	"fmt"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var remainingTickets int = 50
// var slices = make([]UserData, 0)
var slices []UserData;

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets int
}

func main() {

	premium.PremiumUser()

	greetUser()

	// slices = append(slices, "John Smith");
	// fmt.Println("slice in Go is also an array", slices[0]);

	fmt.Println(&conferenceName)
	fmt.Println(&remainingTickets)

	fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		firstName, lastName, email, tickets := getUserInput()

		var middleName *string
		name := "Alice"

		middleName = &name
		fmt.Println("Value of middleName is", *middleName)

		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, tickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			leftTickets := bookingTicket(remainingTickets, firstName, lastName, email, tickets)

			firstNames := getFirstName()
			fmt.Printf("First names of bookings are %v\n", firstNames)

			if leftTickets == 0 {
				fmt.Println("All tickets are sold out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is too short, try again!")
			}
			if !isValidEmail {
				fmt.Println("Your email doesn't contain @ sign, try again!")
			}
			if !isValidTicketNumber {
				fmt.Println("Your ticket number is invalid, try again!")
			}
			break
		}

		var location string
		fmt.Println("Input your prefered event location:")
		fmt.Scan(&location)

		switch location {
		case "London":
			fmt.Println("You are in London")
		case "New York":
			fmt.Println("You are in New York")
		case "Paris":
			fmt.Println("You are in Paris")
		default:
			fmt.Println("You are not in London, New York or Paris")
		}

	}
}

func greetUser() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Println("We are glad to have you here!")
	fmt.Println("We hope you enjoy the conference and have a great time!")
}

func getFirstName() []string {

	firstNames := []string{}
	for _, values := range slices {
		firstNames = append(firstNames, values.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var tickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&tickets)

	return firstName, lastName, email, tickets
}

func bookingTicket(remainingTickets int, firstName string, lastName string, email string, tickets int) int {
	var leftTickets int = remainingTickets - tickets

	//create a map
	// var slices []UserData
	userData := UserData{
        firstName:       firstName,
        lastName:        lastName,
        email:           email,
        numberOfTickets: tickets,
    }

	slices = append(slices, userData)
	fmt.Printf("List  of bookings is: %v \n", slices)

	fmt.Printf("User %v %v with email %v booked %v tickets\n", firstName, lastName, email, tickets)

	return leftTickets

}
