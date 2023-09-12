package routes

import (
	"admodev/invexchange/cmd/gincmd"
	usercontainer "admodev/invexchange/container/userContainer"

	"github.com/gin-gonic/gin"
)

var r = gincmd.R

func healthCheckRoute() {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
}

func createNewUser() {
	r.POST("/user", usercontainer.CreateUser)
}

func InitializeRoutes() {
	healthCheckRoute()
	createNewUser()
}
