package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
func main() {
	fmt.Print("Enter a grade:  ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of", grade, "is", status)
}

/*

1) the Println function, Print doesn’t skip to a new terminal line after printing a message, which lets us keep the prompt and the user’s entry on the same line.

2) way to read (receive and store) input from the program’s standard input, which all keyboard input goes to. The line reader := bufio.NewReader(os.Stdin) stores a bufio.Reader in the reader variable that can do that for us.

3) actually get the user’s input, we call the ReadString method on the Reader. The ReadString method requires an argument with a rune (character) that marks the end of the input. We want to read everything the user types up until they press Enter, so we give ReadString a newline rune.

4) Println to print
Error: ./pass_fail.go:12:11: assignment mismatch: 1 variable but reader.ReadString returns 2 values

5) most programming languages, functions and methods can only have a single return value, but in Go, they can return any number of values. The most common use of multiple return values in Go is to return an additional error value that can be consulted to find out if anything went wrong while the function or method was running.

6) Go doesn’t allow us to declare a variable unless we use it

7) Go requires that every variable that gets declared must also get used somewhere in your program. If we add an err variable and then don’t check it, our code won’t compile.

8) Add err input,err

Error: ./pass_fail.go:12:8: err declared and not used
9) Option 1 to ignore err using _ identifier Assigning a value to the blank identifier essentially discards it

10) it works but if err comes how to handle it

11) The log package has a Fatal function that can do both of these operations for us at once: log a message to the terminal and stop the program. (“Fatal” in this context means reporting an error that “kills” your program.)

12) if our program encounters a problem reading input from the keyboard, we’ve set it up to report the error and stop running. But now, it stops running even when everything’s working correctly!

13) Functions and methods like ReadString return an error value of nil, which basically means “there’s nothing there.”

14) f the value in our err variable is nil, it means reading from the keyboard was successful. Now that we know about if statements, let’s try updating our code to log an error and exit only if err is not nil.

15) Adding if ele loop

ERROR ./pass_fail.go:18:14: invalid operation: input >= 60 (mismatched types string and untyped int)

16) pair of issues to address here:
	The input string still has a newline character on the end, from when the user pressed the Enter key while entering it. We need to strip that off.
	The remainder of the string needs to be converted to a floating-point number.

17) Removing the newline character from the end of the input string will be easy. The strings package has a TrimSpace function that will remove all whitespace characters (newlines, tabs, and regular spaces) from the start and end of a string.

18) All that should remain in the input string now is the number the user entered. We can use the strconv package’s ParseFloat function to convert it to a float64 value. ParseFloat converts the string to a number, and returns it as a float64 value. Like ReadString, it also has a second return value, an error, which will be nil unless there was some problem converting the string. (For example, a string that can’t be converted to a number. We don’t know of a numeric equivalent to "hello"...)

19) It seems a little strange that we’re getting the error twice, but let’s disregard that for now. We’ll add a call to Println to print the percentage grade we were given, and the value of status.

ERROR: ./pass_fail.go:26:3: status declared and not used
./pass_fail.go:28:3: status declared and not used
./pass_fail.go:30:41: undefined: status


20) solution is to move the declaration of the status variable out of the conditional blocks, and up to the function block. Once we do that, the status variable will be in scope both within the nested conditional blocks and at the end of the function block.



*/
