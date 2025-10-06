package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine,version string) {

	//grouping auth routes
	{
	   auth:=r.Group(fmt.Sprintf("/api/%s",version));
       
	   //sign up
	   auth.POST("/signup")
	   //login
	   auth.POST("/login")

	}

}
