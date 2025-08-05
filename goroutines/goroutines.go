package main

import (
	"fmt"
	"go_tutorials/different-packages/helper" // <-- A mágica acontece aqui!
	"time"
)

// Package level variables
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // -> Slice (dinamic index values)

// Works like class from java
type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsletter bool
}

func main() {

	greetUsers()

	fmt.Println("----------------------------------------")

	// Infinite loop
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("our conference is booked out. Come back next year")
				break
			}
		} else {

			if !isValidName {
				println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				println("email adress doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				println("number of tickets you entered is invalid")
			}
		}
	}

	city := "London"

	switch city {
	case "New York":
		// some code here
	case "Singapore":
		// some code here
	case "London":
		// some code here
	case "Berlin":
		// some code here
	case "Mexico City":
		// some code here
	case "Hong Kong":
		// some code here
	default:
		fmt.Println("No valid city selected")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// funct <name-of-function>(<parameter <parameter-type>>) <return-type>{}
func getFirstNames() []string {
	// ForEach loop
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, email)
	fmt.Println("#####################")
}
