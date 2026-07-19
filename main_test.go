package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHello(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := newRouter()

	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}

	expectedBody := `{"message":"hello, stagepass"}`
	if res.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, res.Body.String())
	}
}

func TestHealthz(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := newRouter()

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, res.Code)
	}

	expectedBody := `{"status":"ok"}`
	if res.Body.String() != expectedBody {
		t.Errorf("Expected body %s, but got %s", expectedBody, res.Body.String())
	}
}
