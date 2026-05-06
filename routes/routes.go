package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(server *gin.Engine) {
	server.GET("/events", GetAllEvent)
	server.GET("/events/:id", GetEventByID)
	server.POST("/events", CreateEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
}
