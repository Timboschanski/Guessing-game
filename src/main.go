package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	params := validation()
	param := randomizer(params.From, params.To)
	guesser(param.Rdm)
}

type Parameters struct {
	From int
	To   int
}

//validation To Ensure that Errors will Print a User friendly Error and Check that the used Params are appropriate for the Program.
func validation() Parameters {

	result := Parameters{}

	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("error: 2 command line arguments expected")
		os.Exit(1)
	}

	parseFailure := false

	from, err := strconv.Atoi(args[0])
	if err != nil {
		parseFailure = true
	}

	to, err := strconv.Atoi(args[1])
	if err != nil {
		parseFailure = true
	}

	if parseFailure {
		fmt.Println("error: Please use integers as command line parameters.")
		os.Exit(1)
	}

	if !(from <= to) {
		fmt.Println("error: From must be <= to. Please use valid values.")
		os.Exit(1)
	}
	result.From = from
	result.To = to

	return result
}

type Parameter struct {
	Rdm int
}

func randomizer(from int, to int) Parameter {

	random := Parameter{}

	rand.Seed(time.Now().UnixNano())
	rdm := (rand.Intn(to-from) + from)

	random.Rdm = rdm
	return random

}

func guesser(rdm int) {

	fmt.Println("Guess a number that is between your params")
	fmt.Println("Please input your guess")

	attempts := 0
	for {
		attempts++
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error: An error occured while reading input. Please try again", err)
			continue
		}

		input = strings.TrimSuffix(input, "\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("error: Invalid input. Please use a number")
			continue
		}

		fmt.Println("Your guess is", guess)

		if guess > rdm {
			fmt.Println("Your guess is too high. Try again")
		} else if guess < rdm {
			fmt.Println("Your guess is too low. Try again")
		} else {
			fmt.Println("Woohoo thats correct. Good Job! That took you", attempts, "attempts")
			break
		}
	}
}
