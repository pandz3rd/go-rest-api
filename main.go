package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"net/http"
)

func NewServer(handler *middleware.LogMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8085",
		Handler: handler,
	}
}

func NewHandler(router *httprouter.Router) *middleware.LogMiddleware {
	authMiddleware := middleware.NewAuthMiddleware(router)
	return middleware.NewLogMiddleware(authMiddleware)
}

func main() {
	//validate := validator.New()
	//db := app.NewDB()
	//
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//router := app.NewRouter(categoryController)
	//logMiddleware := NewHandler(router)
	//server := NewServer(logMiddleware)
	
	server := InitializeServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Server is running")
}
