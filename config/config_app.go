package config

import (
	"os"

	"github.com/spf13/viper"
)

var AppConfig Config

const CURRENT_FILE = "configcmd"

func LoadingConfig() {

	viper.SetConfigName("/configmap/application")

	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		getDefaultConfigYML()
		return
	}

	err := viper.Unmarshal(&AppConfig)

	if err != nil {
		getDefaultConfigYML()
	}

}

func getDefaultConfigYML() {

	env := os.Getenv("ENV")

	configName := getConfigName(env)

	viper.SetConfigName(configName)

	viper.AddConfigPath(".")

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err := viper.Unmarshal(&AppConfig)

	if err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return AppConfig
}

func GetApplication() Application {
	return GetConfig().Application
}

func GetServer() Server {
	return GetConfig().Server
}

func GetCronJobs() []Cronjob {
	return GetConfig().Cronjob
}

func GetServiceConfig() []Service {
	return AppConfig.Service
}

func GetElasticConfig() Elastic {
	return AppConfig.Elastic
}

func getConfigName(env string) string {

	configName := "./application"

	switch env {
	case "local":
		configName += ".local"
	}
	return configName
}
