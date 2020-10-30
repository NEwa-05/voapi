package main

import (
	"net/http"
)

func checkMethod(w http.ResponseWriter, r *http.Request) {
	// return content-type as json
	w.Header().Set("Content-Type", "application/json")
	//check if method is GET
	if r.Method != "GET" {
		// return header forbidden since we only want GET
		w.WriteHeader(http.StatusForbidden)
		// return header message as json
		w.Write([]byte(` { "message": "Only GET method accepted" } `))
	}
}
