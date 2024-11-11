package main

import "fmt"

func main() {
	conferenceName := "GopherCon" // variable must be used, imported packages must be used

	const conferenceTickets = 50
	var remainingTickets uint = 50 // uint is an unsigned integer, it can't be negative

	fmt.Printf("Welcome to %v, booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// go is a statically typed language, so tell the type of variable to the compiler when declaring
	// var userName
	var userName string
	var userTicket uint
	fmt.Println("Enter your name to book tickets")
	fmt.Scan(&userName)

	fmt.Println("How many tickets you want to book?")
	fmt.Scan(&userTicket)

	remainingTickets -= userTicket

	fmt.Printf("User %v has booked %v tickets\n", userName, userTicket)
	fmt.Printf("%v remaining tickets are available\n", remainingTickets)
}
