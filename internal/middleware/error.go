package middleware

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Error(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		err := next(c)
		if err != echo.NewHTTPError {
			log.Print(err.Error())
			c.Response().WriteHeader(http.StatusInternalServerError)
			return echo.NewHTTPError(500)
		}
		return err
	}
}
