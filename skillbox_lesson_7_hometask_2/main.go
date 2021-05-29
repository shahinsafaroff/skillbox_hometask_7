package main
import "fmt"
func main() {

	var width, height, row, col int
	var current, temp string
	fmt.Println("Please enter height and width of chess table: \n")
	fmt.Scan(&width, &height)

	for row = 0; row <height; row++ {
	current, temp = "*", " "
			if row%2 != 0 {
				current, temp = temp, current
			}
			for col = 0; col < width; col++{
				fmt.Print(current)
				current, temp = temp, current
			}
			fmt.Println()
	}
	fmt.Println("end")
	}


