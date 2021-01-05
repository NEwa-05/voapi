package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func getSwapi(s string) (peopleInfo SearchPeople) {
	//debug variable info
	fmt.Printf("%s\n", s)

	var peopleSearchInfo SearchPeople
	swapiBase, err := url.Parse("https://swapi.dev/api/people/")
	if err != nil {
		log.Print(err)
	}
	params := url.Values{}
	params.Add("search", s)
	swapiBase.RawQuery = params.Encode()

	//debug
	fmt.Printf("%s\n", swapiBase.String())

	swapiSearch, err := http.Get(swapiBase.String())
	if err != nil {
		log.Print(err)
	}

	err = json.NewDecoder(swapiSearch.Body).Decode(&peopleSearchInfo)
	if err != nil {
		log.Print(err)
	}

	//debug results
	fmt.Printf("%s\n", peopleSearchInfo.Results[0].Name)

	return peopleSearchInfo

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

	fmt.Print(alexaResponse)

	if alexaResponse.Request.Type == "LaunchRequest" {
		alexaResponse.Version = "1.0"
		alexaResponse.Response.OutputSpeech.Type = "PlainText"
		alexaResponse.Response.OutputSpeech.Text = "The dark side is full of informations"
		alexaResponse.Response.ShouldEndSession = false
		alexaResponseByte, err := json.Marshal(alexaResponse)

		if err != nil {
			log.Print(err)
		}
		w.Write(alexaResponseByte)

	} else if alexaResponse.Request.Type == "IntentRequest" {
		if alexaResponse.Request.Intent.Name == "int_sw_people_selection" {
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
						swapiInfo := getSwapi(newSlot.Resolutions.ResolutionsPerAuthority[0].Values[0].Value.Name)
						fmt.Printf("%s\n", swapiInfo.Results[0].Name)
						alexaResponse.Version = "1.0"
						alexaResponse.Response.OutputSpeech.Type = "PlainText"
						alexaResponse.Response.OutputSpeech.Text = "Luke you said"
						alexaResponse.Response.ShouldEndSession = false
						alexaResponseByte, err := json.Marshal(alexaResponse)
						if err != nil {
							log.Print(err)
						}

						w.Write(alexaResponseByte)
					}
				}
			}
		} else if alexaResponse.Request.Intent.Name == "int_sw_people_info" {
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
					if key == "intsl_sw_people_haircolor" {
						alexaResponse.Version = "1.0"
						alexaResponse.Response.OutputSpeech.Type = "PlainText"
						alexaResponse.Response.OutputSpeech.Text = "blond"
						alexaResponse.Response.ShouldEndSession = true
						alexaResponseByte, err := json.Marshal(alexaResponse)
						if err != nil {
							log.Print(err)
						}

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
