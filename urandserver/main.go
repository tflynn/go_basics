// urandserver - Run a server to server random numbers
package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	//	"log"
)

func runServer(port string) {
	var portSuffix = fmt.Sprintf(":%s", port)
	iris.Listen(portSuffix)
}

func addRoutes() {
	iris.HandleFunc("GET", "/health", healthHandler)
	iris.HandleFunc("GET", "/random/:randomType", randomTypeHandler)
}

func healthHandler(c *iris.Context) {
	c.Write("ok")
}

func randomTypeHandler(c *iris.Context) {
	randomType := c.Param("randomType")
	c.Write(randomType)
}

func main() {
	var port = flag.String("p", "8080", "Specify port")
	flag.Parse()

	addRoutes()
	runServer(*port)
}
