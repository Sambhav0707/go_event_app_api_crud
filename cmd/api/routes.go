package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler{
	g := gin.Default()
	v1 := g.Group("/api/v1")

	v1.POST("/event" , app.CreateEvent)
	v1.GET("/events" , app.GetAllEvents)
	v1.GET("/event/:id", app.GetEventById)
	v1.PUT("/event/:id" , app.UpdateEvent)
	v1.DELETE("/events/:id" , app.DeleteEvent)
	return g
}