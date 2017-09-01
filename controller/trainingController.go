package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"whoscoming/mapRepo"
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

type ParticipateDto struct {
	UserName string
}

func CreateTrainingHandler(response http.ResponseWriter, request *http.Request) {
	var createTrainingDto CreateTrainingDto
	decoder := json.NewDecoder(request.Body)
	decodeError := decoder.Decode(&createTrainingDto)

	if decodeError != nil {
		http.Error(response, decodeError.Error(), http.StatusBadRequest)
	} else {
		trainingTime, parseError := time.Parse(time.RFC3339, createTrainingDto.TrainingTime)
		if parseError != nil {
			http.Error(response, parseError.Error(), http.StatusBadRequest)
		} else {
			training := mapRepo.CreateTraining(createTrainingDto.Title, createTrainingDto.Location, trainingTime, createTrainingDto.CreatingUser)
			json.NewEncoder(response).Encode(training)
		}
	}
}

func GetTrainingHandler(response http.ResponseWriter, request *http.Request) {
	trainingId := extractRequestVar("trainingId", request)

	avatar, found := mapRepo.GetTraining(trainingId)
	if found {
		json.NewEncoder(response).Encode(avatar)
	} else {
		error := "training with id: " + trainingId + " not found."
		http.Error(response, error, http.StatusBadRequest)
	}
}

func GetTrainingsHandler(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(mapRepo.GetTrainings())
}

func ParticipateHandler(response http.ResponseWriter, request *http.Request) {
	var participateDto ParticipateDto
	decodeError := json.NewDecoder(request.Body).Decode(&participateDto)

	if decodeError != nil {
		http.Error(response, decodeError.Error(), http.StatusBadRequest)
	} else {
		trainingId := extractRequestVar("trainingId", request)
		updatedTraining, error := mapRepo.Participate(trainingId, participateDto.UserName)
		if error != nil {
			http.Error(response, error.Error(), http.StatusBadRequest)
		} else {
			json.NewEncoder(response).Encode(updatedTraining)
		}
	}
}

func extractRequestVar(variable string, request *http.Request) string {
	vars := mux.Vars(request)
	return vars[variable]
}
