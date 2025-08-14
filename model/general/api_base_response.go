package general

type ApiBaseResponse struct {
	ResponseMessage string      `json:"responseMessage"`
	ResponseCode    string      `json:"responseCode"`
	Data            interface{} `json:"data"`
	TraceId         string      `json:"traceId"`
}
