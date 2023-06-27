package service

import (
	"ClearningPatternGO/modules/v1/utilities/device/models"
	"ClearningPatternGO/modules/v1/utilities/device/repository"
	"encoding/json"
	"fmt"
)

type Service interface {
	ListDevice() ([]models.Device, error)
	GetLatestContent(access_key string) (models.ReceivedData, error)
	GetDatafromContent(input string, DeviceId string) (models.SensorData, error)
	GetDeviceHistory() ([]models.CapacityHistory, error)
	GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.SensorData, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) ListDevice() ([]models.Device, error) {
	allDevice, err := s.repository.ListDevice()
	if err != nil {
		return nil, err
	}
	return allDevice, nil
}

func (n *service) GetLatestContent(access_key string) (models.ReceivedData, error) {
	getLatestData, err := n.repository.GetLatestContent(access_key)
	return getLatestData, err
}

func (n *service) GetDatafromContent(input string, DeviceId string) (models.SensorData, error) {
	var data models.SensorData
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	err = n.repository.ExportSensorData(DeviceId, data)

	return data, nil
}

func (s *service) GetDeviceHistory() ([]models.CapacityHistory, error) {
	capacityHistory, err := s.repository.GetDeviceHistory()
	for i, v := range capacityHistory {
		capacityHistory[i].Date_formatter = v.Date_updated.Format("15:04:05")
		if capacityHistory[i].Capacity > 100 {
			capacityHistory[i].Capacity = 100
		}
	}
	if err != nil {
		return nil, err
	}
	return capacityHistory, nil
}

func (n *service) GetDatafromWebhook(sensorData string, antaresDeviceID string) (models.SensorData, error) {
	var data models.SensorData
	err := json.Unmarshal([]byte(sensorData), &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	fmt.Println("Data sensor terbaru :", data)
	err = n.repository.ExportSensorData(antaresDeviceID, data)
	return data, err
}
