package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"config"
	"dataService"
	"strconv"
)

func getActors(c *gin.Context){
	cn, cnErr := config.LoadDbConfig()
	if cnErr != nil{
		c.JSON(424, gin.H{
			"status":"failed",
			"error": fmt.Sprintf("%b", cnErr),
		})
	}

	srv, srvErr := dataService.Initialize(cn)
	if srvErr != nil {
		c.JSON(424, gin.H{
			"status": "failed",
			"error": fmt.Sprintf("%b", srvErr),
		})
	}

	frame, fErr := srv.GetActors()
	if fErr != nil{
		c.JSON(424, gin.H{
			"status": "failed",
			"error": fmt.Sprintf("%b", fErr),
		})
	}

	c.JSON(200, gin.H{
		"headers": frame.Headers,
		"rows": frame.Rows,
	})
}

func checkHealth(c *gin.Context){
	cn, cnErr := config.LoadDbConfig()
	if cnErr != nil{
		c.JSON(424, gin.H{
			"status": "424",
			"state": "drag brah no config!",
		})
	}

	srv, srvErr := dataService.Initialize(cn)
	if srvErr != nil{
		c.JSON(424, gin.H{
			"status":"424",
			"state": "Awww naawwww bra! Couldn't initialize a dataService!",
		})
	}

	rowCount, rcErr := srv.HealthCheck()
	if rcErr != nil{
		c.JSON(424, gin.H{
			"status": "424",
			"state": "very un-dude",
		})
	}

	c.JSON(200, gin.H{
		"status": "200", 
		"db": "available",
		"state": "funky. sexy. cool.",
		"actorCount": strconv.Itoa(rowCount),
	})
}
