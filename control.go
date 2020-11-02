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

func getEvents(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	var faa request
	err = json.Unmarshal(body, &faa)
	if err != nil {
		log.Print(err)
	}
	event := faa.Text

	var bur response

	switch {
	case event == "pandémie":
		bur.Text = "En 2021 la covid19 decimera la moitié de la population mondiale"
	case event == "aliens":
		bur.Text = "Les aliens ont débarqué en 2029"
	default:
		bur.Text = "je n'ai aps connaissance d'un evenement de ce genre"
	}

	toto, err := json.Marshal(bur.Text)
	if err != nil {
		log.Print(err)
	}

	w.Write(toto)
}
