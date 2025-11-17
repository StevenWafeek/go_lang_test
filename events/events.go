package model

import (
	"OnlineServer/db"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	Query := `
	INSERT INTO events (name, description, location, userID)
	VALUES (?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(Query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetEvents() ([]Event, error) {
	query := `SELECT id, name, description, location, userID FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event
		err = rows.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventByID(eventID int64) (*Event, error) {
	query := " SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, eventID)
	var e Event
	err := row.Scan(&e.Id, &e.Name, &e.Description, &e.Location, &e.UserID)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.Id)
	return err
}

func (e Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.Id)
	return err

}
