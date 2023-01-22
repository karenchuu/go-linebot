package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karenchuu/go-linebot/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMessagesByUserIdHandler(c *gin.Context) {
	type QueryParams struct {
		BotUserId string `form:"bot_user_id" binding:"required"`
		Page      int    `form:"page,default=1" binding:"omitempty"`
		Limit     int    `form:"limit,default=10" binding:"omitempty"`
	}
	var params QueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messages": err.Error()})
		return
	}

	client, ctx, err := db.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messages": fmt.Sprintf("Error connect db, err = %s", err)})
		return
	}

	messagesColl := db.GetCollection(client, "messages")
	filter := bson.M{"botuserid": params.BotUserId}                                                              // Build the filter to query for the user's messages
	findOptions := options.Find().SetSkip(int64((params.Page - 1) * params.Limit)).SetLimit(int64(params.Limit)) // Create the options for the Find() method
	// Execute the query and get the list of messages
	cursor, err := messagesColl.Find(ctx, filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"messages": "Error executing query"})
		return
	}
	defer cursor.Close(ctx)

	// Iterate through the cursor and build the list of messages
	var messages []bson.M
	for cursor.Next(ctx) {
		message := bson.M{}
		cursor.Decode(&message)
		messages = append(messages, message)
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": messages})
}
