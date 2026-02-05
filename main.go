package main

import (
	"net/http"

	"example.com/booking-project/db"
	"example.com/booking-project/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // initialize the database
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD
	server.POST("/events", createEvent)
	server.Run(":8080") // localhost
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
