package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/swiftlio/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func main() {
	// Loading env vars
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatalf("Error loading .env file %v\n", envErr)
	}

	// DB Setup
	var (
		localPort  = os.Getenv("LOCALHOST_PORT")
		pgHost     = os.Getenv("PGHOST")
		pgPort     = os.Getenv("PGPORT")
		pgUser     = os.Getenv("PGUSER")
		pgPassword = os.Getenv("PGPASSWORD")
		pgName     = os.Getenv("PGNAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Initializing echo
	app := echo.New()
	routes.Init(app)

	app.Start("0.0.0.0" + localPort)
}
