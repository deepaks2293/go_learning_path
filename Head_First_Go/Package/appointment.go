package main

import (
	"dates"
	"fmt"
)

func main() {
	days: = 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up", days + dates.DaysInWeek, "days")
}
