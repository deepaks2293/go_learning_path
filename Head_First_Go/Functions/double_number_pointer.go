package main

import (
	"fmt"
)

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int) {
	*number *= 2 // Output 6
}

/*

1) Write aprogram to double number without return
2) Output 6
3) The output must be 12 as per the logic
4) Here Pointer comes // from Pointer.goconst
5) & Get to the Pointer
6) * value of a Pointer

*/
