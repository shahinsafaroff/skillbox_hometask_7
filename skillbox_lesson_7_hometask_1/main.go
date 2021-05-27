package main
import (
	"fmt"
)
func tickets() int {
var i, sum int
sum = 0
for i = 100000; i <= 999999; i++ {
		if (i/100000%10 == i/100%10 && i/10000%10==i/10%10 && i/1000%10 == i%10 ) {
			sum++
		}
	}
return sum
}
func main() {
	var b int
	b = tickets()
	fmt.Println("Количество счастливых билетов: ", b)
}
