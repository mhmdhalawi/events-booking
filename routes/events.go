package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mhmdhalawi/events-booking/middlewares"
	"github.com/mhmdhalawi/events-booking/models"
)

func getEvents(c *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get events"})
		return
	}

	c.JSON(http.StatusOK, events)

}

func createEvent(c *gin.Context) {

	userID := c.GetInt64("userID")

	var event models.Event

	if c.ShouldBindJSON(&event) == nil {
		event.UserID = userID
		err := event.Save()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event"})
			return
		}

		c.JSON(http.StatusCreated, event)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event"})
	}

}

func getEvent(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get event"})
		return
	}

	c.JSON(http.StatusOK, event)

}

func updateEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}

	userID := c.GetInt64("userID")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get event"})
		return
	}

	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if c.ShouldBindJSON(&event) == nil {
		err := event.Update()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event"})
			return
		}

		c.JSON(http.StatusCreated, event)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create event"})
	}
}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id."})
		return
	}
	userID := c.GetInt64("userID")
	event, err := models.GetEventByID(eventID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot get event"})
		return
	}

	if event.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	err = event.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}

func eventsRoutes(route *gin.RouterGroup) {

	eventsRouter := route.Group("/events")

	protectedEventsRouter := route.Group("/events")
	protectedEventsRouter.Use(middlewares.Authenticate)

	{
		eventsRouter.GET("", getEvents)
		eventsRouter.GET("/:id", getEvent)
		protectedEventsRouter.POST("", createEvent)
		protectedEventsRouter.PUT("/:id", updateEvent)
		protectedEventsRouter.DELETE("/:id", deleteEvent)
	}
}
