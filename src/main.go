package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Timboschanski/Ratespiel/src/app"
	"github.com/Timboschanski/Ratespiel/src/controllers"
	"github.com/gorilla/mux"
)

//main Input of parameters and run functions
func main() {
	go server()
	app.StartGame(&os.Args)
}

//server Creates a Web Server
func server() {

	r := mux.NewRouter()

	controllers.MapUrls(r)

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}
