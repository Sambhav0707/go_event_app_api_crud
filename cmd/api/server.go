package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) server() error {
	server :=  &http.Server{
		Addr: fmt.Sprintf(":%d" , app.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}


	fmt.Printf("Server starting on port: %d" ,app.port)

	return server.ListenAndServe()
}