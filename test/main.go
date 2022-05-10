package main

import (
	"errors"

	"github.com/klarkxy/logger"
)

func main() {
	logger.AddFile("log.txt", 3)
	logger.WithField("home", "klarkxy").Info("hello")
	logger.Debug("world")
	logger.WithError(errors.New("??????")).Error("hahaha")
}
