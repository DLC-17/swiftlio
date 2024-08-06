package handler

import (
	"github.com/kelbwah/swiftlio/view/login"
	"github.com/labstack/echo/v4"
)

func HandleLoginShow(c echo.Context) error {
   return render(c, login.Show()) 
}
