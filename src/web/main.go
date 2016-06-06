package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.RouterGroup.GET("/actors", getActors)
	engine.RouterGroup.GET("/health", checkHealth)
	engine.RouterGroup.GET("/config", checkConn)
	engine.Run()
}


