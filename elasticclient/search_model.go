package elasticclient

import (
	"encoding/json"
	"fmt"
)

type ResultSearch struct {
	Took     int        `json:"took"`
	TimedOut bool       `json:"timed_out"`
	Hits     ResultHits `json:"hits"`
}

type ResultHits struct {
	Total TotalHits  `json:"total"`
	Hits  []DataHits `json:"hits"`
}

type TotalHits struct {
	Value    int    `json:"value"`
	Relation string `json:"Relation"`
}

type DataHits struct {
	Source DataSource `json:"_source"`
}

type DataSource struct {
	Action         string `json:"Action"`
	KafkaTimestamp int64  `json:"kafkaTimestamp"`
	Message        string `json:"message"`
	IsRetryMessage bool   `json:"isRetryMessage"`
	// Log            string `json:"log"`
}

type SearchRequest struct {
	Sort  map[string]string      `json:"sort,omitempty"`
	From  *int                   `json:"from,omitempty"`
	Size  *int                   `json:"size,omitempty"`
	Query map[string]interface{} `json:"query,omitempty"`
}

func (e *SearchRequest) Convert2Map() (map[string]interface{}, error) {
	var result map[string]interface{}

	data, err := json.Marshal(e)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &result)

	return result, nil
}

type KV struct {
	Key   string
	Value *string
}

func (v *KV) Convert2String() *string {
	var result *string

	if v.Value != nil {
		s := fmt.Sprintf(`%s:%s`, v.Key, *v.Value)
		result = &s
	}

	return result
}
