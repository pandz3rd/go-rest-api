package app

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/controller"
	"go-rest-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/category/api/v1/list", categoryController.SearchCategoryAll)
	router.POST("/category/api/v1/add", categoryController.InsertCategory)
	router.GET("/category/api/v1/get/:categoryId", categoryController.SearchCategoryById)
	router.PUT("/category/api/v1/edit/:categoryId", categoryController.UpdateCategory)
	router.DELETE("/category/api/v1/delete/:categoryId", categoryController.DeleteCategory)

	router.PanicHandler = exception.ErrorHandler

	return router
}
