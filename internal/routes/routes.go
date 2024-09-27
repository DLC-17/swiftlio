package routes

import (
	"github.com/kelbwah/swiftlio/internal/handlers/api"
	"github.com/kelbwah/swiftlio/internal/handlers/views"
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	ViewInit(app)
	ApiInit(app)
}

func ViewInit(app *echo.Echo) {
	app.Static("/public", "ui/public")
	app.GET("/", views.HandleHomeShow)
	app.GET("/dashboard", views.HandleDashboardShow)
	app.GET("/login", views.HandleLoginShow)
}

func ApiInit(app *echo.Echo) {
	app.POST("/v1/login", api.HandleLogin)
	app.PUT("/v1/user/:userId", api.HandleGetUser)
	app.GET("/v1/user/:userId", api.HandleGetUser)
}
