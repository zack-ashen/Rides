package models

import (
	"context"
	"log"
	"rides/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Organization struct {
	ID       string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name,omitempty" bson:"name,omitempty"`
	Password string   `json:"password" bson:"password"`
	Events   []uint16 `json:"events" bson:"events"`   // ids of events
	Drivers  []uint16 `json:"drivers" bson:"drivers"` // ids of drivers
}

func OrgCollection() *mongo.Collection {
	client, err := util.GetMongoClient()
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("rides").Collection("organizations")
}

func UpdateOrg(org Organization) error {

	return nil
}

func CreateOrg(org Organization) error {

	return nil
}

func FindOrg(id string) (Organization, error) {
	orgCollection := OrgCollection()

	var orgB bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := orgCollection.FindOne(ctx, bson.M{"ID": id}).Decode(&orgB)
	if err != nil {
		return Organization{}, err
	}

	var org Organization
	bsonBytes, _ := bson.Marshal(orgB)
	bson.Unmarshal(bsonBytes, &org)

	return org, err
}
