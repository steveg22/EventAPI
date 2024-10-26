package routes

import (
	"database/sql"
	"example/mysql-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse ID"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch Event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func addEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse event"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserId = userId

	id, err := models.AddEvent(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create event"})
		return
	}

	event.Id = id
	context.JSON(http.StatusCreated, event)
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse Event Id"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update event"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch Event"})
		}
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	updatedEvent.Id = event.Id
	err = models.UpdateEvent(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update event"})
		return
	}

	context.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the Event ID"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find Event with the given ID"})
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	err = models.DeleteEvent(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete the given Event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
