// urandserver - Run a server to server random numbers
package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	//	"log"
)

func RunServer(router *gin.Engine, port string) {
	var portSuffix = fmt.Sprintf(":%s", port)
	router.Run(portSuffix)
}

func AddRoutes(router *gin.Engine) {
	router.GET("/health", healthHandler)
	router.GET("/random/:randomType", randomTypeHandler)
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

func randomTypeHandler(c *gin.Context) {
	randomType := c.Param("randomType")
	total := c.DefaultQuery("total", "10")
	c.String(http.StatusOK, "%s total %s", randomType, total)
}

