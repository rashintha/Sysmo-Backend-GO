package authService

import (
	"context"
	"sysmo/lib/db"
	"sysmo/lib/log"
	"time"
)

func PostUserSignInProcess(userID string) (string, error) {
	db := db.GetDB()

	coll := db.Database("sysmo").Collection("LoginHistory")
	loginHistoryRecord := LoginHistory{UserID: userID, DateTime: time.Now()}

	_, err := coll.InsertOne(context.TODO(), loginHistoryRecord)
	if err != nil {
		log.Errorln("Error saving login history data. " + err.Error())
		return "Error in saving login history data.", err
	}

	return "Login history saved.", nil
}
