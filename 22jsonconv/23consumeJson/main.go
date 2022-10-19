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
	fmt.Println("Welcome to the program for consuming json data---")
	consumejson()
}

func consumejson()  {
	data := []byte(`
		{
		"coursename": "Angular",
		"Price": 300,
		"tags": ["web","frontend"]
		}
	`)
	json.Valid(data)
		//var coursedata course
		var coursedata map[string]interface{}
		json.Unmarshal(data,&coursedata)
		fmt.Printf("The data which is converted from json is %#v\n",coursedata)

		for i,val := range coursedata{
			fmt.Printf("The key is %v and the value is %v\n",i,val)
		}
}
