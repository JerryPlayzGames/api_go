package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message" : "Could not fetch events. Please try again later!"})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return
	}

	event, err := models.GetEventbyID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse data"})
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create event. Please try again later!"})
	}

	context.JSON(http.StatusCreated, gin.H{"message" : "Event created!", "event" : event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return
	}

	_,err = models.GetEventbyID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the event"})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Could not parse data"})
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return
	}

	event, err := models.GetEventbyID(eventId)

	event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Event deleted successfully"})
}