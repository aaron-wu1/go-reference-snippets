package main

// IO
import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// global level scope
// use capitalize var name
// eg. ConferenceName

// package level scope
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
var bookings = make([]UserData, 0)

// // with make guarrentees acesss to a map
// var bookings = make([]map[string]string, 0)

// userdata struct
type UserData struct {
	firstName    string
	lastName     string
	emailAddress string
	userTickets  uint
}

// init as nil, bad practice
// var bookings []map[string]string

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {
	if remainingTickets == 0 {
		// end program
		fmt.Println("Our conference is booked out. Come back next year.")
		// break
	}

	firstName, lastName, emailAddress, userTickets := getUserInput()

	if !helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets) {
		// continue
	}

	// validate ticket count
	for userTickets > remainingTickets {
		fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		fmt.Printf("Enter a new ticket number\n")
		fmt.Scan(&userTickets)
	}

	bookTicket(userTickets, firstName, lastName, emailAddress)
	// add to wait group
	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, emailAddress)

	firstNames := getFirstNames()
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
	// }
	// waits for all threads to be done
	wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)
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
	// bookings = append(bookings, firstName+" "+lastName)

	// using map
	// create empty map with make
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailAddress"] = emailAddress
	// // converts uint to uint64 to string
	// userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//using struct
	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		userTickets:  userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
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

func sendTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	// assume is a slow process simulate with sleep
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###########################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, emailAddress)
	fmt.Println("###########################")
	// removes thread from wait group
	wg.Done()
}
