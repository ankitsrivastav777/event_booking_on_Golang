package main

import (
	"booking/rest-api/db"
	"booking/rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // Initialize the database connection
	server := gin.Default()

	// Define a simple route
	server.GET("/events", getEvents)
	server.POST("/events", createEvent) // Create a new event

	// Start the server on port 8080
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}

func getEvents(context *gin.Context) {
	// This function will handle the /events route
	events, err := models.GetAllEvents() // Retrieve all events using the GetAllEvents function defined in the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
	// Respond with the list of events in JSON format
}

func createEvent(context *gin.Context) {
	// This function will handle the creation of a new event
	var event models.Event

	if err := context.BindJSON(&event); err != nil {
		// If there is an error binding the JSON, respond with a bad request
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = 1        // Assign a new ID to the event (in a real application, this would be handled by the database)
	event.UserID = 1    // Assign a user ID (in a real application, this would	 be derived from the authenticated user)
	err := event.Save() // Save the event using the Save method defined in the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Respond with the created event
	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
	// In a real application, you would typically return the created event with its ID and other
}
