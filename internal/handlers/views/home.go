package views

import (
	"github.com/kelbwah/swiftlio/ui/pages/home"
	u "github.com/kelbwah/swiftlio/utils"
	"github.com/labstack/echo/v4"
)

func HandleHomeShow(c echo.Context) error {
	return u.Render(c, home.Show())
}
