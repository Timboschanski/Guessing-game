package app

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Timboschanski/Ratespiel/src/config"
	"github.com/Timboschanski/Ratespiel/src/utils"
)

type Parameters struct {
	From int
	To   int
	User string
}

func StartGame(args *[]string) {

	params, err := validation(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {

		guesser(utils.Randomizer(params.From, params.To), params.User)

		records := utils.SorteFile(config.SCOREBOARDFILE)
		utils.Writescore(records)

		records2 := utils.Readfile(config.PLAYEDGAMESFILE)
		fmt.Println("Gespielte Spiele:", records2)
		utils.Writegames(records2)
	}
}

//guesser Making the user able to guess the random number and helping to find it
func guesser(rdm int, user string) {
	fmt.Println("------------------------------------------")
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

	utils.AppendScoreboard(attempts, user)
}

//Validation To ensure that errors will print a user friendly error massage and check that the used params are appropriate for the program.
func validation(args []string) (Parameters, error) {

	result := Parameters{}

	if len(args) != 3 {
		err_new := errors.New("error: 3 command line arguments expected(From To User)")
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

	user := args[2]

	if parseFailure {
		err_new := errors.New("error: Please use integers for the first 2 command line parameters and a Username as Third")
		return result, err_new
	}

	if !(from <= to) {
		err_new := errors.New("error: From must be <= to. Please use valid values")
		return result, err_new
	}
	result.From = from
	result.To = to
	result.User = user

	return result, nil
}
