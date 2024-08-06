package main

import (
	"log"
    "fmt"
	"os"
    "database/sql"

    _ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/kelbwah/swiftlio/handler"
	"github.com/labstack/echo/v4"
)

type DB struct {
    Host     string
    Port     string
    User     string
    Password string
    DbName   string
}

func main() {
    app := echo.New()
    if envErr := godotenv.Load(); envErr != nil {
        log.Fatal("Error loading .env file")
    }

    // DB Setup
    var (
        localPort = os.Getenv("LOCALHOST_PORT")
        pgHost = os.Getenv("PGHOST")
        pgPort = os.Getenv("PGPORT")
        pgUser = os.Getenv("PGUSER")
        pgPassword = os.Getenv("PGPASSWORD")
        pgName = os.Getenv("PGNAME")
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
    
    //Routes
    userHandler := handler.UserHandler{}
    app.Static("/public", "public")
    app.GET("/", handler.HandleHomeShow)
    app.GET("/dashboard", handler.HandleDashboardShow)
    app.GET("/login", handler.HandleLoginShow)
    app.GET("/user", userHandler.HandleUserShow)
    
    app.Start(localPort)
}
