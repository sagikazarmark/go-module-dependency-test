package main

import (
	logrusadapter "logur.dev/adapter/logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrusadapter.New(logrus.New())

	logger.Info("hello")
}
