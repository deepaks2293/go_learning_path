If you want calls to your function to include arguments, you’ll need to declare one or more parameters. A parameter is a variable, local to a function, whose value is set when the function is called.

func repeatLine (line string, times int){
	for i:= 0 ; i< times; i++ {
		fmt.Println(line)
}
}

You can declare one or more parameters between the parentheses in the function declaration, separated by commas. As with any variable, you’ll need to provide a name followed by a type (float64, bool, etc.) for each parameter you declare.

    A parameter is a variable, local to a function, whose value is set when the function is called.

package main

import "fmt"

func main () {
	repeatLines("hello",3)
}

func repeatLine (line string, times int) {
        for i:= 0 ; i< times; i++ {
                fmt.Println(line)
}
}

hello
hello
hello

