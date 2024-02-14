package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DropTable(db *sql.DB, tableName string) {
    _, err := db.Exec( fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName))
    if err != nil {
        log.Fatalf("Unable to drop table users: %v\n", err)
    }

    fmt.Println("Table users dropped successfully")
}

func SeedUsers(db *sql.DB) {
    // Drop the table first
    DropTable(db, "users")

    // Seed the users table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS public.users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            email VARCHAR(255) NOT NULL,
            role INT NOT NULL,
            status INT NOT NULL,
            notification_allowed BOOLEAN NOT NULL,
            CONSTRAINT email_unique UNIQUE (email)
        );

        CREATE INDEX idx_email ON public.users (email);
    `)
    if err != nil {
        log.Fatalf("Unable to create table users: %v\n", err)
    }

    // Seed the users table
    _, err = db.Exec( `
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

func SeedUserTags(db *sql.DB) {
    // Drop the table first
    DropTable(db, "user_tags")

    // Seed the user tag table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS public.user_tags (
            id SERIAL,
            teacher_id INT NOT NULL,
            student_id INT NOT NULL,
            CONSTRAINT pk_user_tags PRIMARY KEY (teacher_id, student_id)
        ); 

        CREATE INDEX idx_teacher_id ON public.user_tags (teacher_id);
    `)
    if err != nil {
        log.Fatalf("Unable to create table user_tags: %v\n", err)
    }

    fmt.Println("Data inserted successfully")
}