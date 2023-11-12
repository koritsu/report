package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"restapi-go/config"
	"restapi-go/database"
	"restapi-go/logging"
	"strings"
	"testing"
)

func init() {
	var err error

	err = config.LoadConfig()
	if err != nil {
		fmt.Printf("Config error: %s", err)
		os.Exit(1)
	}

	appConfig := config.Config()

	err = logging.InitLog(appConfig)
	if err != nil {
		fmt.Printf("Log error: %s", err)
		os.Exit(1)
	}

	// init DB
	err = database.InitDB()
	if err != nil {
		fmt.Printf("DB error: %s", err)
		os.Exit(1)
	}
}

func TestReadNotFound(t *testing.T) {
	rPath := "/read/:id"
	router := gin.Default()
	router.GET(rPath, Read)
	req, _ := http.NewRequest("GET", "/read/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "response status should be 404")
}

func TestCreate(t *testing.T) {
	rPath := "/create"
	router := gin.Default()
	router.POST(rPath, Create)

	req, _ := http.NewRequest("POST", "/create", strings.NewReader("{\"name\":\"test.png\"}"))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "response status should be 400")
}

func TestUpdate(t *testing.T) {
	rPath := "/update/:id"
	router := gin.Default()
	router.PUT(rPath, Update)

	req, _ := http.NewRequest("PUT", "/update/1", strings.NewReader("{\"name\":\"test.png\"}"))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "response status should be 400")
}

func TestDelete(t *testing.T) {
	rPath := "/delete/:id"
	router := gin.Default()
	router.DELETE(rPath, Delete)

	req, _ := http.NewRequest("DELETE", "/delete/1", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "response status should be 404")
}
