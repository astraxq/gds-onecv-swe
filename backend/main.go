package main

import (
	"backend/teacher-admin-api/database"
	"backend/teacher-admin-api/server"
	"log"
)

func main() {
    db, err := database.Init()
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
	defer db.Close()

	router := server.SetupRouter(db)
	router.Run("localhost:8000")
}
