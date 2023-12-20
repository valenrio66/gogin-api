package main

import (
	"github.com/valenrio66/gogin-api/config"
	"github.com/valenrio66/gogin-api/route"
)

func main() {
	db, err := config.DBMySQL()
	if err != nil {
		// Handle the error (e.g., log it and exit)
		panic(err)
	}

	// Close the database connection when the main function exits
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// Set up routes and start the server
	r := route.SetupRouter(db)
	r.Run(":8080")
}
