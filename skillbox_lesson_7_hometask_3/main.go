package main
import "fmt"
func main() {
	var height, i, j int
	fmt.Println("Введите высоту  равнобедренного треугольника: ")
	fmt.Scan(&height)

	for i = 0; i < height; i++ {
		for j = 1; j < height - i; j++ {
			fmt.Print(" ")
	}
	for j = height - 2 * i; j <= height; j++{
	fmt.Print("*")
	}
	fmt.Println("\n")
	}
}

