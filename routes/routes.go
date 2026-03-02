package routes

import (
	"example.com/booking-project/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// This function can be used to register all routes in one place
	// functions like getEvents, getEventById, createEvent are defined in routes/events.go
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	// setup a group of routes that require authentication
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// functions like signup are defined in routes/users.go
	server.POST("/signup", signup)
	server.POST("/login", login)
}
