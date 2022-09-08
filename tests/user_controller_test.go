package tests

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"simpleProfile/app"
	"simpleProfile/controllers"
	"simpleProfile/helpers"
	"simpleProfile/middleware"
	"simpleProfile/repositories"
	"simpleProfile/services"
	"strconv"
	"strings"
	"testing"
	"time"
)

const TOKEN = "INITOKENRAHASIA"
const BASE_URL = "http://localhost:8080"

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_simple_profile")
	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserServiceImpl(userRepository, db, validate)
	userController := controllers.NewUserController(userService)

	router := app.NewRouter(userController)

	return middleware.NewAuthMiddleware(router)
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE user")
}

func TestRegisterSuccess(t *testing.T) {
	db := setupTestDB()
	truncateUser(db)
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"username" : "testing", "password" : "abcde54321", "firstname" : "testing"}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/register", reqBody)
	request.Header.Add("Content-Type", "application/json")
	//request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 200, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "OK", bodyResponse["status"])
	assert.Equal(t, "Berhasil Registrasi dengan user name testing", bodyResponse["data"].(map[string]interface{})["message"])
}

func TestRegisterFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/register", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 400, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "bad request", bodyResponse["status"])
}

func TestLoginSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"username" : "testing", "password" : "abcde54321"}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/login", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 200, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "OK", bodyResponse["status"])
	assert.Equal(t, "INITOKENRAHASIA", bodyResponse["data"].(map[string]interface{})["token"])
}

func TestLoginFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"username" : "testin", "password" : "abcde54321"}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/login", reqBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 404, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "not found", bodyResponse["status"])
}

func TestGetProfileSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	userId := 1
	request := httptest.NewRequest(http.MethodGet, BASE_URL+"/profile/"+strconv.Itoa(userId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 200, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "OK", bodyResponse["status"])
	assert.Equal(t, userId, int(bodyResponse["data"].(map[string]interface{})["id"].(float64)))
}

func TestGetProfileFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	userId := 99
	request := httptest.NewRequest(http.MethodGet, BASE_URL+"/profile/"+strconv.Itoa(userId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 404, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "not found", bodyResponse["status"])
}

func TestUpdateProfileSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	userId := 1
	reqBody := strings.NewReader(`{"lastname" : "testing"}`)
	request := httptest.NewRequest(http.MethodPut, BASE_URL+"/profile/"+strconv.Itoa(userId), reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 200, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "OK", bodyResponse["status"])
	assert.Equal(t, userId, int(bodyResponse["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "testing", bodyResponse["data"].(map[string]interface{})["Lastname"])
}

func TestProfileFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	userId := 1
	//reqBody := strings.NewReader(`{"lastname" : "testing"}`)
	request := httptest.NewRequest(http.MethodPut, BASE_URL+"/profile/"+strconv.Itoa(userId), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 500, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 500, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "internal server error", bodyResponse["status"])
}

func TestLogoutSucess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	reqBody := strings.NewReader(`{"id" : 1}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/logout", reqBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 200, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "OK", bodyResponse["status"])
	assert.Equal(t, "Berhasil Logout!", bodyResponse["data"].(map[string]interface{})["message"])
}

func TestLogoutFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	//reqBody := strings.NewReader(`{"id" : 1}`)
	request := httptest.NewRequest(http.MethodPost, BASE_URL+"/logout", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", TOKEN)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 500, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 500, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "internal server error", bodyResponse["status"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)

	userId := 1
	request := httptest.NewRequest(http.MethodGet, BASE_URL+"/profile/"+strconv.Itoa(userId), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 401, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var bodyResponse map[string]interface{}
	json.Unmarshal(body, &bodyResponse)

	assert.Equal(t, 401, int(bodyResponse["code"].(float64)))
	assert.Equal(t, "Unauthorized!", bodyResponse["status"])
}
