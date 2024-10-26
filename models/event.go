package models

import (
	"example/mysql-api/database"
	"fmt"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Datetime    time.Time `binding:"required"`
	UserId      int64
}

func GetAllEvents() ([]Event, error) {
	rows, err := database.DB.Query("SELECT * FROM events")
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var events []Event

	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId); err != nil {
			fmt.Println(err)
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event
	row := database.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	if err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserId); err != nil {
		return nil, err
	}

	return &event, nil
}

func AddEvent(event Event) (int64, error) {
	result, err := database.DB.Exec("INSERT INTO events (name, description, location, datetime, userid) VALUES (?, ?, ?, ?, ?)", event.Name, event.Description, event.Location, event.Datetime, event.UserId)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateEvent(event *Event) error {
	query := `
  UPDATE events
  SET name = ?, description = ?, location = ?, datetime = ?, userid = ?
  WHERE id = ?`

	_, err := database.DB.Exec(query, event.Name, event.Description, event.Location, event.Datetime, event.UserId, event.Id)

	return err
}

func DeleteEvent(id int64) error {
	_, err := database.DB.Exec("DELETE FROM events WHERE id = ?", id)

	return err
}

func (event *Event) Register(userId int64) error {
	_, err := database.DB.Exec("INSERT INTO registrations (eventid, userid) VALUES (?, ?)", event.Id, userId)

	return err
}

func (event *Event) Unregister(userId int64) error {
	_, err := database.DB.Exec("DELETE FROM registrations WHERE eventId = ? AND userid = ?", event.Id, userId)

	return err
}
