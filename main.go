package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"whoscoming/controller"
	"whoscoming/mongodb"
	"flag"
)



func main() {
	dbUrl := flag.String("mongodb","localhost:27017", "Connection String to a MongoDB")
	apiPort := flag.String("apiPort","7000","REST API Port")
	flag.Parse()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/training", controller.CreateTrainingHandler).Methods("POST")
	router.HandleFunc("/training/{trainingId}", controller.GetTrainingHandler).Methods("GET")
	router.HandleFunc("/trainings", controller.GetTrainingsHandler).Methods("GET")
	router.HandleFunc("/training/{trainingId}/participate", controller.ParticipateHandler).Methods("POST")

	session := mongodb.OpenDbConnection(*dbUrl)
	defer session.Close()

	fmt.Println("REST API listening on " + *apiPort)
	log.Fatal(http.ListenAndServe(*apiPort, router))
}

func index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hello")
}
