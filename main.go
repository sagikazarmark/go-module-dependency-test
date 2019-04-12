package main

import (
	"github.com/goph/logur/adapters/logrusadapter"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrusadapter.New(logrus.New())

	logger.Info("hello")
}
