package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a temprature in Fahrenheit: ")
	farhrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (farhrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees celcisus\n", celcius)
}

/*
importing keyboard funtion
Same Function as of pass_fail.go to convert F to C

*/
