The code to prompt for a guess is already in place. We just need to run it more than once. We can use a loop to execute a block of code repeatedly. If you’ve used other programming languages, you’ve probably encountered loops. When you need one or more statements executed over and over, you place them inside a loop.

loops always begin with the for keyword. In one common kind of loop, for is followed by three segments of code that control the loop:

for x:= 4 ; x<=6 ; x++ {
	fmt.Println("x is now", x)
}

x is now 4
x is now 5
x is now 6


An initialization (or init) statement that is usually used to initialize a variable

A condition expression that determines when to break out of the loop

A post statement that runs after each iteration of the loop

Often, the initialization statement is used to initialize a variable, the condition expression keeps the loop running until that variable reaches a certain value, and the post statement is used to update the value of that variable. 

for t:= 3 ; t:= 0 ; t-- {
	fmt.Println(t)
}
fmt.Println("Blast Off")

The ++ and -- statements are frequently used in loop post statements. Each time they’re evaluated, ++ adds 1 to a variable’s value, and -- subtracts 1.

Used in a loop, ++ and -- are convenient for counting up or down.


for t:= 1 ; t <= 3 ; t++ {
        fmt.Println(t)
}

for t:= 3 ; t >= 1 ; t-- {
        fmt.Println(t)
}

Go also includes the assignment operators += and -=. They take the value in a variable, add or subtract another value, and then assign the result back to the variable.


for t:= 1 ; t <= 5 ; t+= 2 {
        fmt.Println(t)
}
1
3
5

for t:= 15 ; t >= 5 ; t-= 5 {
        fmt.Println(t)
}

15
10
5

When the loop finishes, execution will resume with whatever statement follows the loop block. But the loop will keep going as long as the condition expression evaluates to true. It’s possible to abuse this; here are examples of a loop that will run forever, and a loop that will never run at all:

for x:=1 ; true ; x++ {
	fmt.Println(x)
}

for x:=1 ; false ; x++ {
        fmt.Println(x)
}


Init and post statements are optional

If you want, you can leave out the init and post statements from a for loop, leaving only the condition expression (although you still need to make sure the condition eventually evaluates to false, or you could have an infinite loop on your hands).


x:= 1
for x <= 3 {
	fmt.Println(x)
	x++
}

x:= 3
for x >= 1 {
        fmt.Println(x)
        x--
}

