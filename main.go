package main

import (
	"github.com/Lailton-Abdon/opportunities_go/config"
	"github.com/Lailton-Abdon/opportunities_go/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}
