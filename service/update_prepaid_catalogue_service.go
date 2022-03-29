package service

import (
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
)

type UpdatePrepaidCatalogueService Service

func NewUpdatePrepaidCatalogueService(requestBody string, correlateId string, uuidVal string, requestHeader map[string]string) IService {
	serviceConfig, _ := config.GetService(SERVICE_UPDATE_PREPAID_CATALOGUE)

	return &UpdatePrepaidCatalogueService{
		Endpoint:          serviceConfig.Endpoint,
		HttpMethod:        HTTP_METHOD_PUT,
		RequestBody:       requestBody,
		Username:          serviceConfig.User,
		Password:          serviceConfig.Password,
		ReadTimeout:       serviceConfig.ReadTimeout,
		ConnectionTimeout: serviceConfig.ConnectionTimeout,
		UUID:              uuidVal,
		CorrelateId:       correlateId,
		TagsApp:           "UpdatePrepaidCatalogue",
	}
}

func (s *UpdatePrepaidCatalogueService) GetConfig() Service {
	return Service(*s)
}
