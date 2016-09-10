package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/tflynn/gobasics/urand/server"
)

func main() {
	var port = flag.String("p", "8080", "Specify port")
	flag.Parse()

	router := gin.Default()

	server.AddRoutes(router)
	server.RunServer(router, *port)
}

