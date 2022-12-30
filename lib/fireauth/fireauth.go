package fireauth

import (
	"context"
	"strings"
	"sysmo/lib/env"
	"sysmo/lib/log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var opt = option.WithCredentialsFile(env.CONF["FIREBASE_CONFIG_PATH"])
var app, appError = firebase.NewApp(context.Background(), nil, opt)

func init() {
	if appError != nil {
		log.Errorln("Firebase app initialization failed!.")
	} else {
		log.Defaultln("Firebase app initialized.")
	}
}

func ValidateFirebase(idToken string) (email string, provider string) {
	fireAuthApp, authError := app.Auth(context.Background())
	if authError != nil {
		return "", ""
	}

	token, err := fireAuthApp.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return "", ""
	}

	claimedEmail := token.Claims["email"].(string)
	if strings.Contains(claimedEmail, "#") {
		claimedEmail = DecodeEmail(claimedEmail)
	}
	return claimedEmail, token.Firebase.SignInProvider
}

func DecodeEmail(s string) string {
	spited := strings.Split(s, "#")
	spitedEmail := spited[0]
	indexOfLast := strings.LastIndexAny(spitedEmail, "_")
	email := strings.Join([]string{spitedEmail[:indexOfLast], spitedEmail[indexOfLast+1:]}, "@")
	return email
}
