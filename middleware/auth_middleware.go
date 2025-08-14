package middleware

import (
	"go-rest-api/model/general"
	"go-rest-api/util"
	"net/http"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "MULYONOSECRETSERVICE" == request.Header.Get("X-API-Key") {
		middleware.handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		result := general.ApiBaseResponse{
			ResponseCode:    "99",
			ResponseMessage: "Unauthorized",
		}
		util.WriteResponse(writer, result)
	}
}
