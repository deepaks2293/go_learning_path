Conditionals

If our program encounters a problem reading input from the keyboard, we’ve set it up to report the error and stop running. But now, it stops running even when everything’s working correctly!
image

Functions and methods like ReadString return an error value of nil, which basically means “there’s nothing there.” In other words, if err is nil, it means there was no error. But our program is set up to simply report the nil error! What we should do is exit the program only if the err variable has a value other than nil.

We can do this using conditionals: statements that cause a block of code (one or more statements surrounded by {} curly braces) to be executed only if a condition is met.

if 1 < 2 {
   fmt.Println("It's true!")
}

An expression is evaluated, and if its result is true, the code in the conditional block body is executed. If it’s false, the conditional block is skipped.

if true {
    fmt.Println("I'll be Printed")
}

if false {
   fmt.Println("I won't!")
}

As with most other languages, Go supports multiple branches in the conditional. These statements take the form if...else if...else.

if grade == 100 {
       fmt.Println("Perfect!")
} else if grade >= 60 {
       fmt.Println("You pass.")
} else {
       fmt.Println("You fail!")
}

Conditionals rely on a Boolean expression (one that evaluates to true or false) to decide whether the code they contain should be executed.

if 1 == 1{
   fmt.Println("I'll be Printed")
}

if 1> 2 {
   fmt.Println("I won't")
}

When you need to execute code only if a condition is false, you can use !, the Boolean negation operator, which lets you take a true value and make it false, or a false value and make it true.

if !true{
    fmt.Println("I won't")
}

if !false{
    fmt.Println("I will")
}
If you want to run some code only if two conditions are both true, you can use the && (“and”) operator. If you want it to run if either of two conditions is true, you can use the || (“or”) operator.


if true && false {
	fmt.Println(" I won't!")
}

if false || true {
 	fmt.Println("I'll be printed!")
}
