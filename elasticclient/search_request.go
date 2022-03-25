package elasticclient

import "fmt"

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
