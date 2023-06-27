package models

import "time"

type Device struct {
	Device_id    string
	Device_name  string
	Date_created string
}

type CapacityHistory struct {
	Capacity_history_id int
	Device_id           string
	Capacity            int
	Device_name         string
	Status_device       string
	Date_updated        time.Time
	Date_formatter      string
}
