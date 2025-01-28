package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sullyh7/portfolio/view/home"
)

func (a *application) homeHandler(c echo.Context) error {
	return a.render(c, http.StatusOK, home.Index())
}
