package service

import (
	"encoding/json"
	"net/http"

	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
)

type UpsertRetryProcessInfoService Service

func NewUpsertRetryProcessInfoService(requestBody, correlateId, uuidVal string, requestHeader map[string]string) IService {
	serviceConfig, _ := config.GetService(SERVICE_UPSERT_RETRY_PROCESS_INFO)

	return &RemovePrepaidCatalogueService{
		Endpoint:          serviceConfig.Endpoint,
		HttpMethod:        HTTP_METHOD_PUT,
		RequestBody:       requestBody,
		Username:          serviceConfig.User,
		Password:          serviceConfig.Password,
		ReadTimeout:       serviceConfig.ReadTimeout,
		ConnectionTimeout: serviceConfig.ConnectionTimeout,
		UUID:              uuidVal,
		CorrelateId:       correlateId,
		TagsApp:           "UpsertRetryProcessInfo",
	}
}

func (s *UpsertRetryProcessInfoService) GetConfig() Service {
	return Service(*s)
}

func CallUpsertRetryProcessInfoService(requestBody RequestUpsertRetryProcessInfo, correlateId, uuidVal string, requestHeader map[string]string) ([]byte, *http.Response, error) {

	body, err := json.Marshal(requestBody)

	if err != nil {
		return nil, nil, err
	}

	iService := NewUpsertRetryProcessInfoService(string(body), correlateId, uuidVal, requestHeader)
	return DoService(iService)
}
