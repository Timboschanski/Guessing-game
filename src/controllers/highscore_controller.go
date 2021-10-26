package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/Timboschanski/Ratespiel/src/config"
	"github.com/Timboschanski/Ratespiel/src/utils"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hi, here you can get the csv files from the guessing-game."))

}

func Score(w http.ResponseWriter, req *http.Request) {

	score := utils.CutTo(utils.Readall(config.SCOREBOARDFILE), config.NUMSCORES)

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

func Scoreall(w http.ResponseWriter, req *http.Request) {

	out := req.URL.Query().Get("out")
	score := utils.CutTo(utils.Readall(config.SCOREBOARDFILE), config.NUMSCORES)

	var s string
	for _, v := range score {
		s += fmt.Sprintf("%s %s\n", v[0], v[1])
	}

	if out == "csv" {

		w.Write([]byte("HIGHSCORE\n"))
		w.Write([]byte(s))
	}

	type scorestruc struct {
		Attempts string `json:"attempts" yaml:"attempts"`
		User     string `json:"user" yaml:"user"`
	}

	type highscorestruc struct {
		Title  string       `json:"title"`
		Scores []scorestruc `json:"scores"`
	}

	if out == "json" {

		var hscr highscorestruc
		hscr.Title = "HIGHSCORE"

		for _, rec := range score {
			var a scorestruc
			a.Attempts = rec[0]
			a.User = rec[1]
			hscr.Scores = append(hscr.Scores, a)
		}

		json_data, err := json.MarshalIndent(hscr, "", "\t")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		w.Write(json_data)

	}
	if out == "yaml" || out == "yml" {

		var hscr highscorestruc
		hscr.Title = "HIGHSCORE"

		for _, rec := range score {
			var a scorestruc
			a.Attempts = rec[0]
			a.User = rec[1]
			hscr.Scores = append(hscr.Scores, a)
		}

		yaml_data, err := yaml.Marshal(hscr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		w.Write(yaml_data)
	}
}
