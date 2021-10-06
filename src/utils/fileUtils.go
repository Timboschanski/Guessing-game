package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/Timboschanski/Ratespiel/src/config"
)

//SorteFile Read CSV file "scoreboard.csv", takes the Input and sorts it
func SorteFile(filePath string) [][]string {

	records := Readall(filePath)
	sort.Slice(records, func(i, j int) bool {
		return records[i][0] < records[j][0]
	})

	return records

}

func Readall(filePath string) [][]string {

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
func Readfile(filePath string) []string {

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
func Writescore(records [][]string) {

	uniquerecords := [][]string{}
	lastrecord := []string{"xyz", ""}

	i := 0
	for _, v := range records {
		if v[0] != lastrecord[0] || v[1] != lastrecord[1] {
			uniquerecords = append(uniquerecords, v)
			i++
			if i == config.NUMSCORES {
				break
			}
		}
		lastrecord = v
	}

	fmt.Printf("\nCurrent Scoreboard(Top %d):\n", config.NUMSCORES)
	for _, v := range uniquerecords[:] {
		fmt.Println(v)
	}
	f, err := os.OpenFile(config.SCOREBOARDFILE, os.O_WRONLY|os.O_TRUNC, 0644)
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
func Writegames(records2 []string) {

	file, err := os.OpenFile(config.PLAYEDGAMESFILE, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	x := 0
	z, _ := strconv.Atoi(records2[0])
	if x < z {
		x = z + 1
	}
	counter := strconv.Itoa(x)

	_, err = file.WriteString(counter)
	if err != nil {
		fmt.Println("error: can not write into file")
	}
	file.Sync()
}

func AppendScoreboard(attempts int, user string) {
	f, err := os.OpenFile(config.SCOREBOARDFILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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
