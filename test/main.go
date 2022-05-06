package main

import (
	"errors"

	"github.com/klarkxy/logger"
)

func main() {
	log := logger.New(logger.DebugLevel)
	log.AddFile("log.txt", 3)
	log.WithField("home", "klarkxy").Info("hello")
	log.Debug("world")
	log.WithError(errors.New("I don't know")).Error("hahaha")
}
