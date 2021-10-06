package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func MapUrls(r *mux.Router) {

	r.HandleFunc("/", Home).Methods(http.MethodGet)

	r.HandleFunc("/scoreboard", Scoreall).Methods(http.MethodGet)
	r.HandleFunc("/scoreboard/{player}", Score).Methods(http.MethodGet)

}
