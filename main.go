package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

func alexaSlotParsing(slots interface{}) (newSlot slot) {
	slotBytes, err := json.Marshal(slots)
	if err != nil {
		fmt.Print("cannot marshal new slot")
	}
	err = json.Unmarshal(slotBytes, &newSlot)
	if err != nil {
		fmt.Print("cannot unmarshal new value")
	}
	return newSlot
}

func getSlotPeople(alexaResponse backendResponse) (id, name string) {
	nameSlot := alexaSlotParsing(alexaResponse.Request.Intent.Slots["intsl_sw_people_name"])
	name = nameSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.Name
	id = nameSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.ID
	return id, name
}

func helloAlexa(w http.ResponseWriter, r *http.Request) {
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
	switch alexaResponse.Request.Type {
	case "LaunchRequest":
		response := "Holo disque initialisé"
		alexaResponseByte := buildIntentResponse(response)
		w.Write(alexaResponseByte)
	case "IntentRequest":
		if alexaResponse.Request.Intent.Name == "int_sw_people" {
			swapiInfoID, swapiInfoName := getSlotPeople(alexaResponse)
			for key := range alexaResponse.Request.Intent.Slots {
				newSlot := alexaSlotParsing(alexaResponse.Request.Intent.Slots[key])
				if newSlot.Resolutions != nil {
					switch key {
					case "intsl_sw_people_name":
					case "intsl_sw_people_height":
						swapiInfo := getSwapi(swapiInfoID, key)
						response := fmt.Sprintf("la taille de %s est de %s centimetres", swapiInfoName, swapiInfo)
						alexaResponseByte := buildIntentResponse(response)
						w.Write(alexaResponseByte)
					case "intsl_sw_people_birthdate":
						swapiInfo := getSwapi(swapiInfoID, key)
						response := fmt.Sprintf("l'année de naissance de %s est %s", swapiInfoName, swapiInfo)
						alexaResponseByte := buildIntentResponse(response)
						w.Write(alexaResponseByte)
					case "intsl_sw_people_haircolor":
						swapiInfo := getSwapi(swapiInfoID, key)
						switch swapiInfo {
						case "brown":
							response := fmt.Sprintf("les cheveux de %s sont bruns", swapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						case "blond":
							response := fmt.Sprintf("les cheveux de %s sont blonds", swapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						case "red":
							response := fmt.Sprintf("les cheveux de %s sont roux", swapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						case "none":
							response := fmt.Sprintf("%s n'a pas de cheveux sur le caillou", swapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						default:
							response := fmt.Sprintf("aucune info sur les cheveux de %s", swapiInfoName)
							alexaResponseByte := buildIntentResponse(response)
							w.Write(alexaResponseByte)
						}
					default:
						response := fmt.Sprintf("Je ne connais pas cette information à propos de %s", swapiInfoName)
						alexaResponseByte := buildIntentResponse(response)
						w.Write(alexaResponseByte)
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
