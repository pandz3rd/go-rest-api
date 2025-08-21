package util

import (
	"encoding/json"
	"go-rest-api/helper"
	"go-rest-api/model/general"
	"net/http"
)

func WriteResponse(writer http.ResponseWriter, result interface{}) {
	json.Marshal(result)
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	errEncode := encoder.Encode(result)
	helper.PanicIfError(errEncode)
}

func ConstructResponseSuccess(responseJson interface{}, traceId string) general.ApiBaseResponse {
	return general.ApiBaseResponse{
		ResponseCode:    "00",
		ResponseMessage: "SUCCESS",
		Data:            responseJson,
		TraceId:         traceId,
	}

}
