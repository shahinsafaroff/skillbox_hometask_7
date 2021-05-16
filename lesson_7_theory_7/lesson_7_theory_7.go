package main
import "fmt"
func main() {
ForLabel:
	for i := 0; i < 3; i++ {
		fmt.Println("i = ", i)
		for j := 0; j < 3; j++ {
			fmt.Println("j = ", j)
			if i == 1 && j == 1 {
				continue ForLabel
			}
		}
		fmt.Println("Last command in i loop: ")
	}
}
