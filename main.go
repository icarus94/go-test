package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hompage hit")
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage).Methods("GET")
	myRouter.HandleFunc("/register", homepage).Methods("POST")
	myRouter.HandleFunc("/login", homepage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
