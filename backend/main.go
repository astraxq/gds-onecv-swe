package main

import (
	"backend/teacher-admin-api/database"
	"backend/teacher-admin-api/endpoints"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)


func getStudentsByTeacherID(c* gin.Context) {
	// get students by teacher id
	c.IndentedJSON(http.StatusOK, []endpoints.User{})
}

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(ctx context.Context,conn *pgx.Conn) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Set("context", ctx)
        c.Set("databaseConn", conn)
        c.Next()
    }
}

func main() {
	ctx := context.Background()
    conn, err := database.ConnectDatabase(ctx)
	if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer conn.Close(ctx)

    // Seed the users table
    database.SeedUsers(conn)

	// Seed the user tag table
	database.SeedUserTags(conn)

	router := gin.Default()
	router.Use(ApiMiddleware(ctx, conn))
	router.POST("/register", endpoints.RegisterStudents)
	router.GET("/commonstudents", endpoints.CommonStudents)
	router.POST("/suspend", endpoints.SuspendStudent)
	router.POST("/retrievefornotifications", endpoints.RetrieveForNotifications)

	router.Run("localhost:8080")
}
