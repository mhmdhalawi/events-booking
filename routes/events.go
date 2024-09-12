package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	var event models.Event

	if c.ShouldBindJSON(&event) == nil {
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

func eventsRoutes(route *gin.RouterGroup) {
	eventsRouter := route.Group("/events")

	{
		eventsRouter.GET("", getEvents)
		eventsRouter.POST("", createEvent)
	}
}
