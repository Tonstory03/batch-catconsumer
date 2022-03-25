package cronjob

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/robfig/cron/v3"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/elasticclient"
	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"
)

func Init() {
	var startJob bool
	c := cron.New()
	cronJobs := config.GetCronJobs()

	for _, cronJob := range cronJobs {

		if !cronJob.Enable {
			continue
		}

		switch cronJob.Name {
		case NAME_RETRY_PROCESS_KAFKA:
			retryProcessKafka()
			// c.AddFunc(cronJob.Expression, retryProcessKafka)
			// startJob = true

		}

	}
	if startJob {
		c.Start()
	}
}

func retryProcessKafka() {

	es, err := elasticclient.NewClient()

	if err != nil {
		fmt.Println("Error NewClient:", err.Error())
	}

	var buffer bytes.Buffer

	profile := config.GetApplication().Profile

	lastCheckPoint := getLastCheckPoint()

	rangeTimestamp := getRangeTimestamp(lastCheckPoint)

	fmt.Println(rangeTimestamp)

	query, err := elasticclient.GetSearchBodyRetryProcess(profile, "2022-03-15T00:00:00Z", "2022-03-17T14:22:29Z")

	json.NewEncoder(&buffer).Encode(query)

	response, _ := es.Search(es.Search.WithBody(&buffer))

	var result elasticclient.ResultSearch

	json.NewDecoder(response.Body).Decode(&result)

	DataSources := make([]elasticclient.DataSource, 0)

	// mapIndexData := make(map[string]interface{})

	for _, dataHits := range result.Hits.Hits {
		DataSources = append(DataSources, dataHits.Source)
		// mapIndexData[dataHits.Source]
	}
	var mm map[string]interface{}

	json.Unmarshal([]byte(result.Hits.Hits[0].Source.Message), &mm)
	fmt.Println(mm)
	//fmt.Println(result.Hits.Hits[0].Source)
}

func getRangeTimestamp(lastCheckPoint time.Time) QueryTimestamp {
	now := time.Now()

	return QueryTimestamp{
		StartTime: utils.Time2StrFormatUTC(lastCheckPoint),
		EndTime:   utils.Time2StrFormatUTC(now),
	}
}

func getLastCheckPoint() time.Time {
	now := time.Now()

	minute := int(math.Floor(float64(now.Minute()/30)) * 30)

	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), minute, 0, 0, time.Local).Add(-time.Minute * 30)
}
