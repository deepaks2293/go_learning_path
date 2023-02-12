The time package has a Time type that represents a date (year, month, and day) and time (hour, minute, second, etc.). Each time.Time value has a Year method that returns the year. The code below uses this method to print the current year:

The time.Now function returns a new Time value for the current date and time, which we store in the now variable. Then, we call the Year method on the value that now refers to:

The Year method returns an integer with the year, which we then print.


-------------------------------------------------------------------------------
Methods are functions that are associated with values of a particular type.
-------------------------------------------------------------------------------

The strings package has a Replacer type that can search through a string for a substring, and replace each occurrence of that substring with another string. The code below replaces every # symbol in a string with the letter o:


The strings.NewReplacer function takes arguments with a string to replace ("#"), and a string to replace it with ("o"), and returns a strings.Replacer. When we pass a string to the Replacer value’s Replace method, it returns a string with those replacements made.

Whereas the functions we saw earlier belonged to a package, the methods belong to an individual value. That value is what appears to the left of the dot.



