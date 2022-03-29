package service

import "th.truecorp.it.dsm.batch/batch-catconsumer/config"

type RemovePrepaidCatalogueService Service

func NewRemovePrepaidCatalogueService(requestBody string, correlateId string, uuidVal string, requestHeader map[string]string) IService {
	serviceConfig, _ := config.GetService(SERVICE_REMOVE_PREPAID_CATALOGUE)

	return &RemovePrepaidCatalogueService{
		Endpoint:          serviceConfig.Endpoint,
		HttpMethod:        HTTP_METHOD_DELETE,
		RequestBody:       requestBody,
		Username:          serviceConfig.User,
		Password:          serviceConfig.Password,
		ReadTimeout:       serviceConfig.ReadTimeout,
		ConnectionTimeout: serviceConfig.ConnectionTimeout,
		UUID:              uuidVal,
		CorrelateId:       correlateId,
		TagsApp:           "RemovePrepaidCatalogue",
	}
}

func (s *RemovePrepaidCatalogueService) GetConfig() Service {
	return Service(*s)
}
