package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karenchuu/go-linebot/db"
	"github.com/karenchuu/go-linebot/internal/models"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Line    = "line"
	Text    = "text"
	Sticker = "sticker"
)

func LineCallbackHandler(c *gin.Context) {
	bot, err := linebot.New(viper.GetString("linebot.channelSecret"), viper.GetString("linebot.channelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		}
		return
	}

	client, ctx, err := db.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error connect db, err = %s", err)})
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			messageType := ""
			messageText := ""
			messageId := ""
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
				quota, err := bot.GetMessageQuota().Do()
				if err != nil {
					log.Println("Quota err:", err)
				}

				messageId = message.ID
				messageType = Text
				messageText = message.Text

				// Send reply to sender
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("已收到您的訊息！")).Do()
				if err != nil {
					log.Print(err)
				}
				log.Println("msg ID: " + message.ID + ";" + "Get: " + message.Text + " , \n OK! remain message:" + strconv.FormatInt(quota.Value, 10))
			default:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("已收到您的訊息！")).Do(); err != nil {
					log.Print(err)
				}
			}

			// Save message to DB
			usersCollection := db.GetCollection(client, "users")
			userCount, err := usersCollection.CountDocuments(ctx, bson.M{"botuserid": event.Source.UserID})
			if err != nil {
				log.Fatal(fmt.Sprintf("userCount = %s, err: %s", fmt.Sprint(userCount), err))
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
				return
			}

			if userCount == 0 {
				now := time.Now()
				user := models.User{
					Id:        primitive.NewObjectID(),
					Bot:       Line,
					BotUserId: event.Source.UserID,
					CreatedAt: now,
					UpdatedAt: now,
				}
				result, err := usersCollection.InsertOne(ctx, user)
				if err != nil {
					log.Print(err)
					c.JSON(http.StatusInternalServerError, gin.H{"message": "usersCollection Insert Error"})
					return
				}
				fmt.Println(result)
			}

			messagesCollection := db.GetCollection(client, "messages")
			now := time.Now()
			msg := models.Message{
				Id:             primitive.NewObjectID(),
				BotUserId:      event.Source.UserID,
				WebhookEventId: event.WebhookEventID,
				MessageId:      messageId,
				MessageType:    messageType,
				MessageText:    messageText,
				ReplyToken:     event.ReplyToken,
				CreatedAt:      now,
				UpdatedAt:      now,
			}
			result, err := messagesCollection.InsertOne(ctx, msg)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(result)

			defer client.Disconnect(ctx)
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
