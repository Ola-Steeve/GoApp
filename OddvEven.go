package main

import "fmt"

func OddvEven() {
	for {

		var number int
		var check string
		fmt.Printf("Hello %v, i want you to pick any number, and i would determine if it is even or odd\n", Name)
		fmt.Scan(&number)

		switch {
		case number%2 == 0:
			fmt.Println("This is an even number")
		default:
			fmt.Print("This is an Odd number")
		}

		//Ask if they want another one
		fmt.Println("Do you want to check another number?")
		fmt.Scan(&check)
		if check != "yes" {
			break
		}
	}

}
