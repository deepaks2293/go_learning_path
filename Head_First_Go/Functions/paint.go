package main

import "fmt"

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of  %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of  %0.2f is invalid", height)
	}
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / 10.0, nil
}

func main() {
	var total, amount float64
	var err error
	amount, err = paintNeeded(4.2, 3.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
		total += amount
		fmt.Printf("Total: %0.2f liters needed\n", total)
	}

}

/*

1) removing Printf from painNeeded function and adding return

2) Adding retun type as float64

3) Adding var to store amount, total as float64

4) Add amount in total and printf at the end

5) Added multiple return funtion (float64 and error)

6) In main funtion added if else loop to catch error.const


*/
