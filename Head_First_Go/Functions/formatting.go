package main

import "fmt"

func main() {

	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)

	resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
	fmt.Printf(resultString)
	fmt.Println(resultString)

	// Floating-point number
	f := 3.14
	fmt.Printf("Floating-point number: %f\n", f) // Output: Floating-point number: 3.140000

	// Decimal integer
	d := 42
	fmt.Printf("Decimal integer: %d\n", d) // Output: Decimal integer: 42

	// String
	s := "Hello, world!"
	fmt.Printf("String: %s\n", s) // Output: String: Hello, world!

	// Boolean (true or false)
	t := true
	fmt.Printf("Boolean: %t\n", t) // Output: Boolean: true

	// Any value (chooses an appropriate format based on the supplied value’s type)
	v := "abc"
	fmt.Printf("Any value: %v\n", v) // Output: Any value: abc

	// Any value, formatted as it would appear in Go program code
	m := map[string]int{"foo": 1, "bar": 2}
	fmt.Printf("Formatted value: %#v\n", m) // Output: Formatted value: map[string]int{"bar":2, "foo":1}

	// Type of the supplied value (int, string, etc.)
	var x interface{} = 3.14
	fmt.Printf("Type of x: %T\n", x) // Output: Type of x: float64

	// A literal percent sign
	fmt.Printf("A literal percent sign: %%\n") // Output: A literal percent sign: %

}
