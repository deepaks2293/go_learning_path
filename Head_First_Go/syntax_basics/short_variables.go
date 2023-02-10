package main

import (
	"fmt"
	//"reflect" // if not used than there will be error
)

func main() {

	//var quantity int
	//var length, width float64
	//var customerName string

	quantity:= 4
	length, width:= 1.2, 2.4
	customerName:= "LoneRanger"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "sqaure meters")

}
