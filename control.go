package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getEvents(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	var faa prevresp
	err = json.Unmarshal(body, &faa)
	if err != nil {
		log.Print(err)
	}
	event := faa.Text

	var bur prevresp

	switch {
	case event == "pandémie":
		bur.Text = "En 2021 la covid19 decimera la moitié de la population mondiale"
	case event == "aliens":
		bur.Text = "Les aliens ont débarqué en 2029"
	default:
		bur.Text = "je n'ai pas connaissance d'un evenement de ce genre"
	}

	toto, err := json.Marshal(bur.Text)
	if err != nil {
		log.Print(err)
	}

	w.Write(toto)
}

func helloAlexa(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	var alexaLaunchRequest launchRequest
	err = json.Unmarshal(body, &alexaLaunchRequest)
	if err != nil {
		log.Print(err)
	}

	var alexaResponse backendResponse

	if alexaLaunchRequest.Request.Type == "LaunchRequest" {
		alexaResponse.Version = "1.0"
		alexaResponse.Response.OutputSpeech.Type = "PlainText"
		alexaResponse.Response.OutputSpeech.Text = "Bonjour Humain"
		alexaResponse.Response.ShouldEndSession = true
		alexaResponseByte, err := json.Marshal(alexaResponse)
		if err != nil {
			log.Print(err)
		}
		w.Write(alexaResponseByte)
	} else {
		alexaResponse.Version = "1.0"
		alexaResponse.Response.OutputSpeech.Type = "PlainText"
		alexaResponse.Response.OutputSpeech.Text = "Je n'ai pas compris"
		alexaResponse.Response.ShouldEndSession = true
		alexaResponseByte, err := json.Marshal(alexaResponse)
		if err != nil {
			log.Print(err)
		}
		w.Write(alexaResponseByte)
	}
}
