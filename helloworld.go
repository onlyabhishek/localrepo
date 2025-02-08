package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets=50
	var remainingTicket uint=50
	//remainingTicket=-1
	fmt.Println("conferenceTickets is %T,remainingTickets is %T,conferenceName is %T\n",conferenceTickets,remainingTicket,conferenceName)
    fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still availale\n",conferenceTickets,remainingTicket)
    fmt.Println("Get your ticket here to attend")
    // fmt.Println(conferenceName)
	//conferenceTickets=30
	// print(conferenceTickets)
	var userName string
	var userTickets int
	fmt.Scan(&userName)
	userTickets=2
	fmt.Println("User %v booked %v tickets.\n",userName,userTickets)
} 
