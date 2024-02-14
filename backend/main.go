package main

import (
	"backend/teacher-admin-api/database"
	"backend/teacher-admin-api/endpoints"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func getStudentsByTeacherID(c* gin.Context) {
	// get students by teacher id
	c.IndentedJSON(http.StatusOK, []endpoints.User{})
}

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Set("db", db)
        c.Next()
    }
}

func main() {
    db, err := database.ConnectDatabase()
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
	defer db.Close()

    // Seed the users table
    database.SeedUsers(db)

	// Seed the user tag table
	database.SeedUserTags(db)

	router := gin.Default()
	router.Use(ApiMiddleware(db))
	router.POST("/register", endpoints.RegisterStudents)
	router.GET("/commonstudents", endpoints.CommonStudents)
	router.POST("/suspend", endpoints.SuspendStudent)
	router.POST("/retrievefornotifications", endpoints.RetrieveForNotifications)

	router.Run("localhost:8080")
}
