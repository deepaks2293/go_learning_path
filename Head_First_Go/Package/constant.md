Many packages export constants : named values that never change.
 A constant declaration looks a lot like a variable declaration, with a name, optional type, and value for the constant. But the rules are slightly different:
 •  Instead of the var keyword, you use the const keyword.
 •  You must assign a value at the time the constant is declared; you can’t assign a value later as with variables.
 •  Variables have the := short variable declaration syntax available, but there is no equivalent for constants.


The value of a variable can vary , but the value of a constant must remain constant . Attempting to assign a new value to a constant will result in a compile error. This is a safety feature: constants should be used for values that shouldn’t ever change.

