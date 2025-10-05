package main

import (
	"net/http"
	"strconv"

	"github.com/Sambhav0707/go_event_app_api_crud/internal/database"
	"github.com/gin-gonic/gin"
)

func (app *application) CreateEvent(g *gin.Context) {

	// 1st we will bind the incoming json request to the Event struct

	var event database.Event
	if err := g.ShouldBindJSON(&event); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2nd now after a successfull binding we have to insert it in the database
	err := app.models.Events.Insert(&event)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create an Event"})
		return
	}

	//3rd now we have to return the event

	g.JSON(http.StatusAccepted, event)

}

func (app *application) GetAllEvents(g *gin.Context) {
	// Directly get the event from thte database

	events, err := app.models.Events.GetAll()

	if err != nil {
		g.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, events)

}

func (app *application) GetEventById(g *gin.Context) {
	//convert the id from the params to int

	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event, err := app.models.Events.GetEvent(id)

	if event == nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": "No event found",
		})
		return
	}
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, event)

}

func (app *application) UpdateEvent(g *gin.Context) {
	//convert the id from the params to int

	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//get the event from the DB
	event, err := app.models.Events.GetEvent(id)
	if event == nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": "No event found",
		})
		return
	}
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// now bind the req to the updatedEvent struct

	updatedEvent := &database.Event{}

	if err := g.ShouldBindJSON(updatedEvent); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent.Id = id

	if err := app.models.Events.Update(updatedEvent); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//now return the user the updated event

	g.JSON(http.StatusOK, updatedEvent)

}

func (app *application) DeleteEvent(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	g.JSON(http.StatusNoContent, nil)

}
