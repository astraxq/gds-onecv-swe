package endpoints

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetConnection(c* gin.Context) (*sql.DB, error){
	var pgxDB *sql.DB

	db, ok := c.Get("db")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: ": "Database not found"})
		return nil, errors.New("Database not found")
	}

	pgxDB = db.(*sql.DB)

	return pgxDB, nil
}
