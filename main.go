package main

import (
	"github.com/gersonpb/go-opportunities/config"
	"github.com/gersonpb/go-opportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	// Initialize configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	// Initialized router
	router.Initialize()
}
