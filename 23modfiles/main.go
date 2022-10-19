package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the class on mods")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serverfunction)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Welcome to the mode-----")
}

func serverfunction(w http.ResponseWriter, l *http.Request) {
	w.Write([]byte("<h1>Welcome to mod class:"))
}
