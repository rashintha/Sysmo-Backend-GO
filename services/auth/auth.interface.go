package authService

import (
	"time"
)

type LoginHistory struct {
	UserID   string    `bson:"userId, omitempty"`
	DateTime time.Time `bson:"dateTime, omitempty"`
}

type AuthUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
