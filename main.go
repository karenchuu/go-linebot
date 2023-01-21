package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/karenchuu/go-linebot/internal/api"
	"github.com/spf13/viper"
)

func main() {
	r := *api.GetRouter()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	addr := fmt.Sprintf(":%s", viper.GetString("server.port"))
	r.Run(addr)
}
