package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//secret key
var SecretKey = []byte("SECRET_KEY123")

//GenerateToken func takes user_id and email and generate JSON web token signed with secret key.
func GenerateToken(userID int64,email string)(string,error) {
     
	//generating token
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id":userID,
		"email":email,
		 "exp":time.Now().Add(time.Minute*5),
		
	})
	//signed token with secret key
	tokenString,err:= token.SignedString(SecretKey);
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}

//VerifyToken func takes the JWT token and verifies it 
func VerifyToken(token string)error{
	//parsing
   parsedToken,err:= jwt.Parse(token,func(token *jwt.Token)(interface{},error){ 
		
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
            return nil,errors.New("unaccepted signing method") 
		}
		return SecretKey,nil})

	//error handling
	if err!=nil{
		return err
	} 
	//check if the token is invalid
	if !parsedToken.Valid{
		return errors.New("invalid token")
	}
	//check the token claims is valid 
	 claims,ok:=parsedToken.Claims.(jwt.MapClaims)

	 if !ok{
		return errors.New("invalid token")
	 }
     //by default it was any type.thats why type checking
	 email,_:=claims["email"].(string)
	 id:=int64(claims["id"].(float64)) //!! float64 auto converted by Claim store and retrieve process
     
	 fmt.Println(email,id);
	return nil

}