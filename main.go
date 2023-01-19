package main

import (
	"github.com/karenchuu/go-linebot/internal/api"
)

func main() {
	r := *api.GetRouter()
	r.Run(":8080")
}
