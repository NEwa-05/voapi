package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	swapistructs "voapi/swapi"
)

func callswapi() {
	swapiResponse, err := http.Get("https://swapi.dev/api/people/1/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	var response swapistructs.People
	err = json.NewDecoder(swapiResponse.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", response.Name)
	fmt.Printf("%s\n", response.Mass)
}
