package repo

import (
	"time"
	"strconv"
	"fmt"
)

var (
	trainings = make(map[string]Training)
	nextId = trainingIdSeq()
)

type Training struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Location string `json:"location"`
	TrainingTime time.Time `json:"trainingTime"`
	Participants []string `json:"participants"`
}

type Trainings []Training

func CreateTraining(title string, location string, trainingTime time.Time, participant string) Training {
	newId := nextId()
	p := []string{participant}
	training := Training{Id: newId, Title: title, Location: location, TrainingTime: trainingTime, Participants: p}
	trainings[newId] = training
	fmt.Println(training)
	return training
}

func GetTraining(trainingId string) (Training, bool){
	training, found := trainings[trainingId]
	return training, found
}

func GetTrainings() Trainings {
	results := make(Trainings, len(trainings))
	i := 0
	for _, value := range trainings {
		results[i] = value
		i++
	}
	return results
}

func trainingIdSeq() func() string {
	i := 0
	return func() string {
		i += 1
		return strconv.Itoa(i)
	}
}

