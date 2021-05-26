package main
import (
	"fmt"
)
func main() {
var count, i, lastDigit, fifthDigit, fourthDigit, thirdDigit, secondDigit, firstDigit int
	for i = 100000; i <= 999999; i++ {
		for firstDigit = 0; firstDigit < 10; firstDigit++ {
			for secondDigit = 0; secondDigit < 10; secondDigit++ {
				for thirdDigit = 0; thirdDigit < 10; thirdDigit++ {
					for fourthDigit = 0; firstDigit < 10; fourthDigit++ {
						for fifthDigit = 0; fifthDigit < 10; fifthDigit++ {
							for lastDigit = 0; lastDigit < 10; lastDigit++ {
								if ((firstDigit + secondDigit + thirdDigit) == (fourthDigit + fifthDigit + lastDigit)) {
									count++
									fmt.Println(count)
								}
							}
						}
					}
				}
			}
		}
	}
	//lastDigit= i%10
	//fifthDigit = (i/10)%10
	//fourthDigit = (i/100)%10
	//thirdDigit = (i/1000)%10
	//secondDigit = (i/10000)%10
	//firstDigit = (i/100000)%10
	//if lastDigit == firstDigit && fifthDigit==secondDigit && fourthDigit== thirdDigit{
	//count++
	//}
	//}
	//fmt.Println( count)
}
