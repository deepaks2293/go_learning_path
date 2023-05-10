package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func main() {
	fmt.Print("Enter a temprature in Fahrenheit: ")
	farhrenheit, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	celcius := (farhrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees celcisus\n", celcius)
}

/*

Same Function as of pass_fail.go to convert F to C

*/
