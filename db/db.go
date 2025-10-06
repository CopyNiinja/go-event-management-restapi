package db

import (
	"database/sql"
	"fmt"

	"github.com/copyniinja/go-event-management-restapi/queries"
	_ "github.com/mattn/go-sqlite3" // this driver is used under the hood by database/sql
)

//initialize

var DB *sql.DB; // type return by sql.Open() 

func InitDB(){
   var err error;
  DB,err=sql.Open("sqlite3","./api.db") // driver,datasource location (must end with .db)

  //error in opening db
  if err!=nil {
      panic("Failed to connect database") //program will stop
  }
  
  //connection pool options
  DB.SetMaxOpenConns(10) //maximum number of open connections
  DB.SetMaxIdleConns(5) //maximum number of idle connections

  //after successfully initialization:create the events table:
  createTable(); 
}

//create table function
func createTable(){
 //creating user table
  _,err:=DB.Exec(queries.CreateUserTable);
  if err!=nil{
    panic(fmt.Sprintf("Failed to create users table: %v", err))
  }
//creating events table
 if _,err=DB.Exec(queries.CreateEventsTable);err !=nil{
	panic(fmt.Sprintf("Failed to create events table: %v", err))
 }

}