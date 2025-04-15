package main

import (
	"net/http"
	"paper-trail/internal/model"
	"paper-trail/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterBusinessRoutes(router *gin.Engine, businessService *service.BusinessService) {
	// Create a new business
	router.POST("/business", func(c *gin.Context) {
		var business model.Business
		if err := c.ShouldBindJSON(&business); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := businessService.CreateBusiness(&business); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, business)
	})

	// Get a business by ID
	router.GET("/business/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		business, err := businessService.GetBusinessByID(intID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, business)
	})

	// Update a business
	router.PUT("/business/:id", func(c *gin.Context) {
		var business model.Business
		if err := c.ShouldBindJSON(&business); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := businessService.UpdateBusiness(&business); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, business)
	})

	// Delete a business
	router.DELETE("/business/:id", func(c *gin.Context) {
		id := c.Param("id")
		intID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}
		if err := businessService.DeleteBusiness(intID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	})

	// Get paginated businesses
	router.GET("/business", func(c *gin.Context) {
		pageNumber, err := strconv.Atoi(c.DefaultQuery("pageNumber", "0"))
		if err != nil {
			pageNumber = 0
		}
		pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		if err != nil {
			pageSize = 10
		}

		offset := pageNumber * pageSize
		limit := pageSize

		businesses, err := businessService.GetBusinessesPage(offset, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, businesses)
	})
}
