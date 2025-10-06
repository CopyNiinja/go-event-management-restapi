package models

import (
	"time"

	"github.com/copyniinja/go-event-management-restapi/db"
	"github.com/copyniinja/go-event-management-restapi/queries"
)

type Event struct {
	ID          int64
	Title       string `binding:"required"` //`json:"foo" xml:"foo" binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	UserID  int `json:"user_id"`
}


//save to database
func (e Event) Save()error{
   //saving into database

   //insert a event query
   stmt,err:=db.DB.Prepare(queries.InsertEvent);  

   //error
   if err !=nil{
	return err
   }

   //closing the statement
   defer stmt.Close()

   //executing the query with parameters
   result,err:=stmt.Exec(e.Title,e.Description,e.Location,e.Date,e.UserID)  // Exec() -> to changes data Query() -> to query database(return rows) 
   
   if err !=nil{
	return err
   }

   id,err:=result.LastInsertId()

   if err!=nil{
	return err
   }
   //adding the ID (got from inserting into db) to event
   e.ID=id;
   return nil

}

//get all events
func GetAllEvents()[]Event{
   //ignored error for simplicity. (DO NOT TRY THIS AT HOME! XD)
   

   //get all the rows after execution the query
   rows,_ :=db.DB.Query(queries.GetAllEvents);
   
   //closing the statement after finishing
   defer rows.Close();
   
   //slice for appending each row(each event)
   var events []Event;

   //iterate over each row
   for rows.Next(){
      //each event variable
      var e Event;
      //scanning (the order of the argument matters.Arguments should be in the order they were created in the table)
      rows.Scan(&e.ID,&e.Title,&e.Description,&e.Location,&e.Date,&e.UserID)
      events=append(events, e);
   }
   //return all the events
   return  events
}

//get event by id
func GetEvent(id string)(*Event,error){
    //event
    var event Event;
    
    //get the row of the event from db 
    row:= db.DB.QueryRow(queries.GetEventById,id);
    
    //scan all the row column value into event variable
    err:=row.Scan(&event.ID,&event.Title,&event.Description,&event.Location,&event.Date,&event.UserID);
    
    //error handling
    if err!=nil{
      return nil,err
    }
    //returning the event
    return &event,nil


}