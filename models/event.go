package models

import "time"

type Event struct {
	ID          int
	Title       string `binding:"Required"` //`json:"foo" xml:"foo" binding:"required"`
	Description string `binding:"Required"`
	Location    string `binding:"Required"`
	Date        time.Time `binding:"Required"`
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