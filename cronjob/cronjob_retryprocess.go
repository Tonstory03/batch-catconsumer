package cronjob

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/google/uuid"
	"th.truecorp.it.dsm.batch/batch-catconsumer/config"
	"th.truecorp.it.dsm.batch/batch-catconsumer/elasticclient"
	"th.truecorp.it.dsm.batch/batch-catconsumer/service"
	"th.truecorp.it.dsm.batch/batch-catconsumer/utils"

	"th.truecorp.it.dsm.batch/batch-catconsumer/errorutils"
)

func retryProcessKafka() {

	es, err := elasticclient.NewClient()

	if err != nil {
		fmt.Println("Error NewClient:", err.Error())
	}

	profile := config.GetApplication().Profile

	resultQueryFailure, err := searchProcessFailure(es, profile)

	if err != nil {
		panic(err)
	}

	if resultQueryFailure.Hits.Total.Value == 0 {
		// fmt.Println("No process failure ", resultQueryFailure.Hits.Total.Value)
		return
	}

	for _, dataHits := range resultQueryFailure.Hits.Hits {

		var (
			offerCode      = dataHits.Source.OfferCode
			offerName      = dataHits.Source.OfferName
			kafkaTimestamp = dataHits.Source.KafkaTimestamp
			action         = dataHits.Source.Action
			message        = dataHits.Source.Message
			uuidVal        = uuid.New().String()
		)

		if utils.IsEmptyString(offerCode) || utils.IsEmptyString(offerName) {
			continue
		}

		// fmt.Println("Start process ", *offerCode, *offerName, action)

		resultQueryByCodeName, err := searchProcessByCodeName(es, profile, *offerCode, *offerName, kafkaTimestamp)

		if err != nil {
			panic(err)
		}

		var caseRetryProcess bool = true
		// checking total result.
		if resultQueryByCodeName.Hits.Total.Value > 0 {
			caseRetryProcess, err = checkCaseRetryProcess(action, message, resultQueryByCodeName.Hits.Hits)
			if err != nil {
				fmt.Println("Error ", err.Error())
				// continue
			}
		}

		if caseRetryProcess {
			// call retry process
			// fmt.Println("code:", *offerCode, ", name: ", *offerName, " is going to retry process.")

			_, err := callService(message, action, "", uuidVal, kafkaTimestamp)

			if err != nil {
				fmt.Println("Error call service", err)
			}

		} else {
			// fmt.Println("code:", *offerCode, ", name: ", *offerName, " won't retry process.")
		}
	}

}

func callService(msgValue, action, correlatedId, uuidVal string, kafkaTimestamp int64) (bool, error) {

	var serviceName string

	switch a := action; {
	case strings.EqualFold(a, HEADER_OFFER_ACTION_VERSION_EXP):
		serviceName = service.SERVICE_UPDATE_PREPAID_CATALOGUE

	case strings.EqualFold(a, HEADER_OFFER_ACTION_UPSERT):
		serviceName = service.SERVICE_UPSERT_PREPAID_CATALOGUE

	case strings.EqualFold(a, HEADER_OFFER_ACTION_DELETE):
		serviceName = service.SERVICE_REMOVE_PREPAID_CATALOGUE

	case strings.EqualFold(a, HEADER_OFFER_ACTION_FETCHALL):
		serviceName = service.SERVICE_UPSERT_PREPAID_CATALOGUE
	default:
		return false, errors.New(fmt.Sprintf("Error header action: {%s} doesn't match in any case", action))
	}

	resultService, errGetService := getService(msgValue, serviceName, correlatedId, uuidVal, kafkaTimestamp)

	if errGetService != nil {
		return false, errors.New(fmt.Sprintf("Error getService: %s", errGetService))
	}

	return doService(resultService)
}

func doService(iService service.IService) (bool, error) {

	var resp service.ResponsePrepaidCatalogue

	respBytes, response, errorService := service.DoService(iService)

	if errorService != nil {
		return true, errors.New(fmt.Sprintf("Error DoService: %s", errorService))
	}

	errUnmarshal := json.Unmarshal(respBytes, &resp)

	if errUnmarshal != nil {
		return true, errors.New(fmt.Sprintf("Error Unmarshal body: %s", errUnmarshal))
	}

	if resp.ErrorCode == errorutils.ERR_CD_INTERNAL_FAILURE {
		return true, errors.New(fmt.Sprintf("Error response body errorCode: %s", errorutils.ERR_CD_INTERNAL_FAILURE))
	}

	if response.StatusCode == 400 || resp.ErrorCode == errorutils.ERR_CD_REQUIRED_FIELD {
		return false, errors.New(fmt.Sprintf("Error Bad request : %s", resp.Message))
	}

	return false, nil
}

func getService(msgValue, serviceName, correlatedId, uuidVal string, kafkaTimestamp int64) (service.IService, error) {

	var msgData map[string]interface{}

	var result service.IService

	err := json.Unmarshal([]byte(msgValue), &msgData)

	if err != nil {
		return nil, err
	}

	request := service.RequestPrepaidCatalogue{
		CorrelatedId:   correlatedId,
		CallerUuid:     uuidVal,
		Data:           msgData,
		KafkaTimestamp: kafkaTimestamp,
	}

	msgBytes, err := json.Marshal(request)

	if err != nil {
		return nil, err
	}

	switch serviceName {

	case service.SERVICE_UPSERT_PREPAID_CATALOGUE:
		result = service.NewUpsertPrepaidCatalogueService(string(msgBytes), correlatedId, uuidVal, nil)

	case service.SERVICE_UPDATE_PREPAID_CATALOGUE:
		result = service.NewUpdatePrepaidCatalogueService(string(msgBytes), correlatedId, uuidVal, nil)

	case service.SERVICE_REMOVE_PREPAID_CATALOGUE:
		result = service.NewRemovePrepaidCatalogueService(string(msgBytes), correlatedId, uuidVal, nil)

	default:
		return nil, errors.New(fmt.Sprintf("Error Service name: {%s} doesn't support.", serviceName))
	}

	return result, nil
}

func checkCaseRetryProcess(action, message string, dataHits []elasticclient.DataHits) (bool, error) {
	var checkCaseRetry func(dataSource elasticclient.DataSource, id string) bool

	var mapPrepaidCat map[string]interface{}

	err := json.Unmarshal([]byte(message), &mapPrepaidCat)

	if err != nil {
		return false, err
	}

	var result bool = true
	var id string

	if !strings.EqualFold(action, elasticclient.ACTION_DELETE) {
		id = mapPrepaidCat["id"].(string)
	}

	switch a := action; {
	case strings.EqualFold(a, elasticclient.ACTION_VERSION_EXP):
		checkCaseRetry = func(dataSource elasticclient.DataSource, id string) bool {
			var mapData map[string]interface{}

			err := json.Unmarshal([]byte(dataSource.Message), &mapData)

			if err != nil {
				return false
			}

			return strings.EqualFold(dataSource.Action, elasticclient.ACTION_VERSION_EXP) && mapData["id"].(string) != id

		}

	case strings.EqualFold(a, elasticclient.ACTION_UPSERT) || strings.EqualFold(a, elasticclient.ACTION_FETCHALL):
		checkCaseRetry = func(dataSource elasticclient.DataSource, id string) bool {
			var mapData map[string]interface{}

			err := json.Unmarshal([]byte(dataSource.Message), &mapData)

			if err != nil {
				return false
			}

			return strings.EqualFold(dataSource.Action, elasticclient.ACTION_VERSION_EXP) && mapData["id"].(string) != id
		}

	case strings.EqualFold(a, elasticclient.ACTION_DELETE):
		checkCaseRetry = func(dataSource elasticclient.DataSource, id string) bool {
			return strings.EqualFold(dataSource.Action, elasticclient.ACTION_VERSION_EXP)
		}

	default:
		return false, nil
	}

	for _, dataHit := range dataHits {

		dataSource := dataHit.Source

		if dataSource.IsRetryMessage {
			continue
		}

		if !checkCaseRetry(dataSource, id) {
			result = false
		}
	}

	return result, nil
}

func searchProcessFailure(es *elasticsearch.Client, env string) (*elasticclient.ResultSearch, error) {

	var (
		from           = utils.NewIntPointer(0)
		size           = utils.NewIntPointer(9999)
		lastCheckPoint = getLastCheckPoint()
		rangeTimestamp = getRangeTimestamp(lastCheckPoint)
	)

	// fmt.Println(rangeTimestamp)
	// // fix test
	// rangeTimestamp = elasticclient.SearchRangeTimestamp{StartTime: "2022-03-25T00:00:00Z", EndTime: "2022-03-25T14:22:29Z"}

	query, err := getSearchBodyProcessFailure(env, rangeTimestamp, &elasticclient.SearchPaging{From: from, Size: size})

	if err != nil {
		return nil, err
	}

	return elasticclient.Search(es, query)
}

func searchProcessByCodeName(es *elasticsearch.Client, env, offerCode, offerName string, kafkaTimestamp int64) (*elasticclient.ResultSearch, error) {

	var (
		from = utils.NewIntPointer(0)
		size = utils.NewIntPointer(9999)
	)

	query, err := getSearchBodyProcessByCodeName(env, offerCode, offerName, kafkaTimestamp, &elasticclient.SearchPaging{From: from, Size: size})

	if err != nil {
		return nil, err
	}

	return elasticclient.Search(es, query)
}

func getRangeTimestamp(lastCheckPoint time.Time) elasticclient.SearchRangeTimestamp {
	now := time.Now()

	return elasticclient.SearchRangeTimestamp{
		StartTime: utils.Time2StrFormatUTC(lastCheckPoint),
		EndTime:   utils.Time2StrFormatUTC(now),
	}
}

func getLastCheckPoint() time.Time {
	now := time.Now()

	minute := int(math.Floor(float64(now.Minute()/30)) * 30)

	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), minute, 0, 0, time.Local).Add(-time.Minute * 30)
}

func getTopic(env string) string {
	return fmt.Sprintf("%s-cat-offer", env)
}

func getSearchBodyProcessFailure(env string, searchRangeTimestamp elasticclient.SearchRangeTimestamp, searchPaging *elasticclient.SearchPaging) (map[string]interface{}, error) {

	var result elasticclient.SearchRequest = elasticclient.SearchRequest{}

	// Check and add from size.
	if searchPaging != nil && searchPaging.From != nil && searchPaging.Size != nil {
		result.From = searchPaging.From
		result.Size = searchPaging.Size
	}

	// Sort by kafka timestamp
	result.Sort = map[string]string{
		"kafkaTimestamp": "asc",
	}

	// Add query.
	var query string = fmt.Sprintf(elasticclient.QUERY_BODY_PROCESS_FAILURE, getTopic(env), searchRangeTimestamp.StartTime, searchRangeTimestamp.EndTime)
	var mapQuery map[string]interface{}

	err := json.Unmarshal([]byte(query), &mapQuery)

	if err != nil {
		return nil, err
	}

	result.Query = mapQuery

	return result.Convert2Map()
}

func getSearchBodyProcessByCodeName(env, offerCode, offerName string, kafkaTimestamp int64, searchPaging *elasticclient.SearchPaging) (map[string]interface{}, error) {

	var result elasticclient.SearchRequest = elasticclient.SearchRequest{}

	// Check and add from size.
	if searchPaging != nil && searchPaging.From != nil && searchPaging.Size != nil {
		result.From = searchPaging.From
		result.Size = searchPaging.Size
	}

	// Sort by kafka timestamp
	result.Sort = map[string]string{
		"kafkaTimestamp": "asc",
	}

	// Add query.
	var query string = fmt.Sprintf(elasticclient.QUERY_BODY_BY_CODENAME, getTopic(env), strconv.FormatInt(kafkaTimestamp, 10), offerCode, offerName)
	var mapQuery map[string]interface{}

	err := json.Unmarshal([]byte(query), &mapQuery)

	if err != nil {
		return nil, err
	}

	result.Query = mapQuery

	return result.Convert2Map()
}
