package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karenchuu/go-linebot/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type QueryParams struct {
	Page  int `form:"page,default=1" binding:"omitempty"`
	Limit int `form:"limit,default=10" binding:"omitempty"`
}

func GetUsersHandler(c *gin.Context) {
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

	usersColl := db.GetCollection(client, "users")
	findOptions := options.Find().SetSkip(int64((params.Page - 1) * params.Limit)).SetLimit(int64(params.Limit)) // Create the options for the Find() method
	cursor, err := usersColl.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while retrieving users"})
		return
	}
	defer cursor.Close(ctx)

	var users []bson.M
	for cursor.Next(ctx) {
		user := bson.M{}
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": users})
}
