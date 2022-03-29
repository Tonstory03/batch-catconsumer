package locallogging

import (
	"encoding/json"
	"fmt"
	"time"

	"th.truecorp.it.dsm.batch/batch-catconsumer/intutilities"
)

type LocalLogging struct {
	uuid             string
	username         string
	legacyUsername   string
	inputParam       string
	channel          string
	correlateId      string
	ip               string
	tagsApp          string
	tagsEnv          string
	gatewayType      string
	errorCode        string
	message          string
	system           string
	resultStatus     string
	mainInputKey     string
	mainInputValue   string
	outputString     string
	startTime        time.Time
	endTime          time.Time
	stackTrace       string
	errorApplication string
	errorModule      string
	errorFile        string
	errorFunction    string
}

func (logging *LocalLogging) SetErrorInputLoggerBeforeController(errLogger *ErrorInputLogger) {
	// Base Input Logger
	logging.tagsEnv = errLogger.TagsEnv
	logging.tagsApp = errLogger.TagsApp

	// Response Input Logger
	logging.errorCode = errLogger.ErrorCode
	logging.message = errLogger.Message
	logging.errorApplication = errLogger.ErrorApplication
	logging.errorModule = errLogger.ErrorModule
	logging.errorFile = errLogger.ErrorFile
	logging.errorFunction = errLogger.ErrorFunction
}

func (logging *LocalLogging) SetErrorInputLogger(errCode string, msg string, err error, application string, module string, file string, function string) {

	// Response Input Logger
	logging.errorCode = errCode
	logging.message = msg
	logging.errorApplication = application
	logging.errorModule = module
	logging.errorFile = file
	logging.errorFunction = function

	// Convert error to string stacktrace
	logging.stackTrace = Wrap(err).StackTrace
}

func (logging *LocalLogging) WriteLogError() {
	logError := ErrorInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_ERROR,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Response Input Logger
		ErrorCode:        logging.errorCode,
		Message:          logging.message,
		StackTrace:       logging.stackTrace,
		ErrorApplication: logging.errorApplication,
		ErrorModule:      logging.errorModule,
		ErrorFile:        logging.errorFile,
		ErrorFunction:    logging.errorFunction,
		ResultStatus:     "F",
	}

	// Set tags
	logError.Tags[0] = logging.tagsEnv
	logError.Tags[1] = logging.tagsApp
	logError.Tags[2] = TAG_ERROR

	jsErr, err := json.Marshal(logError)

	if err != nil {

		// executes a function asynchronously
		go fmt.Println(err)
		return
	}

	fmt.Println(string(jsErr))
}

func (logging *LocalLoggingLegacy) SetLegacyInputLoggerStart(correlateId string, uuid string, tagsEnv string, tagsApp string, tStartTime time.Time) {
	logging.correlateId = correlateId
	logging.uuid = uuid
	logging.timestamp = intutilities.GetCurrentISO8601()
	logging.tags = []string{tagsEnv, tagsApp, TAG_LEGACY}
	logging.startRequest = intutilities.GetCurrentISO8601()
	logging.startTime = tStartTime

}
func (logging *LocalLoggingLegacy) SetLegacyInputLoggerEnd(request string, response string, httpResponseCode string, httpResponseMessage string, targetEp string, tEndTime time.Time) {
	logging.request = request
	logging.response = response
	logging.httpResponseCode = httpResponseCode
	logging.httpResponseMessage = httpResponseMessage
	logging.targetEp = targetEp
	logging.endRequest = intutilities.GetCurrentISO8601()
	logging.endTime = tEndTime
}

func (logging *LocalLoggingLegacy) WriteLogLegacy() {

	start := logging.startTime.UnixNano() / int64(time.Millisecond)
	end := logging.endTime.UnixNano() / int64(time.Millisecond)
	diff := end - start

	logRes := LegacyInputLogger{
		CorrelateId:    logging.correlateId,
		Uuid:           logging.uuid,
		DateCompletion: intutilities.GetCurrentISO8601(),
		LogType:        LOGTYPE_APPLICATION,
		Timestamp:      intutilities.GetCurrentISO8601(),

		// Legacy Input Logger
		Request:             logging.request,
		Response:            logging.response,
		HttpResponseCode:    logging.httpResponseCode,
		HttpResponseMessage: logging.httpResponseMessage,
		StartRequest:        logging.startRequest,
		EndRequest:          logging.endRequest,
		TargetEp:            logging.targetEp,
		ElapsedTime:         diff,
	}

	// Set tags
	logRes.Tags[0] = logging.tags[0]
	logRes.Tags[1] = logging.tags[1]
	logRes.Tags[2] = logging.tags[2]

	jsRes, err := json.Marshal(logRes)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsRes))

}

type LocalLoggingLegacy struct {
	uuid        string
	correlateId string
	timestamp   string
	tags        []string
	// dateCompletion      string
	// logType             int
	httpResponseCode    string
	httpResponseMessage string
	request             string
	response            string
	targetEp            string
	// elapsedTime         string
	startRequest string
	endRequest   string
	startTime    time.Time
	endTime      time.Time
}
