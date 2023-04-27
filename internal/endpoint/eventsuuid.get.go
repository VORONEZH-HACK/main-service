package endpoint

import (
	"database/sql"
	"encoding/json"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/VORONEZH-HACK/main-service/internal/models"
	"github.com/labstack/echo/v4"
)

func EventsUuidGet(c echo.Context) error {
	type ResponseBody struct {
		models.Event
		Teams []struct {
			Uuid string `json:"uuid"`
			Name string `json:"name"`
		} `json:"teams"`
	}

	var resBody ResponseBody

	uuid := c.Param("uuid")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	row := conn.QueryRow(db.PostgresQLDB.Get("select_event"), uuid)
	err = row.Scan(
		&resBody.Uuid,
		&resBody.Name,
		&resBody.Owner,
		&resBody.Place,
		&resBody.Description,
		&resBody.StartDate,
		&resBody.EndDate,
		&resBody.MinParticipants,
		&resBody.MaxParticipants,
		&resBody.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(404, "not found")
		} else {
			return err
		}
	}

	rows, err := conn.Query(db.PostgresQLDB.Get("select_event_teams"), uuid)
	defer conn.Close()
	if err != nil {
		return echo.NewHTTPError(500)
	}
	for rows.Next() {
		var team struct {
			Uuid string `json:"uuid"`
			Name string `json:"name"`
		}
		err := rows.Scan(&team.Uuid, &team.Name)
		if err != nil {
			return echo.NewHTTPError(500)
		}
		resBody.Teams = append(resBody.Teams, team)
	}
	if err != nil {
		return echo.NewHTTPError(500)
	}

	resBytes, err := json.Marshal(resBody)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write(resBytes)

	return c.String(200, "")
}
