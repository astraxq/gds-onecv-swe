package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

//make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() ( *sql.DB, error) {

   err := godotenv.Load()//by default, it is .env so we don't have to write
   if err != nil {
      fmt.Println("Error is occurred  on .env file please check")
   }
   //we read our .env file
   host := os.Getenv("HOST")
   port := os.Getenv("PORT") 
   user := os.Getenv("USER")
//    password := os.Getenv("PASSWORD")
   dbname := os.Getenv("DB_NAME")

   // connect to database 
   connOpts := fmt.Sprintf("postgresql://%s@%s:%s/%s", user, host, port, dbname)

   db, err := sql.Open("pgx", connOpts)
   
   return db, err
}