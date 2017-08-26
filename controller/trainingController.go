package controller

import (
	"net/http"
	"encoding/json"
)

type CreateTrainingDto struct {
	Title string
	Location string
	Date	string
}

func CreateTrainingHandler(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var createTrainingDto CreateTrainingDto
	error := decoder.Decode(&createTrainingDto)
	if error != nil {
		http.Error(response, error.Error(), http.StatusBadRequest)
	} else {
		println("parsed")
		//avatar := repo.CreateAvatar(createAvatarDto.Name)
		//json.NewEncoder(response).Encode(avatar)
	}
}
