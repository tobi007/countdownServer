package server

import (
	"github.com/tobi007/cdServer/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("serverAddress") +":" + config.GetString("serverPort"))
}