package service

import (
	"IProtect/api"
	"IProtect/model"
	"encoding/json"
	"errors"
	"os"
	"time"
)

const dataFilePath = "./data/blocked_ip.json"

type DataServiceFactory struct {
	api.DataService
}

type DataService struct {
	logger api.LoggerService
	data   []model.BlockedIp
}

func (d DataServiceFactory) NewDataService(logger api.LoggerService) (api.DataService, error) {
	_, err := os.Stat("./data/blocked_ip.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = createDataFile(logger)
			if err != nil {
				logger.Error("Unable to create file blocked_ip.json : " + err.Error())
				return nil, err
			}
		}

		logger.Error("Unable to get stats from file blocked_ip.json : " + err.Error())
		return nil, err
	}

	return &DataService{logger: logger}, nil
}

func (d *DataService) InitData() error {
	bytes, err := os.ReadFile(dataFilePath)
	if err != nil {
		d.logger.Error("Unable to read file blocked_ip.json : " + err.Error())
		return err
	}
	err = json.Unmarshal(bytes, &d.data)
	if err != nil {
		d.logger.Error("Unable to unmarshal json data from file blocked_ip.json : " + err.Error())
		return err
	}

	return nil
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

func createDataFile(logger api.LoggerService) error {
	err := os.MkdirAll("./data/", 0777)
	if err != nil {
		logger.Error("Unable create parents directory for blocked_ip.json : " + err.Error())
		return err
	}

	file, err := os.Create("./data/blocked_ip.json")
	if err != nil {
		logger.Error("Unable to create file blocked_ip.json : " + err.Error())
		return err
	}

	defaultBytes := []byte("[]")
	err = os.WriteFile(file.Name(), defaultBytes, 0777)
	if err != nil {
		logger.Error("Unable to initialize blocked_ip.json : " + err.Error())
		return err
	}
	err = file.Close()
	if err != nil {
		logger.Error("Unable to close file blocked_ip.json : " + err.Error())
		return err
	}
	return nil
}
