//guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've choosen random number between 1 to 100.")
	fmt.Println("Can you guess it?")
	//fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	sucess := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guess left.")
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			sucess = true
			fmt.Println("Good Job!, You have guess right!")
			break
		}
	}
	if !sucess {
		fmt.Println("Sorry, you din't guess my number. it was", target)
	}

}

/*

1) math/rand package has a Intn function that can generate a random number for us, so we’ll need to import math/rand. Then we’ll call rand.Intn to generate the random number.
One is the package’s import path, and the other is the package’s name.

2) we’ll get a random number in the range 0–99. Since we need a number in the range 1–100, we’ll just add 1 to whatever random value we get. We’ll store the result in a variable, target.

3) The rand.Seed function expects an integer, so we can’t pass it a Time value directly. Instead, we call the Unix method on the Time, which will convert it to an integer.

4) Getting an integer from the keyboard

5) That should work in much the same way as when we read in a percentage grade from the keyboard for our grading program.

6) There will be only one difference: instead of converting the input to a float64, we need to convert it to an int (since our guessing game uses only whole numbers). So we’ll pass the string read from the keyboard to the strconv package’s Atoi (string to integer) function instead of its ParseFloat function. Atoi will give us an integer as its return value. (Just like ParseFloat, Atoi might also give us an error if it can’t convert the string. If that happens, we again report the error and exit.)

7) Comparing the guess to the target

8) If guess is less than target, we need to print a message saying the guess was low. Otherwise, if guess is greater than target, we should print a message saying the guess was high. Sounds like we need an if...else if statement. We’ll add it below the other code in our main function.

9) We’ll use an int variable named guesses to track the number of guesses the player has made. In our loop’s init statement, we’ll initialize guesses to 0. We’ll add 1 to guesses with each iteration of the loop, and we’ll stop the loop when guesses reaches 10.We’ll also add a Println statement at the top of the loop’s block to tell the user how many guesses they have left.

10) Skipping parts of a loop with “continue” and “break”

11) Revealing the target

12) If the player makes 10 guesses without finding the target number, the loop will exit. In that event, we need to print a message saying they lost, and tell them what the target was.

13) So, before our guessing loop, we’ll declare a success variable that holds a bool. (We need to declare it before the loop so that it’s still in scope after the loop ends.) We’ll initialize success to a default value of false. Then, if the player guesses correctly, we’ll set success to true, indicating we don’t need to print the failure message.
*/
