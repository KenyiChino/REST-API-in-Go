package main

import(
	"net/http"
	"rest-api-in-gin/internal/database"
	"github.com/gin-gonic/gin"
)

func(app *application) createEvent(c *gin.context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create event"})
		retun
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getEvent(c *gin.context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid event ID"})
	}
		event,err := app.models.Events.Get(id)

		if event == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to retreive event"})
		}

		c.JSON(http.StatusOK, event)
}

func (app* application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid event ID"})
		return
	}

	existingEvent, err := app.models.Events.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retreive event"})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}

	updateEvent := &database.Event{}
	
	if err := c.ShouldBindJSON(updateEvent); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
}