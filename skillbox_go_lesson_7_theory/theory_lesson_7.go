package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <10; i++ {
		fmt.Println(rand.Intn(100))
	}
}

//Math packet

//Abs(x float64) float64 returns module of X
//Max(x, y float64) float64 returns the biggest out of x and y
//Min(x, y float64) float64  returns the smallest out of x and y
//Pow(x, y float64) float64
//Round(x float64) float64
//Sqrt(x float64) float64
//Trunc(x float64) float64 returns whole part of number
