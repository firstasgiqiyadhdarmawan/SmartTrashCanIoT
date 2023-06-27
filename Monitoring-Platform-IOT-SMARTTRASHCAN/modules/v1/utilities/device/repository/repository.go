package repository

import (
	"ClearningPatternGO/modules/v1/utilities/device/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	ListDevice() ([]models.Device, error)
	GetLatestContent(access_key string) (models.ReceivedData, error)
	ExportSensorData(DeviceId string, input models.SensorData) error
	GetDeviceHistory() ([]models.CapacityHistory, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) ListDevice() ([]models.Device, error) {
	var device []models.Device
	err := r.db.Raw("SELECT * from device").Scan(&device).Error
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (n *repository) GetLatestContent(access_key string) (models.ReceivedData, error) {
	data := models.ReceivedData{}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://platform.antares.id:8443/~/antares-cse/antares-id/smartTrashCan/sensor-dalam/la", nil)
	req.Header.Set("X-M2M-Origin", access_key)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer resp.Body.Close()
	isiBody, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(isiBody, &data)
	return data, err
}

func (n *repository) ExportSensorData(DeviceId string, input models.SensorData) error {
	err := n.db.Exec("INSERT INTO capacity_history (device_id, capacity, status_device, date_updated) VALUES (?,?,?,?)", DeviceId, input.Kapasitas, input.Status_device, time.Now()).Error
	return err
}

func (r *repository) GetDeviceHistory() ([]models.CapacityHistory, error) {
	var capacityHistory []models.CapacityHistory
	err := r.db.Raw("SELECT capacity_history.capacity_history_id, device.device_id, device.device_name,  capacity_history.status_device, capacity_history.capacity, capacity_history.date_updated FROM capacity_history INNER JOIN device ON device.device_id = capacity_history.device_id ORDER BY capacity_history.date_updated DESC").Scan(&capacityHistory).Error
	if err != nil {
		return nil, err
	}
	return capacityHistory, nil
}
