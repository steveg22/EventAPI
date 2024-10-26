package routes

import (
	"example/mysql-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the Event ID"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find Event with the given ID"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the Event ID"})
		return
	}

	var event models.Event
	event.Id = eventId

	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find Event with the given ID"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully Unregistered!"})
}

