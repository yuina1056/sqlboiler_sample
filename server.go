package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"test/domain/models"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/boil"
)

var DB *sql.DB

func Init() {
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=" + os.Getenv("POSTGRES_SSLMODE")

	DB, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		fmt.Println("Could not connect to database")
		panic(err)
	}

	// connection pool settings
	// DB.SetMaxIdleConns(10)
	// DB.SetMaxOpenConns(10)
	// DB.SetConnMaxLifetime(300 * time.Second)

	// global connection setting
	boil.SetDB(DB)
	boil.DebugMode = true
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Init()

	insertPilot := models.Pilot{
		ID:   1,
		Name: "hoge",
	}
	err = insertPilot.Insert(context.Background(), DB, boil.Infer())
	if err != nil {
		log.Fatal("Error")
	}

}
