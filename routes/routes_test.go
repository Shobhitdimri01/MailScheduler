package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var testRoutesTable = []struct {
	Method   string
	Endpoint string
}{
	{"GET", "/api/health"},
}

func TestRegisterRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	RegisterRoutes(r, nil) // Pass nil for logger if not required
	apiEndpoint := r.Routes()

	for _, v := range apiEndpoint {
		found := false
		for _, testroutes := range testRoutesTable {
			if v.Method == testroutes.Method && v.Path == testroutes.Endpoint {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("%s api route is not registered in testRoutesTable with method %s", v.Path, v.Method)
		}
	}

}

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Register a single test route !important to register
	r.GET("/api/health", Health)

	req, _ := http.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Error("Unexpected status code")
	}
	var m map[string]string
	json.NewDecoder(w.Result().Body).Decode(&m)
	if m["status"] != "ok" {
		t.Errorf("Unexpected response %s Expected was : <ok>", m["status"])
	}
}
