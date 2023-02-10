package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.Floor(2.75))// donot take strings
	fmt.Println(strings.Title("i will become netops")) // Donot take integers
	//fmt.Println(strings.Title(2.75))
	//fmt.Println(math.Floor("i will become netops"))
}

//command-line-arguments 
//./import_multiple_package.go:12:28: cannot use 2.75 (untyped float constant) as string value in argument to strings.Title
//./import_multiple_package.go:13:25: cannot use "i will become netops" (untyped string constant) as float64 value in argument to math.Floor
