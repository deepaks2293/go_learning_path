Whereas strings are usually used to represent a whole series of text characters, Go’s runes are used to represent single characters.

String literals are written surrounded by double quotation marks ("), but rune literals are written with single quotation marks (').

Go programs can use almost any character from almost any language on earth, because Go uses the Unicode standard for storing runes. Runes are kept as numeric codes, not the characters themselves, and if you pass a rune to fmt.Println, you’ll see that numeric code in the output, not the original character.

Just as with string literals, escape sequences can be used in a rune literal to represent characters that would be hard to include in program code:


package main
import "fmt"

func main() {
	fmt.Println('A')
}

example in hello_world
