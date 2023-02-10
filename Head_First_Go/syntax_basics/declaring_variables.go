package main

import (
	"fmt"
	"reflect"
)

func main() {

	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "LoneRanger"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "sqaure meters")

	// declaring variable before hand same variable cant be redeclared
	var quantity_new = 5
	var length_new,width_new = 1.2, 2.4
	var customerName_new = "Deeapk"
	
	fmt.Println(reflect.TypeOf(quantity_new))
	fmt.Println(reflect.TypeOf(length_new))
	fmt.Println(reflect.TypeOf(width_new))
	fmt.Println(reflect.TypeOf(customerName_new))


	// Zero values If you declare a variable without assigning it a value, that variable will contain the zero value
	var myInt int
	var myFloat float64
	fmt.Println(myInt,myFloat)
}
