Every Go file starts with a package clause. A package is a collection of code that all does similar things, like formatting strings or drawing images. The package clause gives the name of the package that this file’s code will become a part of. In this case, we use the special package main, which is required if this code is going to be run directly (usually from the terminal).

Next, Go files almost always have one or more import statements. Each file needs to import other packages before its code can use the code those other packages contain. Loading all the Go code on your computer at once would result in a big, slow program, so instead you specify only the packages you need by importing them.


The last part of every Go file is the actual code, which is often split up into one or more functions. A function is a group of one or more lines of code that you can call (run) from other places in your program. When a Go program is run, it looks for a function named main and runs that first, which is why we named this function main.
-------------------------------------------------------------------------------
The Println function

Use the Println function when you need to see what your program is doing. Any arguments you pass to it will be printed (displayed) in your terminal, with each argument separated by a space.

After printing all its arguments, Println will skip to a new terminal line. (That’s why “ln” is at the end of its name.)
-------------------------------------------------------------------------------
