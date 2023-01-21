package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID
	Bot       string
	BotUserId string
	CreatedAt time.Time
	UpdatedAt time.Time
}
