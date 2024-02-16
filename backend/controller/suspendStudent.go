package controller

import (
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// User Story #3
func SuspendStudent(c* gin.Context) {
	var request struct {
		StudentEmail string `json:"Student"`
	}

	// Handle invalid request body
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message: invalid json body": err.Error()})
		return
	}

	// Handle empty inputs
	if request.StudentEmail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "student field cannot be empty"})
		return
	}


	pgxDB, err := GetConnection(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database not found"})
		return
	}

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql
	updateQuery := psql.Update("users").Set("status", SUSPENDED).Where(sq.Eq{"email": request.StudentEmail})

	_, sqlErr := updateQuery.RunWith(pgxDB).Exec()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message: fail to update student status": sqlErr.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Students successfully suspended"})
}