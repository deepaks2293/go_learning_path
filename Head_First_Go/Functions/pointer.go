package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	myInt = 4
	fmt.Println(reflect.TypeOf(&myInt)) // Print type of &myInt (Get Pointer to myInt)
	var myIntPointer *int               // Decleare a varilabe that hold pointer to int
	myIntPointer = &myInt
	fmt.Println(myIntPointer)  // Print Pointer
	fmt.Println(*myIntPointer) // Print value of Pointer
	var myFloat float64
	myFloat = 98.6
	fmt.Println(reflect.TypeOf(&myFloat))
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	var myBool bool
	myBool = true
	fmt.Println(reflect.TypeOf(&myBool))
	var myBoolPointer *bool
	myBoolPointer = &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
}
