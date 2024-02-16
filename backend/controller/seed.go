package controller

import (
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// QOL Endpoint
func Seed(c* gin.Context) {
	pgxDB, err := GetConnection(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database not found"})
		return
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql

	deleteQuery := psql.Delete("users")
	_, sqlErr := deleteQuery.RunWith(pgxDB).Exec()
	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message: fail to clear users table":  sqlErr.Error()})
		return
	}

	deleteQuery = psql.Delete("user_tags")
	_, sqlErr = deleteQuery.RunWith(pgxDB).Exec()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message: fail to clear users tags table": sqlErr.Error()})
		return
	}


	insertQuery := psql.Insert("users").Columns("name", "email", "role", "status").
		Values("Ken Doe", "teacherken@example.com", 2, 1).
		Values("Brian Quek", "brianquek@example.com", 2, 1).
		Values("John Tan", "johntan@example.com", 3, 1).
		Values("Jane Smith", "jane.smith@example.com", 3, 1).
		Values("Alice Johnson", "alice.johnson@example.com", 3, 1).
		Values("James Lee", "james.lee@example.com", 3, 1)


	_, sqlErr = insertQuery.RunWith(pgxDB).Exec()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message: fail to seed tables": sqlErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Users and User Tags successfully seeded"})
}