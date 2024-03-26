package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/longhorn/go-common-libs/ns"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.TraceLevel)

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	_, err := ns.ReadFileContent("/proc/mounts") // Will be output by trace log.
	if err != nil {
		logrus.Error(err)
	}

	<-sigChannel
}
