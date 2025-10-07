// user model
package models

import (
	"github.com/copyniinja/go-event-management-restapi/db"
	"github.com/copyniinja/go-event-management-restapi/queries"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	//preparing the query
	stmt, err := db.DB.Prepare(queries.InsertUser)
	//handling error
	if err != nil {
		return err
	}
	//closing the statement
	defer stmt.Close()

	//executing the query
	//parameters: email,password
	res, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}
	//saving the last inserted id
	u.ID, err = res.LastInsertId()

	return err

}

//GetuserById takes an argument email and return the user if its a registered user or error if not
func GetUserByEmail(email string) (*User, error) {
	//database query
	row := db.DB.QueryRow(queries.GetUserByEmail, email)
	var user User
	//scanning the row to get the user details arg ID,Email,Password
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
