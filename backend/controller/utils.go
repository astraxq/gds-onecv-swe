package controller

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetConnection(c* gin.Context) (*sql.DB, error){
	var pgxDB *sql.DB

	db, ok := c.Get("db")
	if !ok {
		return nil, errors.New("database not found")
	}

	pgxDB = db.(*sql.DB)

	return pgxDB, nil
}

