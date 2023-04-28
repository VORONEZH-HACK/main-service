package endpoint

import (
	"database/sql"
	"encoding/json"
	"log"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/labstack/echo/v4"
)

func TeamUuidGet(c echo.Context) error {
	type ResponseBody struct {
		Uuid  string `json:"uuid"`
		Name  string `json:"name"`
		Lead  string `json:"lead"`
		Users []struct {
			Uuid       string `json:"uuid"`
			Name       string `json:"name"`
			Patronomic string `json:"patronomic"`
			Surname    string `json:"surname"`
		} `json:"users"`
	}

	var resBody ResponseBody

	uuid := c.Param("uuid")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	defer conn.Close()
	row := conn.QueryRow(db.PostgresQLDB.Get("select_team"), uuid)
	err = row.Scan(&resBody.Uuid, &resBody.Name, &resBody.Lead)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(404, "not found")
		} else {
			return err
		}
	}

	rows, err := conn.Query(db.PostgresQLDB.Get("select_team_users"), uuid)
	if err != nil {
		log.Print(err.Error())
		return echo.NewHTTPError(500)
	}
	for rows.Next() {
		var user struct {
			Uuid       string `json:"uuid"`
			Name       string `json:"name"`
			Patronomic string `json:"patronomic"`
			Surname    string `json:"surname"`
		}
		err := rows.Scan(&user.Uuid, &user.Name, &user.Patronomic, &user.Surname)
		if err == sql.ErrNoRows {
			break
		} else {
			if err != sql.ErrNoRows && err != nil {
				log.Print(err.Error())
				return echo.NewHTTPError(500)
			}
		}
		resBody.Users = append(resBody.Users, user)
	}

	resBytes, err := json.Marshal(resBody)
	if err != nil {
		log.Print(err.Error())
		return echo.NewHTTPError(500)
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write(resBytes)

	return c.String(200, "")
}
