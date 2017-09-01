package mapRepo

import (
	"time"
	"strconv"
	"fmt"
	"errors"
	. "whoscoming/domain"
)

var (
	trainings = make(map[string]Training)
	nextId = trainingIdSeq()
)

func CreateTraining(title string, location string, trainingTime time.Time, creatingUser string) Training {
	newId := nextId()
	p := []string{creatingUser}
	training := Training{Id: newId, Title: title, Location: location, TrainingTime: trainingTime, Participants: p}
	trainings[newId] = training
	fmt.Println(training)
	return training
}

func Participate(trainingId string, userName string) ([]string, error) {
	training, found := trainings[trainingId]
	if found {
		training.Participants = append(training.Participants, userName)
		trainings[trainingId] = training
		return training.Participants, nil
	} else {
		return nil, errors.New("No training found for id: " + trainingId)
	}
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

