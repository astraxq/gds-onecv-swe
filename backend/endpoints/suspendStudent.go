package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Story #3
func SuspendStudent(c* gin.Context) {
	var request struct {
		StudentEmail string `json:"Student"`
	}

	// Handle invalid request body
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle empty inputs
	if request.StudentEmail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Student email cannot be empty"})
		return
	}


	pgxConn, connCtx, err := GetConnection(c)
	if err != nil {

	}

	// Update student status to suspended
	query := fmt.Sprintf("UPDATE public.users SET status=%d where email='%s'", SUSPENDED, request.StudentEmail)
	_, err = pgxConn.Exec(connCtx, query)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Students successfully suspended"})
}