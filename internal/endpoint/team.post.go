package endpoint

import (
	"database/sql"
	"encoding/json"
	"io"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/VORONEZH-HACK/main-service/internal/models"
	"github.com/labstack/echo/v4"
)

func TeamPost(c echo.Context) error {
	reqBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	var reqBody models.Team
	err = json.Unmarshal(reqBytes, &reqBody)
	if err != nil {
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
	row := conn.QueryRow(db.PostgresQLDB.Get("insert_team"),
		reqBody.Name,
		reqBody.Lead,
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
