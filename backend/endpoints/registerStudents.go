package endpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User Story #1
func RegisterStudents(c* gin.Context) {
	var request struct {
		Teacher string `json:"teacher"`
		Students []string `json:"students"`
	}

	// Handle invalid request body
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle empty inputs
	if request.Teacher == "" || len(request.Students) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Teacher and students cannot be empty"})
		return
	}


	pgxConn, connCtx, err := GetConnection(c)
	if err != nil {

	}

	var teacherId uint64
	query := fmt.Sprintf("SELECT id from public.users where email='%s'", request.Teacher)
	err = pgxConn.QueryRow(connCtx, query).Scan(&teacherId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentIds []uint64
	students, err := sliceToInClause(request.Students)

	query = fmt.Sprintf("SELECT id from public.users where email in (%s)", students)
	ids, err := pgxConn.Query(connCtx, query)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for ids.Next() {
		var uid uint64
		err := ids.Scan(&uid)
		if err != nil {
			log.Fatal(err)
		}
		studentIds = append(studentIds, uid)
	}

	// create teacher-student relationship (fail silently if already exists)
	for _, studentId := range studentIds {
		query = fmt.Sprintf("INSERT INTO public.user_tags (teacher_id, student_id) VALUES (%d, %d) ON CONFLICT DO NOTHING", teacherId, studentId)
		_, err = pgxConn.Exec(connCtx, query)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Should it be 204 though, feel that we should indicate some form of indication (200 e.g "Students registered successfully")
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Students registered successfully"})
}