package controllers

import (
	"crypto/tls"
	"encoding/json"
	"ilo/app/models"
	"net/http"
	"strconv"
	"time"
)

var client *http.Client

const (
	HTTPS                 string = "https://"
	POWER_STATE_URL              = "/rest/v1/Chassis/1/PowerMetrics"
	TEMPERATURE_STATE_URL        = "/rest/v1/Chassis/1/ThermalMetrics"
	FAN_STATE_URL                = "/rest/v1/Chassis/1/ThermalMetrics"
	EVENT_LOG_STATE_URL          = "/rest/v1/Managers/1/Logs/IEL/Entries?page="
	SYSTEM_STATE_URL             = "/rest/v1/Systems/1"
	SYSTEM_LOG_STATE_URL         = "/rest/v1/Systems/1/Logs/IML/Entries?page="
)

type GET_STATE_TYPE uint8

const (
	GET_STATE_FAN GET_STATE_TYPE = iota
	GET_STATE_POWER
	GET_STATE_TEMPERATURE
	GET_STATE_EVENT_LOG
	GET_STATE_SYSTEM
	GET_STATE_SYSTEM_LOG
)

func HttpGetState(state_type GET_STATE_TYPE, ilo models.Ilo, target interface{}) error {
	url := HTTPS + ilo.Host
	switch state_type {
	case GET_STATE_FAN:
		url += FAN_STATE_URL
	case GET_STATE_TEMPERATURE:
		url += TEMPERATURE_STATE_URL
	case GET_STATE_POWER:
		url += POWER_STATE_URL
	case GET_STATE_EVENT_LOG:
		url += (EVENT_LOG_STATE_URL + strconv.Itoa(target.(*models.EventLogJson).Page))
	case GET_STATE_SYSTEM:
		url += SYSTEM_STATE_URL
	case GET_STATE_SYSTEM_LOG:
		url += (SYSTEM_LOG_STATE_URL + strconv.Itoa(target.(*models.SystemLogJson).Page))
	}

	var req *http.Request
	var res *http.Response
	var err error

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

func InitHttp() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	timeout := time.Duration(3 * time.Second)
	client = &http.Client{Transport: tr, Timeout: timeout}
}
