package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the file Writer and reader----------")

	content := "This is the file to go to www.facebook.com"
	file, err := os.Create("./fileReader.txt")
	errorFinder(err)
	length, erro := io.WriteString(file, content)
	errorFinder(erro)
	fmt.Println("The length of the string is : ", length)
	fileReader("./fileReader.txt")
	defer file.Close()
}

func fileReader(file string) {
	data, erre := ioutil.ReadFile(file)
	errorFinder(erre)
	fmt.Println("The file is :   ", string(data))
}

func errorFinder(err error) {
	if err != nil {
		panic(err)
	}
}
