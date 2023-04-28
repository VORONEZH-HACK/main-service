package endpoint

import (
	"database/sql"
	"encoding/json"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/labstack/echo/v4"
)

func InviteGet(c echo.Context) error {
	userId := c.Request().Header.Get("User-Id")

	type ResponseBody struct {
		Invites []struct {
			Uuid     string `json:"uuid"`
			TeamUuid string `json:"team-uuid"`
		} `json:"invites"`
	}

	var resBody ResponseBody

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	rows, err := conn.Query(db.PostgresQLDB.Get("select_invites"), userId)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	for rows.Next() {
		var invite struct {
			Uuid     string `json:"uuid"`
			TeamUuid string `json:"team-uuid"`
		}
		err := rows.Scan(&invite.Uuid, &invite.TeamUuid)
		if err != nil {
			return echo.NewHTTPError(500)
		}
		resBody.Invites = append(resBody.Invites, invite)
	}

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
