package routes

import (
	"booking/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId") // Get the user ID from the context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId) // Retrieve the event by ID
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = event.Register(userId) // Call the RegisterUser method to register the user for the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User registered for event successfully", "event": event})

}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId") // Get the user ID from the context
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	var event models.Event
	event.ID = int(eventId)                // Set the event ID for the event struct
	err = event.CancelRegistration(userId) // Call the UnregisterUser method to unregister the user from the event
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User unregistered from event successfully", "event": event})
}
