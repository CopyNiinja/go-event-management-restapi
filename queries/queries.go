// Package queries provides query for managing events in database
package queries

const (
	GetAllEvents = `SELECT * FROM events`
	GetEventById = `SELECT * FROM events WHERE id=?`
	InsertEvent  = `INSERT INTO events(title,description,location,date,user_id) VALUES(?,?,?,?,?)`
	UpdateEvent  = `UPDATE events SET title=?, description=?, location=?, date=? WHERE id=?`
	DeleteEvent  = `DELETE FROM events WHERE id=?`
)
