package main

import (
	"fmt"
	"strings"
	"sync"
)

const conferenceTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var wg = sync.WaitGroup{}

func main() {
	for remainingTickets > 0 {
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		} else {
			newBooking()
		}
	}
}

func newBooking() {

	greetUsers(conferenceName, conferenceTickets, remainingTickets)
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	userInfo := userData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}
	if isValidName && isValidEmail && isValidTicketNumber {
		userInfo.bookTicket()
		wg.Add(1)
		go userInfo.sendTicket()
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}

	wg.Wait()
}

func (u *userData) bookTicket() {
	remainingTickets = remainingTickets - (*u).tickets
	fmt.Println(remainingTickets)
	bookings = append(bookings, *u)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", (*u).firstName, (*u).lastName, (*u).tickets, (*u).email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func (u *userData) sendTicket() {
	// time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", (*u).tickets, (*u).firstName, (*u).lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, (*u).email)
	fmt.Println("#################")
	wg.Done()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
