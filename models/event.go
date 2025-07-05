package models

import (
	"booking/rest-api/db"
	"time"
)

type Event struct {
	ID          int
	Name        string    `bind:"required"`
	Description string    `bind:"required"`
	Location    string    `bind:"required"`
	DateTime    time.Time `bind:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {

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
	if err != nil {
		return err
	}
	e.ID = int(id)
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Event) Update() error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, datetime = ?, user_id = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
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
