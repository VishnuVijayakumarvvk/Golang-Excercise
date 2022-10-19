package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the class of web request----------")
	const url = "http://localhost:1111/get"
	getter(url)
}

func getter(url string) {
	responce, err := http.Get(url)
	if err!= nil{
		panic(err)
	}
	data, _ := ioutil.ReadAll(responce.Body)
	defer responce.Body.Close()
	var datastring strings.Builder
	databyte,_ := datastring.Write(data)
	fmt.Println("The value in databyte is : ", databyte)
	fmt.Println("The string value in web server is : ", datastring.String())
}

