package elasticclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"
)

func Search(es *elasticsearch.Client, body map[string]interface{}) (*ResultSearch, error) {

	var result ResultSearch
	var buffer bytes.Buffer

	json.NewEncoder(&buffer).Encode(body)
	response, err := es.Search(es.Search.WithBody(&buffer))

	if err != nil {
		return nil, err
	}

	json.NewDecoder(response.Body).Decode(&result)

	return &result, nil
}

// func getSearchBodyProcessFailure(env, startTime, endTime string) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, nil)
// }

// func getSearchBodyProcessFailurePaging(env, startTime, endTime string, from, size int) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, &SearchPaging{From: &from, Size: &size})
// }

// func GetSearchBodyProcessFailure(env, startTime, endTime string, from, size int) (map[string]interface{}, error) {

// 	return getSearchBodyRetryProcess(env, startTime, endTime, SearchRequest{From: &from, Size: &size})
// }

// func getSearchBodyProcessFailure(env, startTime, endTime string, result SearchRequest) (map[string]interface{}, error) {
// 	var query interface{} = fmt.Sprintf(QUERY_BODY_PROCESS_FAILURE, getTopic(env), startTime, endTime)
// 	var sort interface{} = `{"kafkaTimestamp": "asc"}`
// 	var from interface{} = 0
// 	var size interface{} = 9999
// 	qs := getQuery(&query, &sort, &from, &size)

// }

func getQuery(query, sort, from, size *interface{}) string {

	listKV := []KV{
		KV{Key: "query", Value: utils.ToJsonText(query)},
		KV{Key: "sort", Value: utils.ToJsonText(sort)},
		KV{Key: "from", Value: utils.ToJsonText(from)},
		KV{Key: "size", Value: utils.ToJsonText(size)},
	}

	return fmt.Sprintf(`{%s}`, convertKV2String(listKV))
}

func convertKV2String(keyValues []KV) string {

	var arr []string = make([]string, 0)

	for _, keyValue := range keyValues {
		if v := keyValue.Convert2String(); v != nil {
			arr = append(arr, *v)
		}
	}

	return strings.Join(arr, ",")
}
