package main

import (
	"github.com/gin-gonic/gin"
	"github.com/moldaspan/assignment3/connections"
	"github.com/moldaspan/assignment3/routes"
)

func main() {
	router := gin.New()
	connections.Connect()
	routes.BookRoute(router)

	router.Run(":8080")
}
