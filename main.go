package main

import (
	"fmt"
	"mpesa-example/service"

	"github.com/spf13/viper"
)

var appName = "mpesa-example"

func init() {
	viper.SetConfigName("dev")
}
func main() {
	fmt.Printf("Starting %v\n", appName)
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err.Error()))
	}
	service.StartWebServer(viper.GetString("server_port"))
}
