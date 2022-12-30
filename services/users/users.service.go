package usersService

import (
	"context"
	"sysmo/lib/db"
	"sysmo/lib/hashing"
	"sysmo/lib/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func VerifyUser(user User) (*User, string, error) {
	db := db.GetDB()
	coll := db.Database("sysmo").Collection("User")

	var retUser User
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&retUser)

	if err == mongo.ErrNoDocuments {
		log.Warningln("Unknown Email: " + user.Email + ". " + err.Error())
		return nil, "Invalid user.", err
	} else if err != nil {
		log.Errorln("Error in finding document: " + user.Email + ". " + err.Error())
		return nil, "Unknown error occurred.", err
	}

	match, err := hashing.VerifyPassword(user.Password, retUser.Password)

	if err != nil {
		log.Warningln("Error in password Email: " + user.Email + ". " + err.Error())
		return nil, "Error in password. ", err
	}

	if match {
		log.Defaultln("User verification success, Email: " + user.Email + ".")
		return &retUser, "Successfully verified.", nil
	}

	log.Warningln("User verification failed, Email: " + user.Email + ".")
	return nil, "Verification failed.", nil
}
