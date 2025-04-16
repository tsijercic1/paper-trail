package main

import (
	"paper-trail/internal/repository"
	"paper-trail/internal/service"

	"paper-trail/cmd/api-server/routes"

	"github.com/gin-gonic/gin"
)

var businessService *service.BusinessService

func init() {
	repo := repository.NewInMemoryBusinessRepository()
	businessService = service.NewBusinessService(repo)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	routes.RegisterBusinessRoutes(router, businessService)

	port := ":8080"
	router.Run(port)
}
