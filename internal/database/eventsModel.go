package database

import (
	"database/sql"
	"time"
)

type EventModel struct {
   db *sql.DB
}


type Event struct{
	Id          int       `json:"id"`
	OwnerId     int       `json:"owner_id" binding:"required"`
	Name        string    `json:"name" binding:"required,min=7"`
	Description string    `json:"description" binding:"required,min=10"`
	Date        time.Time `json:"date" binding:"required,datetime=2006-01-02"`
}
