package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"go-rest-api/app"
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

func NewHandler(router *httprouter.Router, logger *logrus.Logger) *middleware.LogMiddleware {
	authMiddleware := middleware.NewAuthMiddleware(router)
	return middleware.NewLogMiddleware(authMiddleware, logger)
}

func main() {
	logger := app.InitLogger()
	server := InitializeServer(logger)
	logger.Info("Server start at port ", server.Addr)
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
