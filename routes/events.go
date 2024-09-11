package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mhmdhalawi/events-booking/models"
)

func getEvents(c *gin.Context) {

	events := models.GetAllEvents()

	c.JSON(http.StatusOK, events)

}

func createEvent(c *gin.Context) {

	var event models.Event

	if c.ShouldBindJSON(&event) == nil {
		event.Save()
		c.JSON(http.StatusCreated, event)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot create event"})
	}

}

func eventRoutes(route *gin.RouterGroup) {
	eventRouter := route.Group("/events")

	{
		eventRouter.GET("/", getEvents)
		eventRouter.POST("/", createEvent)
	}
}
