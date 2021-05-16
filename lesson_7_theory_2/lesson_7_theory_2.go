package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	answer := "Yes"
	for answer == "Yes"{
		a := rand.Intn(9) + 1
		b :=rand.Intn(9)  + 1
		result :=  0
		for a*b != result {
			fmt.Print(a, "*", b,  "=")
			fmt.Scan(&result)
			if a*b != result {
				fmt.Println("Answer is not right, try again")
			}
		}//else
		fmt.Println("Answer is Right!!!")
		fmt.Println("Next Question (Yes/No)?")
		fmt.Scan(&answer)
		for answer != "Yes" && answer != "No" {
			fmt.Println("Next Question (Yes/No)?")
			fmt.Scan(&answer)
		}
	}
}
