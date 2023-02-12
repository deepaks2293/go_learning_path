package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
	var broken string = "G# r#cks!"
	var replacer *strings.Replacer = strings.NewReplacer("#", "o")
	var fixed string = replacer.Replace(broken)
	fmt.Println(fixed)
}

/*
In the statement var replacer strings.Replacer = strings.NewReplacer("#", "o") you are defining a variable with the type strings.Replacer and then trying to assign a value of type *strings.Replacer to it. As these are two different types the compiler reports an error. The fix is to use the correct type var replacer *strings.Replacer = strings.NewReplacer("#", "o") (playground).

replacer := strings.NewReplacer("#", "o") works OK because when using a short variable declaration the compiler determines the type (*strings.Replacer) for you.
*/
