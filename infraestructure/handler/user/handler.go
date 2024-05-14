package user

import (
	"e-commercego/domain/user"
	"e-commercego/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct {
	useCase user.UseCase
}

func newHandler(uc user.UseCase) handler {
	return handler{useCase: uc}

}

func (h handler) Create(c echo.Context) error {
	m := model.User{}

	if err := c.Bind(&m); err != nil {

		if err := h.useCase.Create(&m); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})

		}

		return c.JSON(http.StatusCreated, m)

	}


