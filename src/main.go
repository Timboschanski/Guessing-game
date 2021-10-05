package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	SCOREBOARDFILE  = "./utils/scoreboard.csv"
	PLAYEDGAMESFILE = "./utils/playedgames.csv"
)

//main Input of parameters and run functions
func main() {
	go server()
	for {
		params, err := validation(os.Args[1:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		guesser(randomizer(params.From, params.To), params.User)

		records := sorter(SCOREBOARDFILE)
		writescore(records)

		records2 := readfile(PLAYEDGAMESFILE)
		fmt.Println("Gespielte Spiele:", records2)
		writegames(records2)

	}
}

type Parameters struct {
	From int
	To   int
	User string
}

//validation To ensure that errors will print a user friendly error massage and check that the used params are appropriate for the program.
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

	f, err := os.OpenFile(SCOREBOARDFILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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

//sorter Read CSV file "scoreboard.csv", takes the Input and sorts it
func sorter(filePath string) [][]string {

	records := readall(filePath)
	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})

	return records

}

func readall(filePath string) [][]string {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, _ := r.ReadAll()
	return records
}

//readfile Read CSV file "playedgames.csv"
func readfile(filePath string) []string {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, _ := r.Read()
	return records
}

//writescore Read CSV file "scoreboard.csv" and overwrites it
func writescore(records [][]string) {

	uniquerecords := [][]string{}
	lastrecord := []string{"xyz", ""}

	i := 0
	for _, v := range records {
		if v[0] != lastrecord[0] || v[1] != lastrecord[1] {
			uniquerecords = append(uniquerecords, v)
			i++
			if i == 5 {
				break
			}
		}
		lastrecord = v
	}

	fmt.Println("Current Scoreboard(Top 5):")
	for _, v := range uniquerecords[:] {
		fmt.Println(v)
	}
	f, err := os.OpenFile(SCOREBOARDFILE, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	err = w.WriteAll(uniquerecords)

	if err != nil {
		fmt.Println("error: can not write into file")
	}
	f.Sync()
}

//writegames Reads CSV file "playedgames.csv" and Overwrites it
func writegames(records2 []string) {

	file, err := os.OpenFile(PLAYEDGAMESFILE, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	x := 0
	z, _ := strconv.Atoi(records2[0])
	if x < z {
		x = z + 1
	}
	counter := strconv.Itoa(x)

	b, err := file.WriteString(counter)
	if err != nil {
		fmt.Println("error: can not write into file")
	}
	fmt.Printf("%d bytes written\n", b)
	file.Sync()
}

//server Creats a Web Server
func server() {

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods(http.MethodGet)

	r.HandleFunc("/scoreboard/", scoreall).Methods(http.MethodGet)
	r.HandleFunc("/scoreboard/{player}", score).Methods(http.MethodGet)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hi, here you can get the csv files from the guessing-game."))

}

func score(w http.ResponseWriter, req *http.Request) {

	score := readall(SCOREBOARDFILE)

	lenr := len(score)
	if lenr > 5 {
		lenr = 5
	}
	score = score[:lenr]

	vars := mux.Vars(req)
	post_id := vars["player"]

	var s string
	for _, v := range score {
		if v[1] == post_id {
			s += fmt.Sprintf("%s %s\n", v[0], v[1])
		}
	}

	w.Write([]byte("Scoreboard\n"))
	w.Write([]byte(s))
}

func scoreall(w http.ResponseWriter, req *http.Request) {

	score := readall(SCOREBOARDFILE)

	lenr := len(score)
	if lenr > 5 {
		lenr = 5
	}
	score = score[:lenr]

	var s string
	for _, v := range score {
		s += fmt.Sprintf("%s %s\n", v[0], v[1])
	}

	w.Write([]byte("Scoreboard\n"))
	w.Write([]byte(s))
}
