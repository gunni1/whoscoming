package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"whoscoming/controller"
	"whoscoming/mongodb"
)

const (
	port = ":7000"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/training", controller.CreateTrainingHandler).Methods("POST")
	router.HandleFunc("/training/{trainingId}", controller.GetTrainingHandler).Methods("GET")
	router.HandleFunc("/trainings", controller.GetTrainingsHandler).Methods("GET")
	router.HandleFunc("/training/{trainingId}/participate", controller.ParticipateHandler).Methods("POST")

	session := mongodb.OpenDbConnection()
	defer session.Close()

	fmt.Println("listening on " + port)
	log.Fatal(http.ListenAndServe(port, router))
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello")
}
