package controllers

import (
	"github.com/labstack/echo"
)

func GetExample(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"key": "my value test",
	})
}
