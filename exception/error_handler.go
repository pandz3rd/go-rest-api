package exception

import (
	"github.com/go-playground/validator"
	"go-rest-api/model/general"
	"go-rest-api/util"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}
	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(ErrorNotFound)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		result := general.ApiBaseResponse{
			ResponseCode:    "99",
			ResponseMessage: "Call 911",
		}
		util.WriteResponse(writer, result)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		result := general.ApiBaseResponse{
			ResponseCode:    "99",
			ResponseMessage: "Bad Request",
		}
		util.WriteResponse(writer, result)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	result := general.ApiBaseResponse{
		ResponseCode:    "99",
		ResponseMessage: "Error adalah warna dalam kehidupan",
	}
	util.WriteResponse(writer, result)
}
