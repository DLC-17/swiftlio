package routes

import (
	// "github.com/kelbwah/swiftlio/internal/handlers/api"
	"github.com/kelbwah/swiftlio/internal/handlers/views"
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	// HTML View Endpoints
	app.Static("/public", "ui/public")
	app.GET("/", views.HandleHomeShow)
	app.GET("/dashboard", views.HandleDashboardShow)
	app.GET("/login", views.HandleLoginShow)

	// // API Endpoints
	// app.POST("/v1/login", api.HandleLogin)
	// app.PUT("/v1/user/:userId", api.HandleGetUser)
	// app.GET("/v1/user/:userId", api.HandleGetUser)
}
