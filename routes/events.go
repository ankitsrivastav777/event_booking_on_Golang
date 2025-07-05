package routes

import (
	"booking/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func getEvent(context *gin.Context) {
	// This function will handle the /events/:id route
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Parse the event ID from the URL parameter
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Event not found"})
		return
	}
	event, err := models.GetEventByID(eventId) // Retrieve the event by ID using the GetEventByID function defined in the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event) // Respond with the event in JSON format
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

func updateEvent(context *gin.Context) {
	// This function will handle the update of an existing event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Parse the event ID from the URL parameter
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	_, err = models.GetEventByID(eventId) // Retrieve the event by ID using the GetEventByID function defined in the models package
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	var updateEvent models.Event
	err = context.BindJSON(&updateEvent) // Bind the JSON request body to the updateEvent struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateEvent.ID = int(eventId) // Set the ID of the event to be updated
	err = updateEvent.Update()    // Save the updated event using the Save method defined in the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updateEvent})
	// Respond with the updated event

}

func deleteEvent(context *gin.Context) {
	// This function will handle the deletion of an event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Parse the event ID from the URL parameter
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := models.GetEventByID(eventId) // Retrieve the event by ID using the GetEventByID function defined in the models package
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully", "event": event})
}
