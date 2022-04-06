package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(env string) {
	if env != "" {
		viper.SetConfigName("config-" + env) // name of config file (without extension)
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
	}
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	log.Println(viper.GetString("msg"))

	// 配置文件变更时配置热更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		InitConfig(env)
	})
	// 监听配置文件变更的事件
	viper.WatchConfig()
}
