// urandserver - Run a server to server random numbers
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	//	"log"
//	"time"

)

const (
	minsetSize int = 1
	maxsetSize int = 1000
	minRetentionSecs int = 1
	maxRetentionSecs int = 60
	minTotalSets int = 0
	maxTotalSets int = maxRetentionSecs
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
	fmt.Printf("randomSetHandler: %s retention %s setSize %s\n",randomType, retentionParam, setSizeParam)
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
		if ! validationError {
			randomSet,_ := Get(uint64(retention),randomType, uint64(setSize), uint64(totalSets))
			for i := 0 ; i < len(randomSet.entries); i++ {
				entry := randomSet.entries[i]
				switch randomType {
				case "uint64":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexUint64, "values": entry.valuesUint64})
				case "uint32":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexUint32, "values": entry.valuesUint32})
				case "uint16":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexUint16, "values": entry.valuesUint16})
				case "uint8":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexUint8, "values": entry.valuesUint8})
				case "int64":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexInt64, "values": entry.valuesInt64})
				case "int32":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexInt32, "values": entry.valuesInt32})
				case "int16":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexInt16, "values": entry.valuesInt16})
				case "int8":
					c.JSON(http.StatusOK, gin.H{"index": entry.indexInt8, "values": entry.valuesInt8})
				default:
					c.JSON(http.StatusBadRequest,  gin.H{"Error": fmt.Sprintf("Unknown random type %s",randomType)})
				}
			}

			//c.String(http.StatusOK, "%s setSize %d", randomType, setSize)
		}
	}
}

