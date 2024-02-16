package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

//make sure your function start with uppercase to call outside of the directory.
func Init() ( *sql.DB, error) {

   err := godotenv.Load()//by default, it is .env so we don't have to write
   if err != nil {
      fmt.Println("Error is occurred  on .env file please check")
   }
   //we read our .env file
//    host := os.Getenv("HOST")
//    port := os.Getenv("PORT") 
//    user := os.Getenv("POSTGRES_USER")
//    password := os.Getenv("POSTGRES_PASSWORD")
//    dbname := os.Getenv("POSTGRES_DB")

   // connect to database 
   connOpts := os.Getenv("DB_SOURCE")

   db, err := sql.Open("pgx", connOpts)
   
   return db, err
}

func DropTable(db *sql.DB, tableName string) {
   _, err := db.Exec( fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName))
   if err != nil {
       log.Fatalf("Unable to drop table users: %v\n", err)
   }

   fmt.Println("Table users dropped successfully")
}

func Migration(db *sql.DB) {
   // Drop the table first
   DropTable(db, "users")
   DropTable(db, "user_tags")

   // Create the users table
   _, err := db.Exec(`
       CREATE TABLE IF NOT EXISTS users (
           id SERIAL PRIMARY KEY,
           name VARCHAR(255) NOT NULL,
           email VARCHAR(255) NOT NULL,
           role INT NOT NULL,
           status INT NOT NULL,
           notification_allowed BOOLEAN NOT NULL,
           CONSTRAINT email_unique UNIQUE (email)
       );

       CREATE INDEX idx_email ON users (email);
   `)
   if err != nil {
       log.Fatalf("Unable to create table users: %v\n", err)
   }

   // Create the user tag table
   _, err = db.Exec(`
       CREATE TABLE IF NOT EXISTS user_tags (
           id SERIAL,
           teacher_id INT NOT NULL,
           student_id INT NOT NULL,
           CONSTRAINT pk_user_tags PRIMARY KEY (teacher_id, student_id)
       ); 

       CREATE INDEX idx_teacher_id ON user_tags (teacher_id);
   `)
   if err != nil {
       log.Fatalf("Unable to create table user_tags: %v\n", err)
   }

   fmt.Println("Tables created successfully")
}