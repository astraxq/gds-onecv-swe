package server

import (
	"backend/teacher-admin-api/controller"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// DbMiddleware will add the db connection to the context
func DbMiddleware(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Set("db", db)
        c.Next()
    }
}

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(DbMiddleware(db))

	api := router.Group("/api")
	{
		api.POST("/register", controller.RegisterStudents)
		api.GET("/commonstudents", controller.CommonStudents)
		api.POST("/suspend", controller.SuspendStudent)
		api.POST("/retrievefornotifications", controller.RetrieveForNotifications)
		api.POST("/seed", controller.Seed)
	}

	return router
}
