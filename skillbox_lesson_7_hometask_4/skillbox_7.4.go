package main

import "fmt"

func main () {
	var lastDigit, fifthDigit, fourthDigit, thirdDigit, secondDigit, firstDigit, count, distance, n1, n2, max int
	count = 0
	distance = 0
	current:=[3] int {0, 0, 0}
	for lastDigit = 0; lastDigit < 10; lastDigit++ {
		for fifthDigit = 0; fifthDigit < 10; fifthDigit++ {
			for fourthDigit = 0; fourthDigit < 10; fourthDigit++ {
				for thirdDigit = 0; thirdDigit < 10; thirdDigit++ {
					for secondDigit = 0; secondDigit < 10; secondDigit++ {
						for firstDigit = 0; firstDigit < 10; firstDigit++ {
							if (lastDigit+fifthDigit+fourthDigit == thirdDigit+secondDigit+firstDigit) {
								current[count%3] = lastDigit*100000 + fifthDigit*10000 + fourthDigit*1000 + thirdDigit*100 + secondDigit*10 + firstDigit
								if (count>0) {
									distance = current[count%3] - current[(count-1)%3]
									if (distance >= max) {
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
					}
				}
			}
		}
	}
	fmt.Printf("\nCount=%d", count)
	fmt.Printf("\nNumber1=%d, Number2=%d, Distance=%d", n1, n2, max)
}
