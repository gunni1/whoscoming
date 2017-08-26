package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
)

const (
	port = ":7000"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	//router.HandleFunc("/training", )

	fmt.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello")
}
