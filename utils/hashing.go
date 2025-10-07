package utils

import "golang.org/x/crypto/bcrypt"

//HashPassword func takes raw user inputted password and hash the password using bcrypt
func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	//return
	return string(hashedPassword), err
}

//CheckPasswordHash func takes password and hash string stored in DB and compares them and return the boolean value
func CheckPasswordHash(password, hash string) bool {
	//CompareHashAndPassword returns err if the password and hash string are not equal
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
