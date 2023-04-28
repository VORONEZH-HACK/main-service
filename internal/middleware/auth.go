package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

const AUTH_SERVICE_HOST = "158.160.18.237"
const AUTH_SERVICE_PORT = 10000

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Access-Token")

		requestURL := fmt.Sprintf("http://%s:%d/api/v1/token", AUTH_SERVICE_HOST, AUTH_SERVICE_PORT)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Access-Token", token)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		if res.StatusCode == http.StatusOK {
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return err
			}
			type JWT struct {
				Code     int    `json:"code"`
				UserId   string `json:"user"`
				UserRole string `json:"type"`
			}
			var body JWT
			err = json.Unmarshal(resBody, &body)
			if body.Code != 200 {
				return echo.NewHTTPError(body.Code)
			}
			if err != nil {
				return echo.NewHTTPError(400)
			}

			c.Request().Header.Set("User-Id", body.UserId)
			c.Request().Header.Set("User-Role", body.UserRole)
			c.Request().Header.Set("Access-Token", string(resBody))
		} else {
			return echo.NewHTTPError(401)
		}

		return next(c)
	}
}
