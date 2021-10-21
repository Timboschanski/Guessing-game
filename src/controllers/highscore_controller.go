package controllers

import (
	"fmt"
	"net/http"

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

	w.Write([]byte("Scoreboard\n"))
	if out == "json" || out == "yaml" || out == "yml" {
		fmt.Fprint(w, "Out = ", out, "\n")
	}
	w.Write([]byte(s))
}
