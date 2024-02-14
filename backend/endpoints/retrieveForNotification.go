package endpoints

import (
	"net/http"
	"regexp"

	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

// User Story #4
func RetrieveForNotifications(c* gin.Context) {
	var request struct {
		TeacherEmail string `json:"teacher"`
		Notification string `json:"notification"`
	}

	// Handle invalid request body
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Handle empty inputs
	if request.TeacherEmail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Student email cannot be empty"})
		return
	}


	pgxDB, err := GetConnection(c)
	if err != nil {

	}

	// Scrape notification string and retrieve all student emails mentioned in it
	mentionedStudents := getMentionedStudents(request.Notification)

	// Get teacher id
	var teacherId uint64
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql
	sqQuery := psql.Select("id").From("public.users").Where(sq.Eq{"email": request.TeacherEmail})
	
	sqlErr := sqQuery.RunWith(pgxDB).QueryRow().Scan(&teacherId)
	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	// Get students id registered with teacher id
	var studentIDs []uint64
	sqQuery = psql.Select("student_id").From("public.user_tags").Where(sq.Eq{"teacher_id": teacherId})
	rows, sqlErr := sqQuery.RunWith(pgxDB).Query()
	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var uid uint64
		err := rows.Scan(&uid)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		studentIDs = append(studentIDs, uid)
	}

	// Get student emails registered with teacher id and mentioned in notification
	var students []string
	sqQuery = psql.Select("email").From("public.users")

	if len(studentIDs) == 0 && mentionedStudents != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"recipients": []string{}})
		return
	}

	sqQuery = sqQuery.Where(sq.Or{sq.Eq{"id": studentIDs}, sq.Eq{"email": mentionedStudents}}).Where(sq.NotEq{"status": SUSPENDED})

	rows, sqlErr = sqQuery.RunWith(pgxDB).Query()
	if sqlErr != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		students = append(students, email)
	}


	c.IndentedJSON(http.StatusOK, gin.H{"recipients": students})
}


func getMentionedStudents(notification string) []string {
	re := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	emails := re.FindAllString(notification, -1)
	return emails
}