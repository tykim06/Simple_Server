package controllers

import (
	"crypto/tls"
	"encoding/json"
	"ilo/app/models"
	"net/http"
	"strconv"
	"time"
)

const timeout = time.Duration(10 * time.Second)

const (
	HTTPS                     string = "https://"
	POWER_STATE_URL                  = "/rest/v1/Chassis/1/PowerMetrics"
	FAN_TEMPERATURE_STATE_URL        = "/rest/v1/Chassis/1/ThermalMetrics"
	EVENT_LOG_STATE_URL              = "/rest/v1/Managers/1/Logs/IEL/Entries?page="
	SYSTEM_STATE_URL                 = "/rest/v1/Systems/1"
	SYSTEM_LOG_STATE_URL             = "/rest/v1/Systems/1/Logs/IML/Entries?page="
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
	case *models.EventLogJson:
		url += (EVENT_LOG_STATE_URL + strconv.Itoa(target.(*models.EventLogJson).Page))
	case *models.System:
		url += SYSTEM_STATE_URL
	case *models.SystemLogJson:
		url += (SYSTEM_LOG_STATE_URL + strconv.Itoa(target.(*models.SystemLogJson).Page))
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
