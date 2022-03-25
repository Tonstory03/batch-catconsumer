package config

import (
	"errors"
)

var mapService = make(map[string]Service, 0)

func loadingServiceConfig() {
	for _, serviceConfig := range GetServiceConfig() {
		mapService[serviceConfig.Name] = serviceConfig
	}
}

func GetService(name string) (Service, error) {
	if service, ok := mapService[name]; ok {
		return service, nil
	}
	return Service{}, errors.New("Service name not match in any case.")
}
