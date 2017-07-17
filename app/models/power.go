package models

import (
	"fmt"
)

type Power struct {
	Id     int64  `db:"Id"`
	Power1 string `db:"Power1"`
	Power2 string `db:"Power2"`
}

func (p Power) String() string {
	return fmt.Sprintf("Powers(%d, %s, %s)", p.Id, p.Power1, p.Power2)
}

type PowerJson struct {
	PowerSupplies []struct {
		FirmwareVersion      string `json:"FirmwareVersion"`
		LastPowerOutputWatts int    `json:"LastPowerOutputWatts"`
		LineInputVoltage     int    `json:"LineInputVoltage"`
		LineInputVoltageType string `json:"LineInputVoltageType"`
		Model                string `json:"Model"`
		Name                 string `json:"Name"`
		Oem                  struct {
			Hp struct {
				AveragePowerOutputWatts int  `json:"AveragePowerOutputWatts"`
				BayNumber               int  `json:"BayNumber"`
				HotplugCapable          bool `json:"HotplugCapable"`
				MaxPowerOutputWatts     int  `json:"MaxPowerOutputWatts"`
				Mismatched              bool `json:"Mismatched"`
				PowerSupplyStatus       struct {
					State string `json:"State"`
				} `json:"PowerSupplyStatus"`
				Type        string `json:"Type"`
				IPDUCapable bool   `json:"iPDUCapable"`
			} `json:"Hp"`
		} `json:"Oem"`
		PowerCapacityWatts int    `json:"PowerCapacityWatts"`
		PowerSupplyType    string `json:"PowerSupplyType"`
		SerialNumber       string `json:"SerialNumber"`
		SparePartNumber    string `json:"SparePartNumber"`
		Status             struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"PowerSupplies"`
}
