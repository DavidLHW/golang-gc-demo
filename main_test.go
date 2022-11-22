package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	os.Setenv("GOMEMLIMIT", "30MiB")
	os.Setenv("GOGC", "100")
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
