package models

import (
	"time"

	"github.com/mhmdhalawi/events-booking/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id

	return err

}

func (e *Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, datetime = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e *Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {

		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err

}

func GetEventByID(id int64) (*Event, error) {

	var e Event

	err := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id).Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil

}

func GetAllEvents() ([]Event, error) {

	rows, err := db.DB.Query("SELECT * FROM events")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)

	}

	return events, nil
}
