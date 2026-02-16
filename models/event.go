package models

import (
	"time"

	"example.com/booking-project/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)` // Use parameterized queries to prevent SQL injection

	// PREPARE MEANS TO COMPILE THE QUERY AND CHECK FOR SYNTAX ERRORS
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	// execute means to run the query with the provided parameters
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId() // Get the ID of the newly inserted event
	if err != nil {
		return err
	}

	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the rows and scan the data into Event structs
	for rows.Next() {
		var event Event
		// Scan the row data into the event struct
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event

	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) // QueryRow is used for queries that return a single row

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events 
	SET name = ?, description = ?, location = ?, date_time = ? 
	WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}
