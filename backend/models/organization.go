package models

import (
	"context"
	"log"
	"rides/util"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Organization struct {
	ID       string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string   `json:"name,omitempty" bson:"name,omitempty"`
	Password string   `json:"password" bson:"password"`
	Events   []uint16 `json:"events" bson:"events"`   // ids of events
	Drivers  []uint16 `json:"drivers" bson:"drivers"` // ids of drivers
}

type OrganizationClaims struct {
	ID        string    `json:"id"`
	jwt.StandardClaims
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

func OrgToken(org Organization) string {
	orgPayload := &OrganizationClaims{
		ID:        org.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, orgPayload)
	tokenString, err := token.SignedString([]byte(util.AccessToken()))
	if err != nil {
		log.Printf("%v", err)
		return ""
	}

	return tokenString
}

func CreateOrg(org Organization) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(org.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	org.Password = string(hashedPassword)

	_, err = OrgCollection().InsertOne(ctx, org)

	return err
}

func FindOrg(id string) (Organization, error) {
	orgCollection := OrgCollection()

	var orgB bson.M
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := orgCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&orgB)
	if err != nil {
		return Organization{}, err
	}

	var org Organization
	bsonBytes, _ := bson.Marshal(orgB)
	bson.Unmarshal(bsonBytes, &org)

	return org, err
}
