Formatting output with Printf and Sprintf


To deal with these sorts of formatting issues, the fmt package provides the Printf function. 

Printf stands for “print, with formatting.” It takes a string and inserts one or more values into it, formatted in specific ways. Then it prints the resulting string.

fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)

The Sprintf function (also part of the fmt package) works just like Printf, except that it returns a formatted string instead of printing it.


resultString:= fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
fmt.Printf(resultString)

Printf and Sprintf can help us limit our displayed values to the correct number of places.

Formatting verbs (the %0.2f in the strings above is a verb)

Value widths (that’s the 0.2 in the middle of the verb)
----------------------------------------------------------------------------------------------------------------------------------------------------
Formatting verbs

first argument to Printf is a string that will be used to format the output. Most of it is formatted exactly as it appears in the string. Any percent signs (%), however, will be treated as the start of a formatting verb

Verb 	Output
%f 	Floating-point number
%d 	Decimal integer
%s 	String
%t 	Boolean (true or false)
%v 	Any value (chooses an appropriate format based on the supplied value’s type)
%#v 	Any value, formatted as it would appear in Go program code
%T 	Type of the supplied value (int, string, etc.)
%% 	A literal percent sign




