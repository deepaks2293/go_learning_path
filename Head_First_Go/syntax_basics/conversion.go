//Math and comparison operations in Go require that the included values be of the same type. If they’re not, you’ll get an error when trying to run your code.

package main

import (
	"fmt"
)

func main() {
	var price int = 100
	fmt.Println("Price is", price, "rupee.")

	var taxrate float64 = 0.08
	var tax float64 = float64(price) * float64(taxrate)
	fmt.Println("Tax is", tax, "rupee.")

	var total float64 = float64(price) + float64(tax)
	fmt.Println("Total cost is", total, "rupee.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "rupee avilable.")
	fmt.Println("Within Budget?", float64(total) <= float64(availableFunds))
}
