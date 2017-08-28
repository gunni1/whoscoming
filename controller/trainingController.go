package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"whoscoming/repo"
)

var (
	homeLocation, _ = time.LoadLocation("Europe/Berlin")
)

type CreateTrainingDto struct {
	Title        string
	Location     string
	TrainingTime string
	CreatingUser string
}

func CreateTrainingHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var createTrainingDto CreateTrainingDto
	error := decoder.Decode(&createTrainingDto)

	if error != nil {
		http.Error(response, error.Error(), http.StatusBadRequest)
	} else {
		trainingTime, error := time.Parse(time.RFC3339, createTrainingDto.TrainingTime)
		if error != nil {
			http.Error(response, error.Error(), http.StatusBadRequest)
		} else {
			training := repo.CreateTraining(createTrainingDto.Title, createTrainingDto.Location, trainingTime, createTrainingDto.CreatingUser)
			json.NewEncoder(response).Encode(training)
		}
	}
}

func GetTrainingHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	trainingId := vars["trainingId"]

	avatar, found := repo.GetTraining(trainingId)
	if found {
		json.NewEncoder(response).Encode(avatar)
	} else {
		error := "training with id: " + trainingId + " not found."
		http.Error(response, error, http.StatusBadRequest)
	}
}

func GetTrainingsHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(repo.GetTrainings())
}
