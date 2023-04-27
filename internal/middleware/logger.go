package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Print("Logger")
		log.Printf("%s: %s", c.Request().Method, c.Request().URL.Path)
		return next(c)
	}
}
