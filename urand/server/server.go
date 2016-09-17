// urandserver - Run a server to server random numbers
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	//	"log"
	//	"time"
)

const (
	minsetSize       int = 1
	maxsetSize       int = 1000
	minRetentionSecs int = 1
	maxRetentionSecs int = 60
	minTotalSets     int = 0
	maxTotalSets     int = maxRetentionSecs
)

func RunServer(router *gin.Engine, port string) {
	var portSuffix = fmt.Sprintf(":%s", port)
	router.Run(portSuffix)
}

func AddRoutes(router *gin.Engine) {
	router.GET("/health", healthHandler)
	router.GET("/random/set/:randomType", randomSetHandler)
}

func healthHandler(c *gin.Context) {
	//c.String(http.StatusOK, "ok")
	// var msg struct {
	// 	Status string
	// }
	// msg.Status = "ok"
	msg := gin.H{"status": "OK"}
	c.JSON(http.StatusOK, msg)
}

func randomSetHandler(c *gin.Context) {
	randomType := c.Param("randomType")
	setSizeParam := c.DefaultQuery("setSize", "10")
	setSize, err1 := strconv.Atoi(setSizeParam)
	retentionParam := c.DefaultQuery("retention", "1")
	retention, err2 := strconv.Atoi(retentionParam)
	totalSetsParam := c.DefaultQuery("sets", "0")
	totalSets, err3 := strconv.Atoi(totalSetsParam)
	//fmt.Printf("randomSetHandler: %s retention %s setSize %s totalSets %s\n", randomType, retentionParam, setSizeParam, totalSetsParam)
	if err1 != nil || err2 != nil || err3 != nil {
		if err1 != nil {
			c.String(http.StatusBadRequest, "Invalid setSize %s", randomType, setSizeParam)
		}
		if err2 != nil {
			c.String(http.StatusBadRequest, "Invalid retention period %s", randomType, retentionParam)
		}
		if err3 != nil {
			c.String(http.StatusBadRequest, "Invalid total sets %s", randomType, totalSetsParam)
		}
	} else {
		validationError := false
		if setSize < minsetSize || setSize > maxsetSize {
			c.String(http.StatusBadRequest, "Invalid setSize %s", randomType, setSizeParam)
			validationError = true
		}
		if retention < minRetentionSecs || retention > maxRetentionSecs {
			c.String(http.StatusBadRequest, "Invalid retention period %s", randomType, retentionParam)
			validationError = true
		}
		if totalSets < minTotalSets || totalSets > maxTotalSets {
			c.String(http.StatusBadRequest, "Invalid total sets  %s", randomType, totalSetsParam)
			validationError = true
		}
		//TODO Needs stronger validation
		if !strings.Contains(randomType, "int") {
			c.String(http.StatusBadRequest, "Invalid randomType  %s", randomType)
			validationError = true
		}
		if totalSets == 0 {
			totalSets = retention
		}

		if !validationError {
			randomSet, _ := Get(uint64(retention), randomType, uint64(setSize), uint64(totalSets))
			outputObjs := make([]interface{}, totalSets, totalSets)
			for i := 0; i < totalSets; i++ {
				entry := randomSet.entries[i]
				outputObj := make(map[string]interface{})
				outputObj["index"] = entry.index
				outputObj["values"] = entry.values
				outputObjs[i] = outputObj
			}
			c.JSON(http.StatusOK, outputObjs)
		}
	}
}
