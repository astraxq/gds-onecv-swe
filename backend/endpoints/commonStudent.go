package endpoints

import (
	"fmt"
	"log"
	"net/http"

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
	pgxConn, connCtx, err := GetConnection(c)
	if err != nil {

	}

	// Get teacher ids from emails
	var teacherIds []uint64
	teacherEmailQueryString, err := sliceToInClause(teacherEmails)
	
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := fmt.Sprintf("SELECT id from public.users where email in (%s)", teacherEmailQueryString)
	res, err := pgxConn.Query(connCtx, query)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for res.Next() {
		var uid uint64
		err := res.Scan(&uid)
		if err != nil {
			log.Fatal(err)
		}
		teacherIds = append(teacherIds, uid)
	}

	// Handle when no teacher ids are found
	if len(teacherIds) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"students": []string{}})
		return
	}

	// Get students id only if it matches to all the teacher ids
	var studentIds []uint64
	teacherIdsQueryString, err := sliceToInClause(teacherIds)
	query = fmt.Sprintf(`SELECT student_id from public.user_tags where teacher_id in (%s)
		GROUP BY student_id HAVING COUNT(DISTINCT teacher_id) = %d
	`, teacherIdsQueryString, len(teacherIds))
	res, err = pgxConn.Query(connCtx, query)

	if err != nil {
		fmt.Println("HERE", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for res.Next() {
		var id uint64
		err := res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		studentIds = append(studentIds, id)
	}

	// Handle when no students are tagged to the teacher(s)
	if len(studentIds) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"students": []string{}})
		return
	}

	// Get student emails by student ids
	var studentEmails []string
	studentIdsQueryString, err := sliceToInClause(studentIds)
	fmt.Println(studentIdsQueryString)
	query = fmt.Sprintf("SELECT email from public.users where id in (%s)", studentIdsQueryString)
	res, err = pgxConn.Query(connCtx, query)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	for res.Next() {
		var email string
		err := res.Scan(&email)
		if err != nil {
			log.Fatal(err)
		}
		studentEmails = append(studentEmails, email)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"students": studentEmails})
}