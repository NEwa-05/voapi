package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/answers", getAnswers).Methods("GET")
	r.HandleFunc("/questions", getQuestions).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
