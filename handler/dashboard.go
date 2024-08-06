package handler

import (
	"github.com/kelbwah/swiftlio/view/dashboard"
	"github.com/labstack/echo/v4"
)

func HandleDashboardShow(c echo.Context) error {
   return render(c, dashboard.Show()) 
}
