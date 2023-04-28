package endpoint

import (
	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/labstack/echo/v4"
)

func InvitePost(c echo.Context) error {
	userId := c.Request().Header.Get("User-Id")

	invitedUserId := c.QueryParam("userid")
	teamId := c.QueryParam("team")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}

	row := conn.QueryRow(db.PostgresQLDB.Get("select_user_x_team"), userId, teamId)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	err = row.Scan(&invitedUserId)
	if err != nil {
		return echo.NewHTTPError(400)
	}

	_, err = conn.Exec(db.PostgresQLDB.Get("insert_invite"), invitedUserId, teamId)
	if err != nil {
		return echo.NewHTTPError(400)
	}

	return c.String(200, "")
}
