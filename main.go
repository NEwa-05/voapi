package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//SwapiInfoID ...
var SwapiInfoID string

//SwapiInfoName ...
var SwapiInfoName string

func getSwapi(id, caracteristic string) (result string) {
	var peopleSearchInfo People
	swapiSearch, err := http.Get("https://swapi.dev/api/people/" + id)
	if err != nil {
		log.Print(err)
	}
	err = json.NewDecoder(swapiSearch.Body).Decode(&peopleSearchInfo)
	if err != nil {
		log.Print(err)
	}
	switch caracteristic {
	case "intsl_sw_people_height":
		result := peopleSearchInfo.Height
		return result
	case "intsl_sw_people_birthdate":
		result := peopleSearchInfo.BirthYear
		return result
	case "intsl_sw_people_haircolor":
		result := peopleSearchInfo.HairColor
		return result
	}
	return result
}

func buildIntentResponse(response string) (alexaResponseByte []byte) {
	var alexaResponse backendResponse
	alexaResponse.Version = "1.0"
	alexaResponse.Response.OutputSpeech.Type = "PlainText"
	alexaResponse.Response.OutputSpeech.Text = fmt.Sprintf("%s", response)
	alexaResponse.Response.ShouldEndSession = true
	alexaResponseByte, err := json.Marshal(alexaResponse)
	if err != nil {
		log.Print(err)
	}
	return alexaResponseByte
}

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
		response := "Holo disque initialisé"
		alexaResponseByte := buildIntentResponse(response)
		w.Write(alexaResponseByte)
	} else if alexaResponse.Request.Type == "IntentRequest" {
		if alexaResponse.Request.Intent.Name == "int_sw_people" {

			for key := range alexaResponse.Request.Intent.Slots {
				var newSlot Slot
				s, err := json.Marshal(alexaResponse.Request.Intent.Slots[key])
				if err != nil {
					fmt.Print("can not marshal")
				}
				err = json.Unmarshal(s, &newSlot)
				if err != nil {
					fmt.Print("can not unmarshal new value")
				}
				if newSlot.Resolutions != nil {
					if key == "intsl_sw_people_name" {
						SwapiInfoName = newSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.Name
						SwapiInfoID = newSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
					}
				}
			}
			for key := range alexaResponse.Request.Intent.Slots {
				var newSlot Slot
				s, err := json.Marshal(alexaResponse.Request.Intent.Slots[key])
				if err != nil {
					fmt.Print("can not marshal")
				}
				err = json.Unmarshal(s, &newSlot)
				if err != nil {
					fmt.Print("can not unmarshal new value")
				}
				if newSlot.Resolutions != nil {
					switch key {
					case "intsl_sw_people_height":
						swapiInfo := getSwapi(SwapiInfoID, key)
						response := fmt.Sprintf("la taille de %s est de %s centimetres", SwapiInfoName, swapiInfo)
						alexaResponseByte := buildIntentResponse(response)
						w.Write(alexaResponseByte)
						/*_, err = w.Write(alexaResponseByte)
						if err == nil {
							SwapiInfoName = ""
							SwapiInfoID = ""
						}*/
					case "intsl_sw_people_birthdate":
						swapiInfo := getSwapi(SwapiInfoID, key)
						response := fmt.Sprintf("l'année de naissance de %s est %s", SwapiInfoName, swapiInfo)
						alexaResponseByte := buildIntentResponse(response)
						w.Write(alexaResponseByte)
						/*_, err = w.Write(alexaResponseByte)
						if err == nil {
							SwapiInfoName = ""
							SwapiInfoID = ""
						}*/
					case "intsl_sw_people_haircolor":
						swapiInfo := getSwapi(SwapiInfoID, key)
						switch swapiInfo {
						case "n/a":
							response := fmt.Sprintf("Les droids n'ont pas de cheveux")
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						case "brown":
							response := fmt.Sprintf("les cheveux de %s sont brun", SwapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)

						case "blond":
							response := fmt.Sprintf("les cheveux de %s sont %s", SwapiInfoName, swapiInfo)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)

						}
						/*_, err = w.Write(alexaResponseByte)
						if err == nil {
							SwapiInfoName = ""
							SwapiInfoID = ""
						}*/

					}
				}
			}
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/alpha/alexa", helloAlexa).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}

}
