package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
  // event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/:id", updateEvent)
	server.POST("/events", addEvent)
	server.DELETE("/events/:id", deleteEvent)

  // user routes
	server.POST("/signup", signup)
}
