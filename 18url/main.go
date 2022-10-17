package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the class on url-----------")
	var urllink string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
	madeurl, err := url.Parse(urllink)
	errorFinder(err)
	fmt.Println("The url : ", madeurl.Scheme)
	fmt.Println("The url : ", madeurl.Host)
	fmt.Println("The url : ", madeurl.Port())
	fmt.Println("The url : ", madeurl.Path)
	fmt.Println("The url : ", madeurl.RawQuery)
	queries := madeurl.Query()

	for _,val := range queries{
		fmt.Println("the queries are : ", val)
	}
fmt.Println("changing from parts to url------")
partOfUrl := &url.URL{
	Scheme:"https",
	Host: "lco.dev",
	Path: "learn",
	RawQuery: "coursename=reactjs&paymentid=ghbj456ghb",
}
fmt.Println("The given url  is : ",partOfUrl.String())

}
func errorFinder(err error) {
	if err != nil {
		panic(err)
	}
}
