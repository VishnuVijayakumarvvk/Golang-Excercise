package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println("Welcome to the first class on web request----")
	responce, err := http.Get(url)
	errorfinder(err)
	fmt.Printf("The type is : %T", responce)
	data,erro := ioutil.ReadAll(responce.Body)
	errorfinder(erro)
	fmt.Println("The data inside the file is : \n",string(data))
	defer responce.Body.Close()


	// responce, err := http.Get(url)
	// errorfinder(err)
	// fmt.Printf("The type of responce is : %T", responce)
	// data, erro := ioutil.ReadAll(responce.Body)
	// errorfinder(erro)
	// fmt.Println("The content in the file is : ", string(data))
	// defer responce.Body.Close()
}

func errorfinder(err error) {
	if err != nil {
		panic(err)
	}
}
