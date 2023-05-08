package main

import "fmt"

func main() {

	var width, height, area float64

	width = 4.2
	height = 3.0
	area = width * height

	fmt.Println(area/10.0, "liters needed")
	fmt.Printf("%f liters needed\n", area/10.0)

	width = 5.2
	height = 3.5
	area = width * height

	fmt.Println(area/10.0, "liters needed")
	fmt.Printf("%f liters needed\n", area/10.0)

	fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	fmt.Println("------------------------------------")
	fmt.Printf("%12s | %2d\n", "Stamps", 50)
	fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	fmt.Printf("%12s | %2d\n", "Tape", 99)

}

/*

1) repetitive code, we need to calculate the amount of paint needed to cover several walls. The manufacturer says each liter of paint covers 10 square meters. So, we’ll need to multiply each wall’s width (in meters) by its height to get its area, and then divide that by 10 to get the number of liters of paint needed.

2) calculations are slightly off because ordinary floating-point arithmetic on computers is ever-so-slightly inaccurate. (Usually by a few quadrillionths.) The reasons are a little too complicated to get into here, but this problem isn’t exclusive to Go.

3) Formatting output with Printf and Sprintf

4) minimum width after the percent sign for a formatting verb. If the argument matching that verb is shorter than the minimum width, it will be padded with spaces until the minimum width is reached.

5) f := 3.14159265359
fmt.Printf("Floating-point number with precision: %2.1f\n", f)

The % character is used to indicate that a formatted string follows. The 2 between the % and the .1 indicates the minimum width of the printed value, which is 2 characters in this case. If the printed value is less than 2 characters, it will be padded with spaces on the left to reach the minimum width. The .1 after the 2 indicates that the floating-point number should be displayed with a precision of 1 decimal place. The f character indicates that the value to be formatted is a floating-point number.

When this code is executed, the following output will be displayed:

Floating-point number with precision:  3.1


Note the space before the number 3.1. This is because the minimum width is set to 2 characters, but the number itself only takes up 3 characters (the digits 3, . and 1). To fill up the remaining character, a space is added before the number.

If we had used %1.1f instead, the output would have been 3.1 without any leading spaces.

6) Add .2 in paint

*/
