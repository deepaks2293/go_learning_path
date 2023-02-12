ming a variable error would be a bad idea, because it would shadow the name of a type called error.

When you declare a variable, you should make sure it doesn’t have the same name as any existing functions, packages, types, or other variables. If something by the same name exists in the enclosing scope (we’ll talk about scopes shortly), your variable will shadow it—that is, take precedence over it. And all too often, that’s a bad thing.

Here, we declare a variable named int that shadows a type name, a variable named append that shadows a built-in function name (we’ll see the append function in Chapter 6), and a variable named fmt that shadows an imported package name. Those names are awkward, but they don’t cause any errors by themselves...



