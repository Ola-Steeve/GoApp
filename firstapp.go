package main

import (
	"fmt"
	"time"
)

var name string

func main() {

	//First part of challenge, a user gives the system an input and the system greets user
	fmt.Println("Good day my name is R.O.Y or Random Operating System Yay, what's yours?")
	fmt.Scan(&name)
	//fmt.Println(&name)

	//current time
	currentTime := time.Now()
	fmt.Printf("The time is %v\n", currentTime)

	//Get just the hour from the whole time
	currentHour := currentTime.Hour()

	/*
		if currentHour >= 1 && currentHour < 12 {
			fmt.Printf("Good Morning %v, i'm here to serve your every need\n", name)

		} else if currentHour >= 12 && currentHour < 16 {
			fmt.Printf("Good Afternoon %v, i'm here to serve your every need\n", name)

		} else {
			fmt.Printf("Good evening %v, i'm here to serve your every need\n", name)

		}
	*/

	//not optimized but a less redundant way of writing the code
	var greeting string
	switch {
	case currentHour >= 1 && currentHour < 12:
		greeting = "Good Morning"
	case currentHour >= 12 && currentHour < 16:
		greeting = "Good Afternoon"
	default:
		greeting = "Good Evening"
	}

	fmt.Printf("%s %v, i'm here to serve your every need\n", greeting, name)

}
