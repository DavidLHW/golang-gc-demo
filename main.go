package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type configJson struct {
	GOMEMLIMIT string `json:"gomemlimit"` // in human bytes
	GOGC       string `json:"gogc"`       // in integer percentage
}

func main() {
	r := SetUpRouter(true)
	r.Run(":8080")
}

func SetUpRouter(withRoutes bool) *gin.Engine {
	r := gin.New()
	if withRoutes {
		r.GET("/config", readConfigHandler)
	}
	return r
}

func getConfigJson() configJson {
	var configJson configJson

	configJson.GOMEMLIMIT = os.Getenv("GOMEMLIMIT")
	configJson.GOGC = os.Getenv("GOGC")
	return configJson
}

func readConfigHandler(c *gin.Context) {
	currentConfig := getConfigJson()
	log.Printf("current config: %+v", currentConfig)
	c.IndentedJSON(http.StatusOK, currentConfig)
}
