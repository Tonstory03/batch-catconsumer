package elasticclient

import (
	"fmt"
	"strings"

	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"
)

func getSearchBodyRetryProcess(env, startTime, endTime string, result SearchRequest) (map[string]interface{}, error) {

	// sort by kafka timestamp
	sortMap := map[string]string{
		"kafkaTimestamp": "asc",
	}

	mustMap := []map[string]interface{}{

		// query eq topic.
		map[string]interface{}{
			"term": map[string]string{
				"topic.keyword": getTopic(env),
			},
		},

		// query range timestamp.
		map[string]interface{}{
			"range": map[string]interface{}{
				"timestamp": map[string]string{
					"gte": startTime,
					"lt":  endTime,
				},
			},
		},

		// query exist field action.
		map[string]interface{}{
			"exists": map[string]string{
				"field": "action",
			},
		},
	}

	mustNotMap := []map[string]interface{}{
		map[string]interface{}{
			"term": map[string]string{
				"action.keyword": "fetchAll",
			},
		},
	}

	boolMap := map[string]interface{}{
		"must":     mustMap,
		"must_not": mustNotMap,
	}

	result.Query = map[string]interface{}{
		"bool": boolMap,
	}

	result.Sort = sortMap

	return result.Convert2Map()
}

func GetSearchBodyRetryProcess(env, startTime, endTime string) (map[string]interface{}, error) {

	return getSearchBodyRetryProcess(env, startTime, endTime, SearchRequest{})
}

func GetSearchBodyRetryProcessPaging(env, startTime, endTime string, from, size int) (map[string]interface{}, error) {

	return getSearchBodyRetryProcess(env, startTime, endTime, SearchRequest{From: &from, Size: &size})
}

func getTopic(env string) string {
	return fmt.Sprintf("%s-cat-offer", env)
}

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
