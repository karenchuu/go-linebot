package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Id             primitive.ObjectID
	BotUserId      string
	WebhookEventId string
	MessageId      string
	MessageType    string
	MessageText    string
	ReplyToken     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
