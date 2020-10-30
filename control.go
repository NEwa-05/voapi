package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func getYears(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	var foo request
	err = json.Unmarshal(body, &foo)
	if err != nil {
		log.Print(err)
	}
	year, err := strconv.Atoi(foo.Text)
	if err != nil {
		log.Print(err)
	}
	var bar response

	switch {
	case year <= 2021:
		bar.Text = "l'année est déjà passé"
	case year > 2500:
		bar.Text = "la Terre a été detruite"
	default:
		bar.Text = fmt.Sprintf("l'année %d n'est pas encore implémenté", year)
	}

	toto, err := json.Marshal(bar.Text)
	if err != nil {
		log.Print(err)
	}

	w.Write(toto)

}
