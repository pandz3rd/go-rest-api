//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

var middlewareSet = wire.NewSet(
	app.NewRouter,
	NewHandler,
)

var serverSet = wire.NewSet(
	middlewareSet,
	NewServer,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		serverSet,
	)
	return nil
}
