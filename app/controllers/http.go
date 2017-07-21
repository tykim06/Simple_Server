package controllers

import (
	"crypto/tls"
	"encoding/json"
	"ilo/app/models"
	"net/http"
	"time"
)

const timeout = time.Duration(10 * time.Second)

const (
	HTTPS                     string = "https://"
	POWER_STATE_URL                  = "/rest/v1/Chassis/1/PowerMetrics"
	FAN_TEMPERATURE_STATE_URL        = "/rest/v1/Chassis/1/ThermalMetrics"
	SYSTEM_STATE_URL                 = "/rest/v1/Systems/1"
)

func HttpGetState(ilo models.Ilo, target interface{}) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: timeout}

	url := HTTPS + ilo.Host
	switch target.(type) {
	case *models.FanTemperatureJson:
		url += FAN_TEMPERATURE_STATE_URL
	case *models.PowerJson:
		url += POWER_STATE_URL
	case *models.System:
		url += SYSTEM_STATE_URL
	}

	var (
		req *http.Request
		res *http.Response
		err error
	)

	if req, err = http.NewRequest("GET", url, nil); err != nil {
		return err
	}
	req.SetBasicAuth(ilo.User, ilo.Pass)
	if res, err = client.Do(req); err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
