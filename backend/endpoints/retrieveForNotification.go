package endpoints

import (
	"fmt"
	"net/http"
	"regexp"

	"strings"

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


	pgxConn, connCtx, err := GetConnection(c)
	if err != nil {

	}

	// Scrape notification string and retrieve all student emails mentioned in it
	mentionedStudents := getMentionedStudents(request.Notification)

	// Get teacher id
	var teacherId uint64
	query := fmt.Sprintf("SELECT id from public.users where email='%s'", request.TeacherEmail)
	err = pgxConn.QueryRow(connCtx, query).Scan(&teacherId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get students id registered with teacher id
	var idSlice []uint64
	query = fmt.Sprintf("SELECT student_id from public.user_tags where teacher_id=%d", teacherId)
	studentIds, err := pgxConn.Query(connCtx, query)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for studentIds.Next() {
		var uid uint64
		err := studentIds.Scan(&uid)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		idSlice = append(idSlice, uid)
	}

	fmt.Println(idSlice)

	// If no students are registered with teacher id, return mentioned students
	if len(idSlice) == 0 {
		c.IndentedJSON(http.StatusOK, gin.H{"recipients": mentionedStudents})
		return
	}

	// Get student emails registered with teacher id and mentioned in notification
	var students []string
	idSliceString, err := sliceToInClause(idSlice)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	query = fmt.Sprintf("SELECT email from public.users where id in (%s)", idSliceString)
	if mentionedStudents != nil {
		query += fmt.Sprintf(" OR email in ('%s')", strings.Join(mentionedStudents, "','"))
	}

	studentEmails, err := pgxConn.Query(connCtx, query)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for studentEmails.Next() {
		var email string
		err := studentEmails.Scan(&email)
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