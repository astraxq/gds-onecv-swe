package controller

import (
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// User Story #2
func CommonStudents(c* gin.Context) {
	teacherEmails := c.QueryArray("teacher")

	// Handle invalid params
	if len(teacherEmails) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Teacher cannot be empty"})
		return
	}

	// Get connection and context
	pgxDB, err := GetConnection(c)
	if err != nil {

	}

	// Get teacher ids from emails
	var teacherIds []uint64
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql
	sqQuery := psql.Select("id").From("public.users").Where(sq.Eq{"email": teacherEmails})
	rows, sqlErr := sqQuery.RunWith(pgxDB).Query()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var uid uint64
		err := rows.Scan(&uid)
		if err != nil {
			log.Fatal(err)
		}
		teacherIds = append(teacherIds, uid)
	}

	// Get students id only if it matches to all the teacher ids
	var studentIds []uint64
	sqQuery = psql.Select("student_id").From("public.user_tags").Where(sq.Eq{"teacher_id": teacherIds}).
			GroupBy("student_id").Having("COUNT(DISTINCT teacher_id) = ?", len(teacherIds))
	rows, sqlErr = sqQuery.RunWith(pgxDB).Query()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var uid uint64
		err := rows.Scan(&uid)
		if err != nil {
			log.Fatal(err)
		}
		studentIds = append(studentIds, uid)
	}

	// No common students found
	if len(studentIds) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"students": []string{}})
		return
	}

	// Get student emails by student ids
	var studentEmails []string
	sqQuery = psql.Select("email").From("public.users").Where(sq.Eq{"id": studentIds})
	rows, sqlErr = sqQuery.RunWith(pgxDB).Query()

	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		studentEmails = append(studentEmails, email)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"students": studentEmails})
}