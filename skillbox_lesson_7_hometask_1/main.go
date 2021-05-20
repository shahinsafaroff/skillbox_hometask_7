package main
import (
	"fmt"
)
func main() {
var count, i, lastDigit, fifthDigit, fourthDigit, thirdDigit, secondDigit, firstDigit int
	for i = 100000; i <= 999999; i++ {
	lastDigit= i%10
	fifthDigit = (i/10)%10
	fourthDigit = (i/100)%10
	thirdDigit = (i/1000)%10
	secondDigit = (i/10000)%10
	firstDigit = (i/100000)%10
	if lastDigit == firstDigit && fifthDigit==secondDigit && fourthDigit== thirdDigit{
	count++
	}
	}
	fmt.Println( count)
}
