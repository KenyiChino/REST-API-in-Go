package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

type Attendee struct {
	Id      int `json:"id"`
	UserId  int `json:"userId"`
	EventId int `json:"eventId"`
}

func (m *AttendeeModel) Insert(attendee *Attendee) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO attendees (event_id, user_id) VALUES ($1, $2) RETURNING id"
	err := m.DB.QueryRowContext(ctx, query,attendee.EventId, attendee.UserId).Scan(&attendee.id)

	if err != nil {
		return nil, err
	}

	return attendee, nil
}

func (m *AttendeeModel) GetByEventAndAttendee(eventId, userId int) (*Attendee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT * FROM attendees where event_id = $1 AND user_id = $2"
	
	var attendee Attendee
	err := m.DB.QueryRowContext(ctx, query, eventId, userId).Scan(&attendee.Id, &attendee.UserId, &attendee.EventId)
	
}
