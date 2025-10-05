package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/copyniinja/go-event-management-restapi/db"
	"github.com/copyniinja/go-event-management-restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	//DB
	db.InitDB()

	// Create a Gin router with default middleware (logger and recovery)
	r:=gin.Default()

   //endpoints
	r.GET("/events",getAllEvents)
    r.POST("/events",createEvent)

	//port
	var port  = 8080;
	flag.IntVar(&port,"port",port,"server port number");
    //parsing the flag
	flag.Parse()

	//run and listening the server
	r.Run(":"+strconv.Itoa(port));
}

func getAllEvents(c *gin.Context){
  //Get all events
  events:=models.GetAllEvents()

  //JSON response
  c.JSON(http.StatusOK,gin.H{
	"events":events,
  })
}

func createEvent(c *gin.Context){
  var event models.Event;
  //binding event with request body
  if  err:=c.ShouldBindJSON(&event);err!=nil{
	//error response
	c.JSON(http.StatusBadRequest,gin.H{
		"message":"Failed to parse Request",
	})

  }else{
   //successfully parsing body
   //saving the event to db   
   event.Save();
    
    //dummy data
    event.ID=1
	event.UserID=100
   //response
   c.JSON(http.StatusCreated,gin.H{
	"message":"Event created successfully",
	"event":event,
   })
  }

}