package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name    string `json:"coursename"`
	Price   int
	Wesite  string `json:"-"`
	Content []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to th e program-- how to convert data to json---------")
	datas := []course{
		{"Angular", 300, "lco", []string{"web", "frontend"}},
		{"react", 400, "lco", []string{"web", "frontend"}},
		{"node js", 500, "lco", []string{"web", "frontend"}},
		{"goland", 650, "lco", []string{"web", "frontend"}},
	}
	
	// jsondatas, err := json.Marshal(datas)
	jsondatas, err := json.MarshalIndent(datas, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("The data is :", string(jsondatas))
}
