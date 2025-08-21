package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	request2 "go-rest-api/model/dto/request"
	"go-rest-api/service"
	"go-rest-api/util"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
	logger          *logrus.Logger
}

func NewCategoryController(service service.CategoryService, logger *logrus.Logger) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: service,
		logger:          logger,
	}
}

func (controller *CategoryControllerImpl) InsertCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	traceId, _ := ctx.Value("traceId").(string)

	controller.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start API create new category: ", request.RequestURI)
	requestJson := request2.InsertCategoryReq{}
	util.ReadFromRequest(request, &requestJson)

	responseJson := controller.CategoryService.InsertCategory(ctx, requestJson)
	result := util.ConstructResponseSuccess(responseJson, traceId)

	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) SearchCategoryById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	traceId, _ := ctx.Value("traceId").(string)
	controller.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start API search category by id: ", request.RequestURI)
	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	responseJson := controller.CategoryService.SearchCategoryById(ctx, id)
	result := util.ConstructResponseSuccess(responseJson, traceId)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) SearchCategoryAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	traceId, _ := ctx.Value("traceId").(string)
	controller.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start API search category all: ", request.RequestURI)
	responseJson := controller.CategoryService.SearchCategoryAll(ctx)
	result := util.ConstructResponseSuccess(responseJson, traceId)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) UpdateCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	traceId, _ := ctx.Value("traceId").(string)
	controller.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start API update new category: ", request.RequestURI)
	requestJson := request2.UpdateCategoryReq{}
	util.ReadFromRequest(request, &requestJson)

	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	requestJson.Id = id

	responseJson := controller.CategoryService.UpdateCategory(ctx, requestJson)
	result := util.ConstructResponseSuccess(responseJson, traceId)
	util.WriteResponse(writer, result)
}

func (controller *CategoryControllerImpl) DeleteCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ctx := request.Context()
	traceId, _ := ctx.Value("traceId").(string)
	controller.logger.WithFields(logrus.Fields{"traceId": traceId}).Info("Start API delete new category: ", request.RequestURI)
	categoryId := params.ByName("categoryId")
	id, errConvert := strconv.Atoi(categoryId)
	helper.PanicIfError(errConvert)

	controller.CategoryService.DeleteCategory(ctx, id)
	result := util.ConstructResponseSuccess(nil, traceId)
	util.WriteResponse(writer, result)
}
