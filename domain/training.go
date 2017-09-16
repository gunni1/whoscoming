package domain

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Training struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	Location string `json:"location" bson:"location"`
	TrainingTime time.Time `json:"trainingTime" bson:"trainingTime"`
	Participants []string `json:"participants" bson:"participants"`
}

type Trainings []Training