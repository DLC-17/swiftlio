package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelbwah/swiftlio/internal/routes"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

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

	echoErr := app.Start(localPort)
	if echoErr != nil {
		fmt.Printf("Error while starting echo: %s\n", echoErr.Error())
		return
	} else {
		fmt.Println("Echo started successfully")
	}
}
