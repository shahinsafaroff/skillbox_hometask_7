package main
import "fmt"
func main() {
	var chessSize, i, j int
	fmt.Scan(&chessSize)
	fmt.Println("Here's Chess table on your front:\n")
	for i=0; i<=chessSize; i++ {
		for j=0; j<=chessSize; j++ {
			if(i%2==1) {
				if (j%2 == 1) {
					fmt.Printf(" ")
				} else {
					fmt.Printf("*")
				}
			} else {
				if (j%2 == 1) {
					fmt.Printf("*")
				}else{
					fmt.Printf(" ")
					}
				}
			}
		fmt.Printf("\n")
	}
	}


