package domain

import (
	"time"
)

type Training struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Location string `json:"location"`
	TrainingTime time.Time `json:"trainingTime"`
	Participants []string `json:"participants"`
}

type Trainings []Training