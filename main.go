package main

import (
	"IProtect/service"
)

func main() {
	loggerFactory := service.ZapLoggerFactory{}
	logger := loggerFactory.NewLogger()

	dataServiceFactory := service.DataServiceFactory{}
	dataService := dataServiceFactory.NewDataService(logger)

	requestManagerListenerFactory := service.RequestManagerListenerFactory{}
	requestManagerListener := requestManagerListenerFactory.NewRequestManagerListener(logger, dataService)
	err := requestManagerListener.Start()
	if err != nil {
		logger.Fatal("Unable to start listeners : " + err.Error())
		return
	}
}
