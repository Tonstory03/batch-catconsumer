package service

import (
	"encoding/json"
	"net/http"

	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
)

type GetRetryProcessInfoService Service

func NewGetRetryProcessInfoService(correlateId string, uuidVal string, requestHeader map[string]string) IService {
	serviceConfig, _ := config.GetService(SERVICE_GET_RETRY_PROCESS_INFO)

	return &RemovePrepaidCatalogueService{
		Endpoint:          serviceConfig.Endpoint,
		HttpMethod:        HTTP_METHOD_GET,
		Username:          serviceConfig.User,
		Password:          serviceConfig.Password,
		ReadTimeout:       serviceConfig.ReadTimeout,
		ConnectionTimeout: serviceConfig.ConnectionTimeout,
		UUID:              uuidVal,
		CorrelateId:       correlateId,
		TagsApp:           "GetRetryProcessInfo",
	}
}

func (s *GetRetryProcessInfoService) GetConfig() Service {
	return Service(*s)
}

func CallGetRetryProcessInfoService(correlateId string, uuidVal string, requestHeader map[string]string) (*ResponseGetRetryProcessInfo, *http.Response, error) {
	iService := NewGetRetryProcessInfoService(correlateId, uuidVal, requestHeader)

	var result ResponseGetRetryProcessInfo

	responseBody, httpResponse, err := DoService(iService)

	errUnmarshal := json.Unmarshal(responseBody, &result)

	if errUnmarshal != nil {
		return nil, httpResponse, err
	}

	return &result, httpResponse, err
}
