package routes

import (
	"github.com/kelbwah/swiftlio/handlers"
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	app.Static("/public", "public")
	app.GET("/", handlers.HandleHomeShow)
	app.GET("/dashboard", handlers.HandleDashboardShow)
	app.GET("/login", handlers.HandleLoginShow)
}
