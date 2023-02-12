It’s true that when the same variable name is declared twice in the same scope, we get a compile error:

a := 1
a := 2

But as long as at least one variable name in a short variable declaration is new, it’s allowed. The new variable names are treated as a declaration, and the existing names are treated as an assignment.

a := 1
b,a := 2,3
a,c := 4,5
fmt.Println(a,b,c) // 4 2 5

There’s a reason for this special handling: a lot of Go functions return multiple values. It would be a pain if you had to declare all the variables separately just because you want to reuse one of them.

Instead, Go lets you use a short variable declaration for everything, even if for one of the variables, it’s actually an assignment.
