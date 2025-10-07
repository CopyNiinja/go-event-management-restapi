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

//login handler
func LoginHandler(c *gin.Context){
  
  //getting user info from request body
  var user models.User;
  c.ShouldBindJSON(&user);

  //checking the user is a valid registered user
  fetchedUser,err:=models.GetUserByEmail(user.Email);
   
  //handling wrong credentials error
  if err!=nil{
    c.JSON(http.StatusUnauthorized,gin.H{
      "message":"Failed to login"+err.Error(),
    })
   return
  }
  //if the user id exists in registered user list
  //comparing the password
  if utils.CheckPasswordHash(user.Password,fetchedUser.Password) {
    //successfully login TODO: JWT implementation
       
    //dummy: TODO: remove this code!
     c.JSON(http.StatusOK,gin.H{
      "message":"Successfully logged In.",
    })
   return

  } else{
    //wrong credentials
     c.JSON(http.StatusUnauthorized,gin.H{
      "message":"Failed to login:wrong credentials"+err.Error(),
    })
   return
  }

}
