package endpoints

import (
	"log"
	"net/http"

	sq "github.com/Masterminds/squirrel"
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


	pgxDB, err := GetConnection(c)
	if err != nil {

	}

	var teacherId uint64
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql
	sqQuery := psql.Select("id").From("public.users").Where(sq.Eq{"email": request.Teacher})
	err = sqQuery.RunWith(pgxDB).QueryRow().Scan(&teacherId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var studentIds []uint64
	sqQuery = psql.Select("id").From("public.users").Where(sq.Eq{"email": request.Students})
	rows, err := sqQuery.RunWith(pgxDB).Query()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	// create teacher-student relationship (fail silently if already exists)
	for _, studentId := range studentIds {
		insertQuery := psql.Insert("public.user_tags").Columns("teacher_id", "student_id").Values(teacherId, studentId).Suffix("ON CONFLICT DO NOTHING")
		_, err = insertQuery.RunWith(pgxDB).Exec()
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	// Should it be 204 though, feel that we should indicate some form of indication (200 e.g "Students registered successfully")
	c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Students registered successfully"})
}