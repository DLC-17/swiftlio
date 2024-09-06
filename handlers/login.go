package handlers

import (
	"github.com/kelbwah/swiftlio/ui/pages/login"
	u "github.com/kelbwah/swiftlio/utils"
	"github.com/labstack/echo/v4"
)

func HandleLoginShow(c echo.Context) error {
	return u.Render(c, login.Show())
}
