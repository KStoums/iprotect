package service

import (
	"IProtect/api"
	"IProtect/model"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type DataServiceFactory struct {
	api.DataService
}

type DataService struct {
	logger api.LoggerService
	data   []model.BlockedIp
}

func (d DataServiceFactory) NewDataService(logger api.LoggerService) api.DataService {
	bytes, err := os.ReadFile("./data/blocked_ip.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll("./data/blocked_ip.json", 777)
			if err != nil {
				logger.Error("Unable create file blocked_ip.json : " + err.Error())
				return nil
			}
		}
		logger.Error("Unable to read file blocked_ip.json : " + err.Error())
		return nil
	}

	err = json.Unmarshal(bytes, d.DataService)
	if err != nil {
		logger.Error("Unable to unmarshal json data : " + err.Error())
		return nil
	}

	return &DataService{logger: logger}
}

func (d *DataService) GetAddressState(address string) bool {
	for _, addr := range d.data {
		return addr.Ip == address
	}

	return false
}

func (d *DataService) AddBlockedAddress(address string) {
	d.logger.Info("Adding " + address + "to blocked address data...")

	for _, addr := range d.data {
		if addr.Ip == address {
			d.logger.Info("Address " + address + " has already blocked!")
			return
		}
	}

	d.data = append(d.data, model.BlockedIp{
		Ip:        address,
		BlockedAt: time.Now(),
	})
}

func (d *DataService) RemoveBlockAddress(address string) {
	d.logger.Info("Deleting " + address + " to blocked address data...")

	for i, addr := range d.data {
		if addr.Ip == address {
			d.data = append(d.data[:i], d.data[i+1:]...)
			d.logger.Info("Address " + address + "has removed to blocked address data!")
			return
		}
	}

	d.logger.Info("Address " + address + " are not blocked!")
}
