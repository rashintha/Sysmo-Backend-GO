package usersService

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"Name"`
	Email     string             `json:"email" form:"email" bson:"email"`
	Password  string             `json:"password" form:"password" bson:"password"`
	Status    int                `json:"status" bson:"status"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
}
