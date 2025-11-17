package routes

import (
	model "OnlineServer/events"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handelEvent(context *gin.Context) {
	e, err := model.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, e)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func PostEvent(context *gin.Context) {
	var e model.Event
	err := context.ShouldBindJSON(&e)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := context.GetInt64("userID")
	e.UserID = userID

	err = e.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "done yasta", "event": e})
}

func UpdateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userID := context.GetInt64("userID")
	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to update event!"})
		return
	}

	var updatedEvent model.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.Id = eventId

	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func DeleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	userID := context.GetInt64("userID")
	event, err := model.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized to delete event!"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event deleted successfully"})
}
