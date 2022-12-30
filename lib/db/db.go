package db

import (
	"context"
	"sysmo/lib/env"
	"sysmo/lib/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func init() {
	open()
}

func open() {
	uri := env.CONF["DATABASE_URL"]

	if uri == "" {
		log.ErrorFatal("Error in your environment file. DATABASE_URL not found.")
	}

	var err error

	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.ErrorFatal("Error in connecting to MongoDB instance: " + err.Error())
	}
}

func ping() error {
	err := db.Ping(context.TODO(), nil)

	if err == mongo.ErrClientDisconnected {
		open()
	} else if err != nil {
		log.Errorln("Error in pinging to MongoDB instance: " + err.Error())
		return err
	}

	return nil
}

func GetDB() *mongo.Client {
	ping()

	return db
}
