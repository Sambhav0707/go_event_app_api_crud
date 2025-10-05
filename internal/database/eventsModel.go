package database

import (
	"context"
	"database/sql"
	"time"
)

type EventModel struct {
	Db *sql.DB
}

type Event struct {
	Id          int       `json:"id"`
	OwnerId     int       `json:"owner_id" binding:"required"`
	Name        string    `json:"name" binding:"required,min=7"`
	Description string    `json:"description" binding:"required,min=10"`
	Date        time.Time `json:"date" binding:"required,datetime=2006-01-02"`
	Location    string    `json:"location" binding:"required , min=3`
}

func (e *EventModel) Insert(event *Event) error {

	cntxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "INSERT INTO events (owner_id , name , description , date , location) VALUES ($1, $2 , $3 , $4 , $5)"

	return e.Db.QueryRowContext(cntxt, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(event.Id)

}

func (e *EventModel) GetAll() ([]*Event, error) {

	cntxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "SELECT * FROM events"

	rows, err := e.Db.QueryContext(cntxt, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []*Event

	for rows.Next() {
		var event Event

		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)

		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil

}

func (model *EventModel) GetEvent(id int) (*Event, error) {
	cntxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "SELECT id ,owner_id , name , description , date , location WHERE id = ?"

	var event Event

	err := model.Db.QueryRowContext(cntxt, query, id).Scan(

		&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err

	}

	return &event, nil

}

func (model *EventModel) Update(event *Event) error {
	cntxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "UPDATE events SET name = $1 , description = $2 , date = $3 , location = $4 WHERE id = $5"

	_, err := model.Db.ExecContext(cntxt, query, &event.Name, &event.Description, &event.Date, &event.Location, &event.Id)

	if err != nil {
		return err
	}

	return nil

}

func (model *EventModel) Delete(id int) error {
	cntxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := "DELETE FROM events WHERE id = $5"

	_, err := model.Db.ExecContext(cntxt, query, id)

	if err != nil {
		return err
	}

	return nil

}
