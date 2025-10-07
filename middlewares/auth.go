package middlewares

import (
	"net/http"
	"strings"

	"github.com/copyniinja/go-event-management-restapi/utils"
	"github.com/gin-gonic/gin"
)

//Authorization func is a middleware that check and verify the json token and  add userid
func Authorization(c *gin.Context) {

  //get token from header
  bearerToken:= c.Request.Header.Get("Authorization");
  //checks if its empty
  if bearerToken == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
        c.Abort()
        return
    }
  //get the token part //Bearer <token>
  parts:=strings.Split(bearerToken, " ");
  
  //verify token
  id,err:=utils.VerifyToken(parts[1]);

   if err!=nil{
	//In middleware using gin,Need to Abort ,otherwise response will be send from next middlewares 
	c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
		"message":"Unauthorized access.please login",
	})
	return
   }
   //save the id 
   c.Set("id",id);
   
   //next() 
   c.Next()

}