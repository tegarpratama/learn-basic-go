package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

/*
	STRUCT
	type -> mean creates a new data type with name you specify
	"Create a type called "UserData" based on a struct of firstName, lastName, ..."
*/
type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

/*
	Waits for the launched goroutine to finish
	Add: sets the number of goroutines to wait for (increases the counter by the provided number)
	Wait: Blokcs until the WaitGroup counter is 0
*/
var wg = sync.WaitGroup{}

/*
	GLOBAL VARIABEL / VARIABLE PACKAGE LEVEL
*/
var conferenceName string = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50

//1. Without struct
// var bookings = make([]map[string]string, 0) -> Empty slice of map

// 2. With Struct
var bookings = make([]UserData, 0)

/*
	Slice and array is same, but array length must set first and not dinamycly
	1. Array:
	var bookings [50]string
	Declare array and fill it with values
	var bookings = [50]string{"Tegar", "Pratama"}

	2. Slice:
	var bookings = []string{}
	bookings = [
		{ firstName: "Tegar", lastName: "Pratama" },
		{ firstName: "Tegar", lastName: "Pratama" },
	]
*/

func main() {
	greetUsers()

	// for remainingTickets > 0 && len(bookings) < 50{
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName)

			// 1. Run this func with single thread (sync)
			// sendTicket(userTickets, firstName, lastName, email)
			
			// 2. Run this func with multi thread (async)
			// "go" mean starts a new goroutine
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
		
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are : %v\n", firstNames)
			fmt.Println(getFirstNames())

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Your email is invalid")
			}

			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}

			if userTickets > remainingTickets {
				fmt.Printf("We only have %v tickets remaingin, so you can't book %v tickets", remainingTickets, userTickets)
				// break
			}
		}

		// Wait task async until complete before program is finish
		wg.Wait()
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}

	/*
		FOR EACH LOOP
		_ -> Variable not used
	*/
	for _, booking := range bookings {
		// 1. MAP WAY
		// firstNames = append(firstNames,  booking["firstName"])
		
		// 2. STRUCT WAY
		firstNames = append(firstNames,  booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("How many tickets do you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string) {
	remainingTickets = remainingTickets - userTickets

	/*
		MAP
		map[key]value
		Concept map same with array associative in php
	*/

	// 1. MAP WAY
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// 2. STRUCT WAY
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		numberOfTickets: userTickets,
	}

	/*
		INSERT DATA TO ARRAY
		bookings[0] = firstName

		INSET DATA TO SLICE
		bookings = append(bookings, firstName + " " + lastName)
	*/
	bookings = append(bookings, userData) 

	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v for booked %v tickets.\n", firstName, userTickets)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string ) {
	// Simulate below task is slow for 10 seconds
	time.Sleep(10 * time.Second)

	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("######################")

	// INFORM TO "wg.Wait()"  THIS CODE IS FINISH
	wg.Done()
}