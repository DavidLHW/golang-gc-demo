package main

import "github.com/gin-gonic/gin"

func main() {
	r := SetUpRouter()
	r.Run(":8080")
}

func SetUpRouter() *gin.Engine {
	r := gin.New()
	return r
}
