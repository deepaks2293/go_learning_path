package main

import "fmt"

func main() {

	//var int int = 12
	var count int = 12
	//var append string = "minutes of bonus footage"
	var suffix string = "minutes of bonus footage"
	//var fmt string = "DVD"
	var format string = "DVD"
	var languages = append([]string{}, "Espanol")
	fmt.Println(count, suffix, "on", format, languages)

}

/*

Shadowing error

# command-line-arguments
./shadowing_names.go:3:8: "fmt" imported and not used
./shadowing_names.go:10:12: int (variable of type int) is not a type
./shadowing_names.go:11:18: invalid operation: cannot call non-function append (variable of type string)
./shadowing_names.go:12:6: fmt.Println undefined (type string has no field or method Println)
./shadowing_names.go:12:38: undefined: langugaes


Changing variable

*/
