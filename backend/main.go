package main

import (
	"backend/teacher-admin-api/controller"
	"backend/teacher-admin-api/database"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Set("db", db)
        c.Next()
    }
}

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(ApiMiddleware(db))

	router.POST("/register", controller.RegisterStudents)
	router.GET("/commonstudents", controller.CommonStudents)
	router.POST("/suspend", controller.SuspendStudent)
	router.POST("/retrievefornotifications", controller.RetrieveForNotifications)
	return router
}

func main() {
    db, err := database.Init()
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
	defer db.Close()

	// Create tables in the database
	database.Migration(db)

    // Seed the users table
    database.SeedUsers(db)


	router := SetupRouter(db)
	router.Run("localhost:8080")
}
