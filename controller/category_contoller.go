package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	InsertCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SearchCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SearchCategoryAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
