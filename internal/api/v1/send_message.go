package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func SendMessage(c *gin.Context) {
	bot, err := linebot.New(viper.GetString("linebot.channelSecret"), viper.GetString("linebot.channelAccessToken"))
	log.Println("Bot:", bot, " err:", err)

	// Parse the incoming JSON payload
	var request struct {
		BotUserId string `json:"bot_user_id" binding:"required"`
		Message   string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Send the message
	to, err := bot.PushMessage(request.BotUserId, linebot.NewTextMessage(request.Message)).Do()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error sending message"})
		log.Print(fmt.Sprintf("error sending message to [%s], err = %s", to, err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "line message sent!"})
}
