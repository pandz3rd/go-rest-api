package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"go-rest-api/app"
	"go-rest-api/controller"
	"go-rest-api/helper"
	"go-rest-api/middleware"
	"go-rest-api/model/dao"
	"go-rest-api/model/general"
	"go-rest-api/repository"
	"go-rest-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func NewDBTest() *sql.DB {
	db, err := sql.Open("mysql", "root:masuk();@tcp(127.0.0.1:3306)/learn_test?parseTime=true")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	return middleware.NewLogMiddleware(authMiddleware)
}

func clearDb(db *sql.DB) {
	db.Exec("TRUNCATE category;")
}

func insertSampleCategory(db *sql.DB) int {
	categoryRepository := repository.NewCategoryRepository()

	tx, _ := db.Begin()
	newCategory := categoryRepository.Insert(context.Background(), tx, dao.Category{
		Name: "Barang Haram",
	})
	tx.Commit()

	return newCategory.Id
}

func TestCreateCategorySuccess(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	requestString := strings.NewReader("{\"name\":\"Baran terlarang\"}")
	request := httptest.NewRequest(http.MethodPost, "/category/api/v1/add", requestString)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestCreateCategoryFailed(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	requestString := strings.NewReader("{}")
	request := httptest.NewRequest(http.MethodPost, "/category/api/v1/add", requestString)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Bad Request", responseJson.ResponseMessage)
}

func TestFindAllCategorySuccess(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	insertSampleCategory(db)

	request := httptest.NewRequest(http.MethodGet, "/category/api/v1/list", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestFindAllCategoryFailed(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/category/api/v1/list", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestFindIdCategorySuccess(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	id := insertSampleCategory(db)

	request := httptest.NewRequest(http.MethodGet, "/category/api/v1/get/"+strconv.Itoa(id), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestFindIdCategoryFailed(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/category/api/v1/get/"+"5", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, "Call 911", responseJson.ResponseMessage)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	id := insertSampleCategory(db)

	requestString := strings.NewReader("{\"name\":\"Baran terlarang Edited\"}")
	request := httptest.NewRequest(http.MethodPut, "/category/api/v1/edit/"+strconv.Itoa(id), requestString)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	requestString := strings.NewReader("{\"name\":\"Baran terlarang Edited\"}")
	request := httptest.NewRequest(http.MethodPut, "/category/api/v1/edit/"+"1", requestString)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, "Call 911", responseJson.ResponseMessage)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	id := insertSampleCategory(db)

	request := httptest.NewRequest(http.MethodDelete, "/category/api/v1/delete/"+strconv.Itoa(id), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "SUCCESS", responseJson.ResponseMessage)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "/category/api/v1/delete/"+"5", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", "MULYONOSECRETSERVICE")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, "Call 911", responseJson.ResponseMessage)
}

func TestUnauthorized(t *testing.T) {
	db := NewDBTest()
	clearDb(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "/category/api/v1/list", nil)
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	bodyByte, _ := io.ReadAll(response.Body)
	responseJson := general.ApiBaseResponse{}
	json.Unmarshal(bodyByte, &responseJson)
	fmt.Println(responseJson)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.Equal(t, "Unauthorized", responseJson.ResponseMessage)
}
