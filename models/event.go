package models

import "time"

type Event struct {
	ID          int
	Title       string `binding:"required"` //`json:"foo" xml:"foo" binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time `binding:"required"`
	UserID  int
}

//database
var events []Event=[]Event{};

func (e Event) Save(){
   //saving into database
	events=append(events,e);
}

func GetAllEvents()[]Event{
	return events
}