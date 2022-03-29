package service

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/locallogging"
)

type Service struct {
	Endpoint          string
	Headers           map[string]string
	HttpMethod        string
	RequestBody       string
	Username          string
	Password          string
	XChannel          string
	CorrelateId       string
	UUID              string
	TagsApp           string
	ReadTimeout       *int
	ConnectionTimeout *int
}

type IService interface {
	GetConfig() Service
}

func DoService(service IService) ([]byte, *http.Response, error) {

	serviceConfig := service.GetConfig()

	var correlateId string = serviceConfig.CorrelateId
	var UUID string = serviceConfig.UUID
	var tagsApp string = serviceConfig.TagsApp
	var profile string = config.GetApplication().Profile
	var endpoint string = serviceConfig.Endpoint
	var dataBytes []byte
	var err error
	var response *http.Response
	var statusCode string = "-"
	var statusMessage = "-"
	var startTime time.Time = time.Now()

	// set log error.
	// defer func() {
	// 	if err != nil {

	// 	}
	// }()

	// set log legacy.
	defer func() {
		logLegacy := locallogging.LocalLoggingLegacy{}

		if response != nil {
			statusCode = strconv.Itoa(response.StatusCode)
			statusMessage = response.Status
		}

		logLegacy.SetLegacyInputLoggerStart(correlateId, UUID, profile, tagsApp, startTime)
		logLegacy.SetLegacyInputLoggerEnd(serviceConfig.RequestBody, string(dataBytes), statusCode, statusMessage, endpoint, time.Now())
		logLegacy.WriteLogLegacy()
	}()

	dataBytes, response, err = doService(service)

	return dataBytes, response, err
}

func doService(service IService) ([]byte, *http.Response, error) {

	serviceConfig := service.GetConfig()

	var req *http.Request
	var client http.Client = http.Client{}
	var XChannel string = serviceConfig.XChannel
	var httpMethod string = serviceConfig.HttpMethod
	var endpoint string = serviceConfig.Endpoint
	var errorRequest error

	if serviceConfig.ReadTimeout != nil {
		client.Timeout = time.Duration(*serviceConfig.ReadTimeout) * time.Millisecond
	}

	switch httpMethod {
	case HTTP_METHOD_GET:
		req, errorRequest = http.NewRequest(HTTP_METHOD_GET, endpoint, nil)
	case HTTP_METHOD_POST:
		reqBodyBuffer := bytes.NewBuffer([]byte(serviceConfig.RequestBody))
		req, errorRequest = http.NewRequest(HTTP_METHOD_POST, endpoint, reqBodyBuffer)
		req.Header.Set("Content-Type", "application/json")
	case HTTP_METHOD_PUT:
		reqBodyBuffer := bytes.NewBuffer([]byte(serviceConfig.RequestBody))
		req, errorRequest = http.NewRequest(HTTP_METHOD_PUT, endpoint, reqBodyBuffer)
		req.Header.Set("Content-Type", "application/json")
	case HTTP_METHOD_DELETE:
		reqBodyBuffer := bytes.NewBuffer([]byte(serviceConfig.RequestBody))
		req, errorRequest = http.NewRequest(HTTP_METHOD_DELETE, endpoint, reqBodyBuffer)
		req.Header.Set("Content-Type", "application/json")
	}

	if errorRequest != nil {
		return nil, nil, errors.New(errorRequest.Error())
	}

	// set basic auth
	if serviceConfig.Username != "" && serviceConfig.Password != "" {
		req.SetBasicAuth(serviceConfig.Username, serviceConfig.Password)
	}

	// set headers
	if serviceConfig.Headers != nil {
		for k, v := range serviceConfig.Headers {
			req.Header.Set(k, v)
		}
	}

	req.Header.Add("Accept", `application/json`)
	req.Header.Set("X-Channel", XChannel)

	response, errorResponse := client.Do(req)

	defer func() {
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}()

	if errorResponse != nil {
		return nil, nil, errors.New(errorResponse.Error())
	}

	dataBytes, errRead := ioutil.ReadAll(response.Body)

	if errRead != nil {
		return nil, nil, errors.New(errRead.Error())
	}

	return dataBytes, response, nil
}
