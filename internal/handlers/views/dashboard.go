package views

import (
	"github.com/kelbwah/swiftlio/ui/pages/dashboard"
	u "github.com/kelbwah/swiftlio/utils"
	"github.com/labstack/echo/v4"
)

func HandleDashboardShow(c echo.Context) error {
	return u.Render(c, dashboard.Show())
}
