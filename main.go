package main

import (
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/arl/statsviz"
	example "github.com/arl/statsviz/_example"
	"github.com/docker/go-units"
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
	// Force the GC to work to make the plots "move".
	go example.Work()

	r := gin.New()
	if withRoutes {
		r.GET("/config", readConfigHandler)
		r.POST("/config", updateConfigHandler)
		r.GET("/debug/statsviz/*filepath", statsvizHandler)
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

func updateConfigHandler(c *gin.Context) {
	var req *configJson
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Printf("requested config: %+v", req)

	requestedMemoryLimit, _ := units.FromHumanSize(req.GOMEMLIMIT)
	if requestedMemoryLimit > 0 {
		debug.SetMemoryLimit(requestedMemoryLimit)
		log.Printf("memory limit set to %d bytes", requestedMemoryLimit)
		os.Setenv("GOMEMLIMIT", req.GOMEMLIMIT)
	}

	requestedGOGC, _ := strconv.Atoi(req.GOGC)
	if requestedGOGC > 0 {
		debug.SetGCPercent(requestedGOGC)
		log.Printf("GOGC set to %d percent", requestedGOGC)
		os.Setenv("GOGC", req.GOGC)
	} else if req.GOGC == "off" {
		debug.SetGCPercent(-1)
		log.Printf("GOGC limit set to off")
		os.Setenv("GOGC", req.GOGC)
	}

	updated := getConfigJson()
	log.Printf("updated config: %+v", updated)

	c.IndentedJSON(http.StatusOK, req)
}

func statsvizHandler(c *gin.Context) {
	if c.Param("filepath") == "/ws" {
		statsviz.Ws(c.Writer, c.Request)
		return
	}
	statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(c.Writer, c.Request)
}
