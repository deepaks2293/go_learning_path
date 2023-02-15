package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0
}

func main() {
	var total, amount float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n",amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n",amount)
	total += amount
	fmt.Printf("%0.2f liters needed\n",total)

}

/*

1) removing Printf from painNeeded function and adding return

2) Adding retun type as float64

3) Adding var to store amount, total as float64

4) Add amount in total and printf at the end

*/
