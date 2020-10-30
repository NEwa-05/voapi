package main

//Requests structure
type Requests struct {
	Request string `json:"request"`
}

//Replies structure
type Replies struct {
	Reply string `json:"reply"`
}

//var replies []Replies = []Replies{}
var requests = []Requests{{"2021"}, {"2022"}}

var replies = []Replies{{"fire"}, {"pandemic"}}
