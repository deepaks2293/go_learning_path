simple function declaration might look like this:

func sayHi () {
	fmt.Ptinln("Hi!")
}

A declaration begins with the func keyword, followed by the name you want the function to have, a pair of parentheses (), and then a block containing the function’s code.

Once you’ve declared a function, you can call it elsewhere in your package simply by typing its name, followed by a pair of parentheses. When you do, the code in the function’s block will be run.

package main

import "fmt"

func sayHi () {
        fmt.Ptinln("Hi!")
}

func main() {
	sayHi()
}

we call sayHi, we’re not typing the package name and a dot before the function name. When you call a function that’s defined in the current package, you should not specify the package name. (Typing main.sayHi() would result in a compile error.)

The rules for function names are the same as the rules for variable names:

    A name must begin with a letter, followed by any number of additional letters and numbers. (You’ll get a compile error if you break this rule.)

    Functions whose name begins with a capital letter are exported, and can be used outside the current package. If you only need to use a function inside the current package, you should start its name with a lowercase letter.

    Names with multiple words should use camelCase.


    addNumbers - This function name uses camelCase to separate the words "add" and "numbers." It begins with a lowercase letter, which indicates that it is only intended to be used within the current package.

    CalculateSum - This function name also uses camelCase to separate the words "calculate" and "sum." However, it begins with an uppercase letter, which indicates that it is intended to be used outside the current package.

    validateEmail - This function name uses camelCase to separate the words "validate" and "email." It begins with a lowercase letter, indicating that it is only intended to be used within the current package.

    CreateUser - This function name uses camelCase to separate the words "create" and "user." It begins with an uppercase letter, indicating that it is intended to be used outside the current package.

    formatPhoneNumber - This function name uses camelCase to separate the words "format" and "phone" and "number." It begins with a lowercase letter, indicating that it is only intended to be used within the current package.

In each of these examples, the function names begin with a letter, use only letters and numbers, and follow the convention of using camelCase for names with multiple words. Additionally, the use of uppercase or lowercase letters at the beginning of the name indicates whether the function is intended to be used outside or inside the current package.



