package endpoints

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func GetConnection(c* gin.Context) (*pgx.Conn, context.Context, error){
	var pgxConn *pgx.Conn
	var connCtx context.Context

	// get conn and context
	conn, ok := c.Get("databaseConn")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: ": "Database connection not found"})
		return &pgx.Conn{}, nil, errors.New("Database connection not found")
	}

	ctx, ok := c.Get("context")
	if !ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Context not found"})
		return &pgx.Conn{}, nil, errors.New("Context not found")
	}

	pgxConn = conn.(*pgx.Conn)
	connCtx = ctx.(context.Context)

	return pgxConn, connCtx, nil
}

func sliceToInClause(slice interface{}) (string, error) {
    switch v := slice.(type) {
    case []uint64:
        s := make([]string, len(v))
        for i, val := range v {
            s[i] = fmt.Sprint(val)
        }
        return strings.Join(s, ","), nil
    case []string:
        s := make([]string, len(v))
        for i, val := range v {
            s[i] = fmt.Sprintf("'%s'", val)
        }
        return strings.Join(s, ","), nil
    default:
        return "", fmt.Errorf("unsupported type: %T", slice)
    }
}