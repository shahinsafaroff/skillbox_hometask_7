package main
import (
	"fmt"
)

func main() {
const min = 100_000
const max = 999_000
var counter int
for ticket :=min; ticket<= max; ticket++ {
	var rev int
	number :=ticket
	for number > 999 {
		digit := number % 10
		rev = rev*10 + digit
		number /=10
	}
	high := ticket / 1_000
	if high == rev {
		counter++
	}
}
fmt.Println(counter)
}
