package main

import (
	"github.com/odeassis/goodoor/config"
	"github.com/odeassis/goodoor/routes"
)

var (
	logger *config.Looger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialize error: %v", err)
		return
	}

	routes.Initialize()
}
