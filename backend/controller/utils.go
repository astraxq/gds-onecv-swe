package controller

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetConnection(c* gin.Context) (*sql.DB, error){
	if c == nil {
		return nil, errors.New("gin context not found")
	}

	var pgxDB *sql.DB

	db, ok := c.Get("db")
	if !ok {
		return nil, errors.New("database not found")
	}

	pgxDB = db.(*sql.DB)

	return pgxDB, nil
}

