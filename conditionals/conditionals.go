package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := [50]string{} // -> Array (fixed index values)
	bookings2 := []string{}  // -> Slice (dinamic index values)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("----------------------------------------")

	// Infinite loop

	for {
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

		if userTickets < 50 {
			// It needs to be the same type (uint - uint)
			remainingTickets = remainingTickets - userTickets

			bookings[0] = firstName + " " + lastName
			bookings2 = append(bookings2, firstName+" "+lastName)
			fmt.Println("-----------------------------------------------------")
			fmt.Println("Array bookings [50]string")

			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Array type: %T\n", bookings)
			fmt.Printf("Array length: %v\n", len(bookings))

			fmt.Println("-----------------------------------------------------")
			fmt.Println("Slice bookings2 []string")

			fmt.Printf("The whole array: %v\n", bookings2)
			fmt.Printf("The first value: %v\n", bookings2[0])
			fmt.Printf("Array type: %T\n", bookings2)
			fmt.Printf("Array length: %v\n", len(bookings2))

			fmt.Println("-----------------------------------------------------")
			fmt.Println("Result")

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining\n", remainingTickets)

			// ForEach loop
			firstNames := []string{}
			for _, booking := range bookings2 {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are all our bookins: %v\n", bookings2)
			fmt.Printf("The first name of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("our conference is booked out. Come back next year")
				break
			}
		} else {

			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}

}
