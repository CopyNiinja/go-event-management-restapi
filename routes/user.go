package routes

import (
	"fmt"

	"github.com/copyniinja/go-event-management-restapi/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, version string) {

	//grouping auth routes
	{
		user := r.Group(fmt.Sprintf("/api/%s", version))

		//sign up
		user.POST("/signup", handlers.SignupHandler)
		//login
		user.POST("/login", handlers.LoginHandler)

	}

}
