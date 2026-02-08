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
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request body"})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}
