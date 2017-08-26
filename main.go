package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"whoscoming/controller"
)

const (
	port = ":7000"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/training", controller.CreateTrainingHandler)
	router.HandleFunc("/training/{trainingId}", controller.GetTrainingHandler)
	router.HandleFunc("/trainings", controller.GetTrainingsHandler)

	fmt.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello")
}
