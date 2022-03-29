package locallogging

type RequestInputLogger struct {
	// Base Input Logger
	CorrelateId    string    `json:"correlateId" binding:"required"`
	TagsEnv        string    `json:"tagsEnv,omitempty"`
	TagsApp        string    `json:"tagsApp,omitempty"`
	TagsParameter  string    `json:"tagsParameter,omitempty"`
	Tags           [3]string `json:"tags" binding:"required"`
	Uuid           string    `json:"uuid" binding:"required"`
	DateCompletion string    `json:"dateCompletion" binding:"required"`
	LogType        int       `json:"logType" binding:"required"`
	Timestamp      string    `json:"@timestamp" binding:"required"`

	// Request Input Logger
	LegacyUsername   string `json:"legacyUsername" binding:"required"`
	Username         string `json:"username" binding:"required"`
	InputParamLog    string `json:"inputParamLog" binding:"required"`
	Host             string `json:"host" binding:"required"`
	Ip               string `json:"ip" binding:"required"`
	InputParamLogObj string `json:"inputParamLogObj,omitempty"`
	Channel          string `json:"channel" binding:"required"`
	GatewayType      string `json:"gatewayType" binding:"required"`
	StartTime        string `json:"startTime,omitempty"`
}

type ResponseInputLogger struct {
	// Base Input Logger
	CorrelateId    string    `json:"correlateId" binding:"required"`
	TagsEnv        string    `json:"tagsEnv,omitempty"`
	TagsApp        string    `json:"tagsApp,omitempty"`
	TagsParameter  string    `json:"tagsParameter,omitempty"`
	Tags           [3]string `json:"tags" binding:"required"`
	Uuid           string    `json:"uuid" binding:"required"`
	DateCompletion string    `json:"dateCompletion" binding:"required"`
	LogType        int       `json:"logType" binding:"required"`
	Timestamp      string    `json:"@timestamp" binding:"required"`

	// Response Input Logger
	Username     string `json:"username" binding:"required"`
	ErrorCode    string `json:"errorCode" binding:"required"`
	Message      string `json:"message" binding:"required"`
	ResultStatus string `json:"resultStatus" binding:"required"`
	OutputString string `json:"outputString" binding:"required"`

	EndTime string `json:"endTime,omitempty"`
}

type ElapsedLogger struct {
	// Base Input Logger
	CorrelateId    string    `json:"correlateId" binding:"required"`
	TagsEnv        string    `json:"tagsEnv,omitempty"`
	TagsApp        string    `json:"tagsApp,omitempty"`
	TagsParameter  string    `json:"tagsParameter,omitempty"`
	Tags           [3]string `json:"tags" binding:"required"`
	Uuid           string    `json:"uuid" binding:"required"`
	DateCompletion string    `json:"dateCompletion" binding:"required"`
	LogType        int       `json:"logType" binding:"required"`
	Timestamp      string    `json:"@timestamp" binding:"required"`

	// Response Input Logger
	Username       string `json:"username" binding:"required"`
	ErrorCode      string `json:"errorCode" binding:"required"`
	Message        string `json:"message" binding:"required"`
	System         string `json:"system" binding:"required"`
	MainInputKey   string `json:"mainInputKey" binding:"required"`
	MainInputValue string `json:"mainInputValue" binding:"required"`
	ElapsedTime    int64  `json:"elapsedTime" binding:"required"`
}

// ErrorInputLogger
type ErrorInputLogger struct {
	// Base Input Logger
	CorrelateId    string    `json:"correlateId" binding:"required"`
	TagsEnv        string    `json:"tagsEnv,omitempty"`
	TagsApp        string    `json:"tagsApp,omitempty"`
	TagsParameter  string    `json:"tagsParameter,omitempty"`
	Tags           [3]string `json:"tags" binding:"required"`
	Uuid           string    `json:"uuid" binding:"required"`
	DateCompletion string    `json:"dateCompletion" binding:"required"`
	LogType        int       `json:"logType" binding:"required"`
	Timestamp      string    `json:"@timestamp" binding:"required"`

	// Response Input Logger
	ErrorCode        string `json:"errorCode" binding:"required"`
	Message          string `json:"message" binding:"required"`
	ResultStatus     string `json:"resultStatus" binding:"required"`
	StackTrace       string `json:"StackTrace" binding:"required"`
	ErrorApplication string `json:"errorApplication" binding:"required"`
	ErrorModule      string `json:"errorModule" binding:"required"`
	ErrorFile        string `json:"errorFile" binding:"required"`
	ErrorFunction    string `json:"errorFunction" binding:"required"`
}

type LegacyInputLogger struct {
	// Base Input Logger
	CorrelateId    string    `json:"correlateId" binding:"required"`
	TagsEnv        string    `json:"tagsEnv,omitempty"`
	TagsApp        string    `json:"tagsApp,omitempty"`
	TagsParameter  string    `json:"tagsParameter,omitempty"`
	Tags           [3]string `json:"tags" binding:"required"`
	Uuid           string    `json:"uuid" binding:"required"`
	DateCompletion string    `json:"dateCompletion" binding:"required"`
	LogType        int       `json:"logType" binding:"required"`
	Timestamp      string    `json:"@timestamp" binding:"required"`

	// Legacy Input Logger
	Request             string `json:"request,omitempty"`
	Response            string `json:"response,omitempty"`
	HttpResponseCode    string `json:"httpResponseCode,omitempty"`
	HttpResponseMessage string `json:"httpResponseMessage,omitempty"`
	StartRequest        string `json:"startRequest,omitempty"`
	EndRequest          string `json:"endRequest,omitempty"`
	TargetEp            string `json:"targetEp,omitempty"`
	ElapsedTime         int64  `json:"elapsedTime,omitempty"`

	EndTime string `json:"endTime,omitempty"`
}

type kafkaReadMessageLogger struct {
	Action         string   `json:"action"`
	UUID           string   `json:"uuid"`
	CorrelatedId   string   `json:"correlatedId"`
	Timestamp      string   `json:"timestamp"`
	Tags           []string `json:"tags"`
	KafkaTimestamp int64    `json:"kafkaTimestamp"`
	Topic          string   `json:"topic"`
	Partition      string   `json:"partition"`
	Offset         string   `json:"offset"`
	Header         string   `json:"header"`
	OfferCode      string   `json:"offerCode"`
	OfferName      string   `json:"offerName"`
	Message        string   `json:"message"`
	ResultStatus   string   `json:"resultStatus"`
	IsRetryMessage bool     `json:"isRetryMessage"`
	ErrorMessage   string   `json:"errorMessage"`
	ElapsedTime    int64    `json:"elapsedTime"`
}
