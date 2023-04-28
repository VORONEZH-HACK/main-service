package endpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/VORONEZH-HACK/main-service/internal/models"
	"github.com/labstack/echo/v4"
)

func EventsPost(c echo.Context) error {
	userId := c.Request().Header.Get("User-Id")
	// userRole := c.Request().Header.Get("User-Role")
	// if userRole != "organisation" {
	// 	return echo.NewHTTPError(401)
	// }

	reqBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(400)
	}
	var reqBody models.Event
	err = json.Unmarshal(reqBytes, &reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(400)
	}

	type ResponseBody struct {
		Uuid string `json:"uuid"`
	}
	var resBody ResponseBody

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	row := conn.QueryRow(db.PostgresQLDB.Get("insert_event"),
		reqBody.Name,
		reqBody.Description,
		userId,
		reqBody.MinParticipants,
		reqBody.MaxParticipants,
		reqBody.StartDate,
		reqBody.EndDate,
		reqBody.Place,
	)
	err = row.Scan(&resBody.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(500)
		} else {
			return err
		}
	}

	resBytes, err := json.Marshal(resBody)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write(resBytes)

	return c.String(200, "")
}
