package handler

import (
	"github.com/kelbwah/swiftlio/model"
	"github.com/kelbwah/swiftlio/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {}

func (h UserHandler) HandleUserShow(c echo.Context) error {
    testUser := model.User{
        Email: "email.com",
    }
    return render(c, user.Show(testUser))
}
