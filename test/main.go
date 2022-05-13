package main

import (
	"errors"

	"github.com/klarkxy/logger"
)

type XXX struct {
	Str     string `yaml:"string"`
	PtrInt  *int   `yaml:"ptr_int,omitempty"`
	PtrInt2 *int   `yaml:"ptr_int2,omitempty"`
	Int     int
}

func main() {
	logger.AddFile("log.txt", 3)
	logger.WithField("home", "klarkxy").Info("hello")
	logger.Debug("world")
	logger.WithError(errors.New("??????")).Error("hahaha")
	logger.WithField("xxx", &XXX{Str: "hello", PtrInt: nil, PtrInt2: new(int), Int: 1}).Info("wakaka")
}
