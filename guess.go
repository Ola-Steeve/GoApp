package main

import (
	"fmt"
	"math/rand"
	"time"
)

func guess() {
	// Create a new random number generator

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	//Generate the number btwn 1 - 100 random number generator
	rng := r.Intn(100) + 1
	var random int
	var chances int
	fmt.Println("I would generate a random number and i want you to guess this number in 5 chances")
	//1 second delay
	time.Sleep(1 * time.Second)
	for chances < 5 {

		fmt.Println("Pick any number: ")
		fmt.Scan(&random)
		if random < rng {
			fmt.Println("TOO LOW!!")

		} else if random > rng {
			fmt.Println("TOO HIGH")
		} else {
			fmt.Println("SPOT ON!!")
		}
		chances++
	}

	fmt.Println("Sorry you used all your chances")

}
