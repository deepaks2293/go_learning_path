Functions always return values of a specific type (and only that type). To declare that a function returns a value, add the type of that return value following the parameters in the function declaration. Then use the return keyword in the function block, followed by the value you want to return.

func double(number float64) float64 {
	return number * 2
}

func main() {
	dozen:= double (6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}

When a return statement runs, the function exits immediately, without running any code that follows it. You can use this together with an if statement to exit the function in conditions where there’s no point in running the remaining code (due to an error or some other condition).

func double(number float64) float64 {
	return number * 2
	fmt.Println(number *2 )
}

Error : missing return at the end of function

func double(number float64) float64 {
        return int(number * 2)
}

ERROR: As return type is float can return int


