package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
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

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			printFirstNames(bookings)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if	!isValidName {
				fmt.Println("first name or last you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't containt @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

	// city := "London"

	// switch city {
	// 	case "New York":
	// 		// execute code for booking New York econference tickets
	// 	case "Singapore", "Hong Kong":
	// 		// some code here for Singapore & Hong Kong
	// 	case "London", "Berlin":
	// 		// some code here for London & Berlin
	// 	case "Mexico City":
	// 		// some code here
	// 	default:
	// 		fmt.Println("No valid city selected")
	// }
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
}