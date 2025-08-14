package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	request2 "go-rest-api/model/dto/request"
	"go-rest-api/service"
	"go-rest-api/util"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}

func (controller *CategoryControllerImpl) InsertCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestJson := request2.InsertCategoryReq{}
	util.ReadFromRequest(request, &requestJson)

	ctx := request.Context()

	responseJson := controller.CategoryService.InsertCategory(ctx, requestJson)
	result := util.ConstructResponseSuccess(responseJson)

	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) SearchCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	responseJson := controller.CategoryService.SearchCategoryById(ctx, id)
	result := util.ConstructResponseSuccess(responseJson)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) SearchCategoryAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	responseJson := controller.CategoryService.SearchCategoryAll(ctx)
	result := util.ConstructResponseSuccess(responseJson)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestJson := request2.UpdateCategoryReq{}
	util.ReadFromRequest(request, &requestJson)

	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	requestJson.Id = id
	ctx := request.Context()

	responseJson := controller.CategoryService.UpdateCategory(ctx, requestJson)
	result := util.ConstructResponseSuccess(responseJson)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	controller.CategoryService.DeleteCategory(ctx, id)
	result := util.ConstructResponseSuccess(nil)
	util.WriteResponse(writer, result)
}
