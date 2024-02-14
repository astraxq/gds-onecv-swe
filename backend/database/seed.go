package database

import (
	"database/sql"
	"fmt"
	"log"
)


func SeedUsers(db *sql.DB) {
    // Seed the users table
    _, err := db.Exec( `
        INSERT INTO public.users (name, email, role, status, notification_allowed)
        VALUES
        ('Ken Doe', 'teacherken@example.com', 2, 1, true),
        ('Jane Smith', 'jane.smith@example.com', 3, 1, true),
        ('Alice Johnson', 'alice.johnson@example.com', 3, 1, true);
    `)
    if err != nil {
        log.Fatalf("Unable to insert data into users table: %v\n", err)
    }

    fmt.Println("Data inserted successfully")
}
