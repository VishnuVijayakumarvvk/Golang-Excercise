package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the program on post forms------")
	const urllink = "http://localhost:1111/postform"
	postformreq(urllink)
}

func postformreq(urllink string) {
	data := url.Values{}
	data.Add("Name", "Vishnu")
	data.Add("Age", "24")
	data.Add("Gender", "male")
	data.Add("Status", "Available")
	responce, err := http.PostForm(urllink, data)
	if err != nil {
		panic(err)
	}
	var databyte strings.Builder
	datab, _ := ioutil.ReadAll(responce.Body)
	databb, _ := databyte.Write(datab)
	fmt.Println("The byte count is :", databb)
	fmt.Println("The post form is : ", databyte.String())
}
