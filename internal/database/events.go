package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Event struct {
	Id           int     `json:"id"`
	OwnerId      int     `json:"ownerId" binding:"required"`
	Name         string  `json:"name" binding:"required,min=3"`
	Description  string  `json:"description" binding:"required, min=10"`
	Date         string  `json:"date" binding:"required, datetime=2006-01-02"`
	Location     string  `json:"location" binding:"required, min=3"`
}

func (m *EventModel) Insert(event *Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO events (owner_id, name, description, date, location) VALUES ($1, $2, $3, $4, $5) "

	return m.DB.QueryRowContext(ctx, query, event.OwnerId, event.Name, event.Description, event.Date, event.Location).Scan(&event.Id)
}

func (m *EventModel) GetAll() ([]*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events"

	rows, err := m.DB.QueryContent(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []*Event{}

	for rows.Next(){
		var event Event

		err := rows.Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)

		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err = rows.Err(); err !=nil {
		return nil, err
	}

	return events, nil
}

func (m *EventModel) Get(id int) (*Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM events WHERE id = $1"

	var event Event
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&event.Id, &event.OwnerId, &event.Name, &event.Description, &event.Date, &event.Location)
	if err != nil {
		if err  == sql.ErrNoRows {
			return nil, nil
		}
		return nil,err
	}

	return &event,  nil
}

func (m *EventModel) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "DELETE FROM events WHERE id = $1"

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}