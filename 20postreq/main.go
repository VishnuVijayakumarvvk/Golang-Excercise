package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the class of post request---------")
	urlLink := "http://localhost:1111/post"
	postreq(urlLink)
}

func postreq(urllink string) {
	jsonfile := strings.NewReader(`
		{
			"name" : "Vishnu",
  			"age" : 24,
  			"gender" : "male"
		}
	`)

	responce, err := http.Post(urllink, "application/json", jsonfile)
	if err != nil {
		panic(err)
	}
	//var newst strings.Builder
	datab, _ := ioutil.ReadAll(responce.Body)
	// defer responce.Body.Close()
	// databyte, _ := newst.Write(datab)
	// fmt.Println("the byte vvalue is :", databyte)
	fmt.Println("The string formate is : \n", string(datab))
}
