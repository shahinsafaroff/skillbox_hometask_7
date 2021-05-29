package main

import "fmt"

func main () {
	var i, count, distance, n1, n2, max int
	count = 0
	distance = 0
	current := [3]int{0, 0, 0}
	for i = 100000; i <= 999999; i++ {
		if i/100000%10+i/10000%10+i/1000%10 == i/100%10+i/10%10+i%10 {
			current[count%3] = i/100000%10 + i/10000%10 + i/1000%10 + i/100%10 + i/10%10 + i%10
			if count > 0 {
				distance = current[count%3] - current[(count-1)%3]
				if distance >= max {
					max = distance
					n1 = current[(count-1)%3]
					n2 = current[count%3]
				}
			} else {
				n1 = 0
				n2 = max
				max = current[0]
			}
			count++
		}
	}
		fmt.Printf("\nCount=%d", count)
		fmt.Printf("\nNumber1=%d, Number2=%d, Distance=%d", n1, n2, max)
	}