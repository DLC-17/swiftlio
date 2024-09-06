package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/swiftlio/internal/routes"
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

	var (
		dbUrl     = os.Getenv("DATABASE_URL")
		localPort = os.Getenv("LOCALHOST_PORT")
	)

	// DB Setup
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// dbQueries := database.New(db)

	// Initializing echo
	app := echo.New()
	routes.Init(app)

	app.Start("0.0.0.0" + localPort)
}
