package mongodb

import (
	"time"
	"errors"
	. "whoscoming/domain"
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)

const (
	trainingDbName = "whoscomingDB"
	trainingCollectionName = "trainings"
)

var(
	//Alternativ: type TrainingDB
	db *mgo.Database
)

func OpenDbConnection(dbUrl string) mgo.Session {
	log.Println("connect to mongodb using url: " + dbUrl)
	session, error := mgo.Dial(dbUrl)
	if error != nil {
		panic(error)
	}
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(trainingDbName)
	return *session
}

func CreateTraining(title string, location string, trainingTime time.Time, creatingUser string) Training {
	trainings := db.C(trainingCollectionName)
	parts := []string{creatingUser}
	training := Training{Title: title, Location: location, TrainingTime: trainingTime, Participants: parts}
	err := trainings.Insert(&training)
	if err != nil {
		log.Println(err)
	}
	return training
}

func Participate(trainingId string, userName string) ([]string, error) {
	trainings := db.C(trainingCollectionName)
	training, found  := GetTraining(trainingId)

	if found {
		training.Participants = append(training.Participants, userName)
		trainings.UpdateId(bson.ObjectIdHex(trainingId), training)
		return training.Participants, nil
	} else {
		return nil, errors.New("No training found for id: " + trainingId)
	}
}

func GetTraining(trainingId string) (Training, bool){
	trainings := db.C(trainingCollectionName)
	result := Training{}

	err := trainings.FindId(bson.ObjectIdHex(trainingId)).One(&result)

	if err != nil {
		log.Println(err)
		return result, false
	} else{
		return result, true
	}
}

func GetTrainings() Trainings {
	trainings := db.C(trainingCollectionName)

	result := Trainings{}
	trainings.Find(bson.M{}).All(&result)

	return result
}