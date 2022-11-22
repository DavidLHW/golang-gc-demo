package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	os.Setenv("GOMEMLIMIT", "30MiB")
	os.Setenv("GOGC", "100")
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func TestGetConfigJson(t *testing.T) {
	test := struct {
		want configJson
	}{
		want: configJson{
			GOMEMLIMIT: "30MiB",
			GOGC:       "100",
		},
	}

	originalGOGC := os.Getenv("GOGC")
	defer os.Setenv("GOGC", originalGOGC)

	originalMemoryLimit := os.Getenv("GOMEMLIMIT")
	defer os.Setenv("GOMEMLIMIT", originalMemoryLimit)

	os.Setenv("GOGC", test.want.GOGC)
	os.Setenv("GOMEMLIMIT", test.want.GOMEMLIMIT)

	got := getConfigJson()

	if got != test.want {
		t.Errorf("TestGetConfigJson: want %q, got %q", test.want, got)
	}
}

func TestReadConfigHandler(t *testing.T) {
	test := struct {
		want configJson
	}{
		want: configJson{
			GOMEMLIMIT: "30MiB",
			GOGC:       "100",
		},
	}

	originalGOGC := os.Getenv("GOGC")
	defer os.Setenv("GOGC", originalGOGC)

	originalMemoryLimit := os.Getenv("GOMEMLIMIT")
	defer os.Setenv("GOMEMLIMIT", originalMemoryLimit)

	os.Setenv("GOGC", test.want.GOGC)
	os.Setenv("GOMEMLIMIT", test.want.GOMEMLIMIT)

	r := SetUpRouter(false)
	r.GET("/config", readConfigHandler)

	req, _ := http.NewRequest("GET", "/config", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	responseData := w.Body.Bytes()
	got := configJson{}
	if err := json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if got != test.want {
		t.Errorf("TestReadConfigHandler: want %q, got %q", test.want, got)
	}
}

func TestUpdateConfigHandler(t *testing.T) {
	test := struct {
		initial configJson
		want    configJson
	}{
		initial: configJson{
			GOMEMLIMIT: "100MiB",
			GOGC:       "off",
		},
		want: configJson{
			GOMEMLIMIT: "30MiB",
			GOGC:       "100",
		},
	}

	originalGOGC := os.Getenv("GOGC")
	defer os.Setenv("GOGC", originalGOGC)

	originalMemoryLimit := os.Getenv("GOMEMLIMIT")
	defer os.Setenv("GOMEMLIMIT", originalMemoryLimit)

	os.Setenv("GOGC", test.initial.GOGC)
	os.Setenv("GOMEMLIMIT", test.initial.GOMEMLIMIT)

	r := SetUpRouter(false)
	r.POST("/config", updateConfigHandler)

	requestPayload, err := json.Marshal(test.want)
	if err != nil {
		t.Errorf("error marshalling request payload: %v", err)
	}
	req, _ := http.NewRequest(
		"POST",
		"/config",
		strings.NewReader(string(requestPayload)),
	)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	responseData := w.Body.Bytes()
	got := configJson{}
	if err := json.Unmarshal(responseData, &got); err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if got != test.want {
		t.Errorf("TestUpdateConfigHandler: want %q, got %q", test.want, got)
	}
}
