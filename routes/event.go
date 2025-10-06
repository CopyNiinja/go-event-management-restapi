package routes

import (
	"fmt"

	"github.com/copyniinja/go-event-management-restapi/handlers"
	"github.com/gin-gonic/gin"
)

func EventRoutes(r *gin.Engine,version string) {
    
	//grouping routes for /events path
	{
	   events:=r.Group(fmt.Sprintf("/api/%s/events",version))
       //get event by ID
	   events.GET("/:id",handlers.GetEventById)
	   //get all events 
	   events.GET("/",handlers.GetAllEvents)
	   //create an event
	   events.POST("/",handlers.CreateEvent)
	   //update an event by ID
	   events.PUT("/:id",handlers.UpdateEventById)
	
	}
}