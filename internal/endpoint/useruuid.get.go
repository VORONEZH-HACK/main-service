package endpoint

import (
	"database/sql"
	"encoding/json"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/labstack/echo/v4"
)

func UserUuidGet(c echo.Context) error {
	type ResponseBody struct {
		Uuid       string `json:"uuid"`
		Email      string `json:"-"`
		Password   string `json:"-"`
		Education  string `json:"education"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Surname    string `json:"surname"`
		Rating     string `json:"rating"`
	}

	var body ResponseBody

	uuid := c.Param("uuid")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	defer conn.Close()
	row := conn.QueryRow(db.PostgresQLDB.Get("select_user"), uuid)
	err = row.Scan(
		&body.Uuid,
		&body.Email,
		&body.Password,
		&body.Education,
		&body.Name,
		&body.Patronymic,
		&body.Surname,
		&body.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(404, "not found")
		} else {
			return err
		}
	}
	resBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write(resBytes)

	return c.String(200, "")
}
