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
	r.GET("/events/:id",getEventById) //dynamic event: to get a single event using id

	//port
	var port  = 8080;
	flag.IntVar(&port,"port",port,"server port number");
    //parsing the flag
	flag.Parse()

	//run and listening the server
	r.Run(":"+strconv.Itoa(port));
}


// getAllEvents handles GET /events requests.
// It fetches all the events and returns JSON response
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
   //after successfully parsing body
     
   //dummy user id (TODO: Authentication,login ,sign up)
    event.UserID=100

   //saving the event to db  
    err:=event.Save();
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
	"message":"Failed to create event",
	
   })
	}
    
	
   //JSON response
   c.JSON(http.StatusCreated,gin.H{
	"message":"Event created successfully",
	"event":event,
   })
  }

}

//get single event by id
func getEventById(c *gin.Context){
  //the id of the event from params
  id := c.Param("id"); //events/1  -> id="1"
   
  //get the event 
  event,err:=models.GetEvent(id);
   
  //handling error
  if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{
		"message":"Failed to get the event:"+err.Error(),
	})
	return
  }
  //JSON response after successfully getting the event
  c.JSON(http.StatusOK,gin.H{
	"message":"success",
	"event":event,
  })

}