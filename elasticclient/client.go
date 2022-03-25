package elasticclient

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"
)

func NewClient() (*elasticsearch.Client, error) {
	elasticConfig := config.GetElasticConfig()

	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticConfig.Endpoint,
		},
		Header: http.Header{},
	}

	if elasticConfig.EnableAuth {
		cfg.Header["Authorization"] = []string{utils.GetBasicAuth(elasticConfig.Username, elasticConfig.Password)}
	}

	return elasticsearch.NewClient(cfg)
}
