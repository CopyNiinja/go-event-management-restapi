// this package contains handlers for user,events
package handlers

import (
	"net/http"

	"github.com/copyniinja/go-event-management-restapi/models"
	"github.com/copyniinja/go-event-management-restapi/utils"
	"github.com/gin-gonic/gin"
)

//user handlers
//
func SignupHandler(c *gin.Context){
  //get user data from request body
   var user models.User
   //binding JSON
   err:=c.ShouldBindJSON(&user)
   //error handling
   if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{
		"message":"Failed to parse request body:"+err.Error(),
	})
	return
   }
   //hashing the password
   hashedPassword,err:= utils.HashPassword(user.Password);

   //handling error occurred in hashing password
   if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{
		"message":"Failed to hash user password:"+err.Error(),
	})
   }
   //rewriting the user raw password into hashed password
   user.Password=hashedPassword;
  
   //saving user to database
   err=user.Save()

    if err!=nil{
    c.JSON(http.StatusInternalServerError,gin.H{
		"message":"Failed to parse request body:"+err.Error(),
	})
	return
   }
   //JSON response
   c.JSON(http.StatusOK,gin.H{
		"message":"Successfully added the user",
		"user":user,
	})
}



