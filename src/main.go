package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//main Input of Parameters and run functions
func main() {
	params, err := validation(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	guesser(randomizer(params.From, params.To))
}

type Parameters struct {
	From int
	To   int
}

//validation To Ensure that Errors will Print a User friendly Error and Check that the used Params are appropriate for the Program.
func validation(args []string) (Parameters, error) {

	result := Parameters{}

	if len(args) != 2 {
		err_new := errors.New("error: 2 command line arguments expected")
		return result, err_new
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
		err_new := errors.New("error: Please use integers as command line parameters")
		return result, err_new
	}

	if !(from <= to) {
		err_new := errors.New("error: From must be <= to. Please use valid values")
		return result, err_new
	}
	result.From = from
	result.To = to

	return result, nil
}

//randomizer Randomize a number inbetween given parameters
func randomizer(from int, to int) int {

	rand.Seed(time.Now().UnixNano())
	rdm := (rand.Intn(to-from) + from)
	return rdm
}

//guesser Making the user able to guess the random number and helping to find it
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
