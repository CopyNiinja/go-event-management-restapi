// Package queries provides query for managing events in database
package queries

const (
	//users
	CreateUserTable=`CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL
	)`
    InsertUser = `INSERT INTO users(email,password) VALUES(?,?)`
	GetUserByEmail = `SELECT * FROM users WHERE email=?`
	//events
	CreateEventsTable=`CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL ,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	date DATETIME NOT NULL,
	user_id INTEGER ,
	FOREIGN KEY (user_id) REFERENCES users(id) 
	)`
	GetAllEvents = `SELECT * FROM events`
	GetEventById = `SELECT * FROM events WHERE id=?`
	InsertEvent  = `INSERT INTO events(title,description,location,date,user_id) VALUES(?,?,?,?,?)`
	UpdateEvent  = `UPDATE events SET title=?, description=?, location=?, date=? WHERE id=?`
	DeleteEvent  = `DELETE FROM events WHERE id=?`
)
