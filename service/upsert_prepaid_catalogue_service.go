package service

import (
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
)

type UpsertPrepaidCatalogueService Service

func NewUpsertPrepaidCatalogueService(requestBody string, correlateId string, uuidVal string, requestHeader map[string]string) IService {
	serviceConfig, _ := config.GetService(SERVICE_UPSERT_PREPAID_CATALOGUE)

	return &UpsertPrepaidCatalogueService{
		Endpoint:          serviceConfig.Endpoint,
		HttpMethod:        HTTP_METHOD_PUT,
		RequestBody:       requestBody,
		Username:          serviceConfig.User,
		Password:          serviceConfig.Password,
		ReadTimeout:       serviceConfig.ReadTimeout,
		ConnectionTimeout: serviceConfig.ConnectionTimeout,
		UUID:              uuidVal,
		CorrelateId:       correlateId,
		TagsApp:           "UpsertPrepaidCatalogue",
	}
}

func (s *UpsertPrepaidCatalogueService) GetConfig() Service {
	return Service(*s)
}
