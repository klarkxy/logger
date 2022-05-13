package main

import (
	"errors"

	"github.com/klarkxy/logger"
)

type XXX struct {
	Str    string `json:"string"`
	PtrInt *int   `json:"ptr_int"`
	Int    int
}

func main() {
	logger.AddFile("log.txt", 3)
	logger.WithField("home", "klarkxy").Info("hello")
	logger.Debug("world")
	logger.WithError(errors.New("??????")).Error("hahaha")
	logger.WithField("xxx", &XXX{Str: "hello", PtrInt: nil, Int: 1}).Info("wakaka")
}
