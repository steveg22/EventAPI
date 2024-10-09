package routes

import (
	"database/sql"
	"example/mysql-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse ID"})
		return
	}
	event, err := models.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch Event"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func addEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse event"})
		return
	}

	id, err := models.AddEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create event"})
		return
	}

	event.Id = id
	c.JSON(http.StatusCreated, event)
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse Event Id"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update event"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch Event"})
		}
		return
	}

	updatedEvent.Id = event.Id
	err = models.UpdateEvent(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update event"})
		return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the Event ID"})
		return
	}

	_, err = models.GetEventById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Could not find Event with the given ID"})
		return
	}

	err = models.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete the given Event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event Deleted"})
}
