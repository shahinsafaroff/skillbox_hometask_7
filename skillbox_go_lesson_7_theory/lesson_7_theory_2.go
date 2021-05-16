package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i :=0;  i<10;  i++ {
		a := rand.Intn(9) + 1
		b :=rand.Intn(9)  + 1
		result :=  0
		for a*b != result {
			fmt.Print(a, "*", b,  "=")
			fmt.Scan(&result)
			if a*b != result {
				fmt.Println("Answer is not right, try again")
			}
		}
		fmt.Println("Answer is Right!!!")
	}
}
