package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `json:"username"`
	Password string             `json:"-"`
}
