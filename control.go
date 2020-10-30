package main

import (
	"encoding/json"
	"net/http"
)

func getAnswers(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(replies)
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(requests)
}
