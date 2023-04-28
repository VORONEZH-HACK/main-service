package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Error(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			log.Print(err.Error())
		}
		return err
	}
}
