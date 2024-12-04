package main

// IO
import (
	"fmt"
	"strings"
)

// globals
// syntatic sugar for init var with type inference
var conferenceName = "Go Conference"

// explicit types
const conferenceTickets int = 50

var remainingTickets uint = 50

// Array:
// assigns inital values with {}
// multiple {"name0", "name1", "name 2"}
// initial size 50
// var bookings  = [50]string{}
// defines booking type
// var bookings [50]string

// Slice - array without explicit size
// dynamic array
var bookings []string

func main() {

	greetUsers()

	for {
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}

		firstName, lastName, emailAddress, userTickets := getUserInput()

		// validate names, names have to be > than 2 chars
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		// checks if string contains an @
		isValidEmail := strings.Contains(emailAddress, "@")
		isValidTicketNumber := userTickets > 0
		if !isValidName {
			fmt.Printf(("invalid first or last name\n"))
			continue
		} else if !isValidEmail {
			fmt.Printf(("invalid email\n"))
			continue
		} else if !isValidTicketNumber {
			fmt.Printf(("invalid ticket number\n"))
			continue
		}
		// validate ticket count
		for userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			fmt.Printf("Enter a new ticket number\n")
			fmt.Scan(&userTickets)
		}

		bookTicket(userTickets, firstName, lastName, emailAddress)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}
}

func greetUsers() {
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Fields(booking)[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets uint
	// get user input, assigns input to address val
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&emailAddress)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	// add to booking
	// bookings[0] = firstName + " " + lastName
	// adding using slice
	bookings = append(bookings, firstName+" "+lastName)

	// // has spaces, represents the rest of the array, doesn't splice out whitespace
	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Array Type: %T\n", bookings)
	// fmt.Printf("Array length: %v\n", len(bookings))

	// // slice, note no white space in print
	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice Type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	// decrement tickets
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v booked %v tickets.\nYou will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
