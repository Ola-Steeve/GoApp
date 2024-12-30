package main

import (
	"fmt"
	"math"
)

func prime() {
	// var count int
	// primeStore := []int{}
	var primeNumber int

	fmt.Println("Give me any number and i would tel you if it is a prime number or not")
	fmt.Scan(&primeNumber)

	switch {
	case primeNumber == 1:
		fmt.Println("Not a prime, 1 cannot be a prime")
	default:
		for i := 2; i < int(math.Sqrt(float64(primeNumber))); i++ {
			if primeNumber%i == 0 {
				fmt.Println("Not a prime")
			} else {
				fmt.Println("This is a prime")
			}
		}
	}
}
