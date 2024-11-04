package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing Game")
	fmt.Println(
		"A number is chosen randomly. Try to guess it right. The number is an intenger between 0 and 100.",
	)
	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("What is your guess?")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)
		if err != nil {
			fmt.Println("Your guess has to be an integer number")
			return
		}

		switch {
		case guessInt < x:
			fmt.Println("You were wrong. The number drawn is bigger than ", guessInt)
		case guessInt > x:
			fmt.Println("You were wrong. The number drawn is smaller than ", guessInt)
		case guessInt == x:
			fmt.Printf(
				"Congratulations! you got the right number: %d\n"+
					"You got it right in %d tries.\n"+
					"These were your other guesses: %v",
				x, i+1, guesses[:i],
			)

			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"Sadly, you didn't get the right number, the number was: %d\n"+
			"You had 10 tries.\n"+
			"These were your guesses: %v",
		x, guesses,
	)

}
