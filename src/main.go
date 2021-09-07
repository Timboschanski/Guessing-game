package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
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

	guesser(randomizer(params.From, params.To), params.User)

	records := readCsvFile("scoreboard.csv")
	for _, v := range records[:5] {
		fmt.Println(v)
	}
	fmt.Println("Gespielte Spiele:", len(records))
}

type Parameters struct {
	From int
	To   int
	User string
}

//validation To Ensure that Errors will Print a User friendly Error and Check that the used Params are appropriate for the Program.
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

//randomizer Randomize a number inbetween given parameters
func randomizer(from int, to int) int {

	rand.Seed(time.Now().UnixNano())
	rdm := (rand.Intn(to-from) + from)
	return rdm
}

//guesser Making the user able to guess the random number and helping to find it
func guesser(rdm int, user string) {

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

	f, err := os.OpenFile("scoreboard.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	output := "\n" + strconv.Itoa(attempts) + "," + user

	a, err := f.WriteString(output)
	if err != nil {
		fmt.Println("error: can not write into file")
	}
	fmt.Printf("%d bytes written\n", a)
	f.Sync()
}

//readCsvFile read csv file, take the Input and sorts it
func readCsvFile(filePath string) [][]string {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})

	fmt.Println("Current Scoreboard(Top 5):")
	return records
}
