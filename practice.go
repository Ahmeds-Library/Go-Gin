package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET: List of users",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST: New user created",
		})
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE: User " + id + " deleted",
		})
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT: User " + id + " updated",
		})
	})

	router.PATCH("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "PATCH: User " + id + " partially updated",
		})
	})

	router.Run(":8080")
}
