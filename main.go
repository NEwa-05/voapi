package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloAlexa(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
	}
	var alexaResponse backendResponse
	err = json.Unmarshal(body, &alexaResponse)
	if err != nil {
		log.Print(err)
	}

	if alexaResponse.Request.Type == "LaunchRequest" {
		alexaResponse.Version = "1.0"
		alexaResponse.Response.OutputSpeech.Type = "PlainText"
		alexaResponse.Response.OutputSpeech.Text = "Bonjour, Humain !"
		alexaResponse.Response.ShouldEndSession = false
		alexaResponseByte, err := json.Marshal(alexaResponse)

		if err != nil {
			log.Print(err)
		}
		w.Write(alexaResponseByte)

	} else if alexaResponse.Request.Type == "IntentRequest" {
		if alexaResponse.Request.Intent.Name == "DoomPrediction" {

			alexaResponse.Version = "1.0"
			alexaResponse.Response.OutputSpeech.Type = "PlainText"
			alexaResponse.Response.OutputSpeech.Text = "En 2020, forme de vie inferieur"
			alexaResponse.Response.ShouldEndSession = true
			alexaResponseByte, err := json.Marshal(alexaResponse)

			if err != nil {
				log.Print(err)
			}

			w.Write(alexaResponseByte)
		}
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

func main() {
	callswapi()
	r := mux.NewRouter()
	r.HandleFunc("/alpha/alexa", helloAlexa).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
