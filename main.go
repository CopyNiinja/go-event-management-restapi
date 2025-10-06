package main

import (
	"flag"
	"strconv"

	"github.com/copyniinja/go-event-management-restapi/db"
	"github.com/copyniinja/go-event-management-restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//DB
	db.InitDB()

	//port
	var port  = 8080;
	flag.IntVar(&port,"port",port,"server port number");
    
	//version of api
	var version string;
	flag.StringVar(&version,"v","v1","the version of api endpoint")
    //parsing the flag
	flag.Parse()

	// Create a Gin router with default middleware (logger and recovery)
	r:=gin.Default()

    //routes
	 routes.EventRoutes(r,version); 

	

	//run and listening the server
	r.Run(":"+strconv.Itoa(port));
}

