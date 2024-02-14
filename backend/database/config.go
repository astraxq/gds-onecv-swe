package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
)

var Db *pgx.Conn //created outside to make it global.

//make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase(ctx context.Context) (*pgx.Conn, error) {

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

   conn, err := pgx.Connect(ctx, connOpts)

   if err != nil {
	 Db = conn
   }
   
   return conn, err
}