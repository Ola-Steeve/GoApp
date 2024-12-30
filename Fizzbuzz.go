package main

import (
	"fmt"
)

func fizzbuzz() {

	fizz := "fizz"
	count := []string{}
	buzz := "buzz"
	// duckula := []int
	// var duckula []interface{} for slices with different types stored
	fmt.Println("I would print  numbers from 1 - 100 and replace multiples of 3 with fizz and multiples od 5 wiht buzz")

	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			count = append(count, fizz)

		} else if i%5 == 0 {
			count = append(count, buzz)

		} else {
			count = append(count, fmt.Sprintf("%d", i)) //Sprintf is a string formatter that turns an interger to a string

		}

	}
	fmt.Println(count)
}
