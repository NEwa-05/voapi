package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	callswapi()
	r := mux.NewRouter()
	r.HandleFunc("/alpha/events", getEvents).Methods("POST")
	r.HandleFunc("/alpha/alexa", helloAlexa).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
