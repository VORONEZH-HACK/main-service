package endpoint

import (
	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/labstack/echo/v4"
)

func InvitePatch(c echo.Context) error {
	userId := c.Request().Header.Get("User-Id")

	inviteUuid := c.QueryParam("invite-uuid")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	_, err = conn.Exec(db.PostgresQLDB.Get("accept_invite"), inviteUuid, userId)
	if err != nil {
		return echo.NewHTTPError(400)
	}

	return c.String(200, "")
}
