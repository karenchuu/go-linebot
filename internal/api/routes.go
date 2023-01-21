package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/karenchuu/go-linebot/internal/api/v1"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/linecallback", LineCallbackHandler)

	apiv1 := r.Group("/v1")
	{
		apiv1.GET("/users", v1.GetUsersHandler)
		apiv1.GET("/messages", v1.GetMessagesByUserIdHandler)
		apiv1.POST("/sendMessage", v1.SendMessage)
	}
	return r
}
