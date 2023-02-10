package main

import "fmt"

func main() {
    str := "Hello World"
    for _, c := range str {
        fmt.Printf("'%d' ", c)
    }
    fmt.Println()
}
//In this example, the for loop iterates over each character in the string str. The fmt.Printf function is then used to display each character's integer representation, using the %d format specifier and single quotes to delimit the representation. This will display each character as an integer, represented as '<integer>'.
