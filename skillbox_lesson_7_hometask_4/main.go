package main

import (
	"fmt"
	"math"
)

func main () {
	var maxScopeTickets, happyTicketNumber, ticketNumber int
	happyTicketNumber = 100000

	for ticketNumber = 100000; ticketNumber <=999999; ticketNumber++{
		var sumOfFirstHalf, sumOfSecondHalf int

		for n := 0; n < 3; n++ {
			sumOfFirstHalf += ticketNumber / int(math.Pow10(n)) % 10
			sumOfSecondHalf += ticketNumber / int(math.Pow10(n+3)) % 10
		}
		if sumOfFirstHalf == sumOfSecondHalf {
			scope := ticketNumber - happyTicketNumber
			if scope > maxScopeTickets {
				maxScopeTickets = scope
			}
			happyTicketNumber = ticketNumber
		}
	}
	fmt.Println("Tickets are needed to buy: ", maxScopeTickets)
}