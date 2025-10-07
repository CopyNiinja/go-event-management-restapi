package handlers

import (
	"net/http"
	"strconv"

	"github.com/copyniinja/go-event-management-restapi/models"
	"github.com/gin-gonic/gin"
)

// getAllEvents handles GET /events requests.
// It fetches all the events and returns JSON response
func GetAllEvents(c *gin.Context){
  //Get all events
  events:=models.GetAllEvents()

  //JSON response
  c.JSON(http.StatusOK,gin.H{
	"events":events,
  })
}


//create a event
func CreateEvent(c *gin.Context){
  var event models.Event;
  //binding event with request body
  if  err:=c.ShouldBindJSON(&event);err!=nil{
	//error response
	c.JSON(http.StatusBadRequest,gin.H{
		"message":"Failed to parse Request",
	})

  }else{
   //after successfully parsing body
     
  //  //dummy user id ( Authentication,login ,sign up)
  //   event.UserID=100
    id,_:=c.Get("id")
    //any type
    idVal := id.(int64)
    event.UserID=idVal

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
func GetEventById(c *gin.Context){
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

//update an event by ID
func UpdateEventById(c *gin.Context){
  //get the id from params
  id:=c.Param("id")
  //query to get the event
  staleEvent,err:= models.GetEvent(id)
  //handling error
  if err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{
      "Message":"Failed to get the event"+err.Error(),
    })
    return
  }
   var updatedEvent models.Event; 
  //get the updated event from body
  err=c.ShouldBindJSON(&updatedEvent);

  if err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{
      "Message":"Failed to parse the event "+err.Error(),
    })
   return
  }

  //set the id,userID of the event in updatedEvent
  updatedEvent.UserID=staleEvent.UserID ;
  updatedEvent.ID,err=strconv.ParseInt(id,10,64) ;
  

  if err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{
      "Message":"Failed to parse the id "+err.Error(),
    })
    return
  }

  err=updatedEvent.Update()
  //handling error while updating the db
  if err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{
      "Message":"Failed to update the event "+err.Error(),
    })
    return
  }

  ///JSON response for successfully updating the event
  c.JSON(http.StatusOK,gin.H{
    "message":"The event updated successfully",
    "updatedEvent":updatedEvent,
  })



}

//Delete an event by ID
func DeleteEventById(c *gin.Context){
  //get the id from params
  id:=c.Param("id")
  //query to get the event
  _,err:= models.GetEvent(id)
  //handling error
  if err!=nil {
    c.JSON(http.StatusBadRequest,gin.H{
      "Message":"Failed to get the event"+err.Error(),
    })
    return
  } 
  //deleting the event
  err= models.DeleteEvent(id);

  if err!=nil {
    c.JSON(http.StatusInternalServerError,gin.H{
      "Message":"Failed to delete the event"+err.Error(),
    })
    return
  } 
  //JSON response after successfully deleting the event
  c.JSON(http.StatusOK,gin.H{
      "Message":"successfully deleted the event.",
    })

}