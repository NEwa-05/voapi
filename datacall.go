package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func callswapi() {
	swapiResponse, err := http.Get("https://swapi.dev/api/planets/8/")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	var response Planets
	err = json.NewDecoder(swapiResponse.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", response.Name)
	fmt.Printf("%s\n", response.Residents)
}
