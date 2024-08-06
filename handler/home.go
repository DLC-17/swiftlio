package handler

import (
	"github.com/kelbwah/swiftlio/view/home"
	"github.com/labstack/echo/v4"
)

func HandleHomeShow(c echo.Context) error {
   return render(c, home.Show()) 
}
