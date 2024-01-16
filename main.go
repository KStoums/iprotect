package main

import (
	"IProtect/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	loggerFactory := service.ZapLoggerFactory{}
	logger := loggerFactory.NewLogger()

	dataServiceFactory := service.DataServiceFactory{}
	dataService, err := dataServiceFactory.NewDataService(logger)
	if err != nil {
		logger.Fatal("Unable to initialize data service : " + err.Error())
		return
	}

	requestManagerListenerFactory := service.RequestManagerListenerFactory{}
	requestManagerListener := requestManagerListenerFactory.NewRequestManagerListener(logger, dataService)
	err = requestManagerListener.Start()
	if err != nil {
		logger.Fatal("Unable to start listeners : " + err.Error())
		return
	}

	errChan := make(chan error)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill, syscall.SIGQUIT, syscall.SIGTERM)

	logger.Info("IProtect is ready!")
	go func() {
		<-signalChan
		close(errChan)
	}()

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	err = <-errChan
	if err != nil {
		logger.Errorw(err.Error(), "hostname", hostname)
	}

	logger.Infow("Shutting down...", "hostname", hostname)
}
