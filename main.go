package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/repository"
	"go-rest-api/service"
	"net/http"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	logMiddleware := middleware.NewLogMiddleware(authMiddleware)

	server := http.Server{
		Addr:    "localhost:8085",
		Handler: logMiddleware,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println("Server is running")
}
