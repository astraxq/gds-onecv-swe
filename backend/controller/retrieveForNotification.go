package controller

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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error: invalid json body": err.Error()})
		return
	}

	// Handle empty inputs
	if request.TeacherEmail == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "teacher field cannot be empty"})
		return
	}


	pgxDB, err := GetConnection(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: ": "Database not found"})
		return
	}

	// Scrape notification string and retrieve all student emails mentioned in it
	mentionedStudents := getMentionedStudents(request.Notification)

	// Get teacher id
	var teacherId uint64
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar) // required for psql
	sqQuery := psql.Select("id").From("users").Where(sq.Eq{"email": request.TeacherEmail})
	
	sqlErr := sqQuery.RunWith(pgxDB).QueryRow().Scan(&teacherId)
	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: fail to get teacher id": sqlErr.Error()})
		return
	}

	// Get students id registered with teacher id
	var studentIDs []uint64
	sqQuery = psql.Select("student_id").From("user_tags").Where(sq.Eq{"teacher_id": teacherId})
	rows, sqlErr := sqQuery.RunWith(pgxDB).Query()
	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: fail to get students ids": sqlErr.Error()})
		return
	}

	for rows.Next() {
		var uid uint64
		err := rows.Scan(&uid)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: fail to scan student ids": err.Error()})
			return
		}
		studentIDs = append(studentIDs, uid)
	}

	// Get student emails registered with teacher id and mentioned in notification
	sqQuery = psql.Select("email").From("users")

	// exit safely if no students are found
	if len(studentIDs) == 0 && mentionedStudents != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"recipients": []string{}})
		return
	}

	sqQuery = sqQuery.Where(sq.Or{sq.Eq{"id": studentIDs}, sq.Eq{"email": mentionedStudents}}).Where(sq.NotEq{"status": SUSPENDED})

	rows, sqlErr = sqQuery.RunWith(pgxDB).Query()
	if sqlErr != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: fail to get students emails": sqlErr.Error()})
		return
	}

	var studentEmails []string
	for rows.Next() {
		var email string
		err := rows.Scan(&email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error: fail to scan student emails": err.Error()})
			return
		}
		studentEmails = append(studentEmails, email)
	}


	c.IndentedJSON(http.StatusOK, gin.H{"recipients": studentEmails})
}


func getMentionedStudents(notification string) []string {
	re := regexp.MustCompile(`\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,}\b`)
	emails := re.FindAllString(notification, -1)
	return emails
}