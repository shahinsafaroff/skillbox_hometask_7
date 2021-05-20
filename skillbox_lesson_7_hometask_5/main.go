package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var num int
	target := numRand(1, 100)
	fmt.Println("Find the number within 1 and 100")
	for attempts := 0; attempts < 20; attempts++ {
		fmt.Println(10-attempts, " tries left ")
		fmt.Print("Please enter guessed number: \n")
		_, _ = fmt.Scan(&num)
		fmt.Print(" You've entered number is : ", num)
		if num > target {
			fmt.Println(" Your guess is higher, enter lesser number")
		} else if num < target {
			fmt.Println(" Your guess is lesser, enter higher number")
		} else if num <0 || num > 100{
			fmt.Println(" Your guess is FAR AWAY from SCOPE!!!!")
		} else {
			fmt.Println("BINGO you've found number, try quantity was", target,  attempts, "th try found the number ")
			break
		}

	}

}
func numRand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}