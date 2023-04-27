package endpoint

import (
	"database/sql"
	"encoding/json"
	"strconv"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/VORONEZH-HACK/main-service/internal/models"
	"github.com/labstack/echo/v4"
)

func getEvents(conn *sql.DB, events []models.EventMin, rType string, startDate string, endDate string) ([]models.EventMin, error) {
	if startDate == "" || endDate == "" {
		return events, echo.NewHTTPError(400)
	}

	startDateUnix, err := strconv.Atoi(startDate)
	if err != nil {
		return events, echo.NewHTTPError(400)
	}
	endDateUnix, err := strconv.Atoi(endDate)
	if err != nil {
		return events, echo.NewHTTPError(400)
	}

	rows, err := conn.Query(db.PostgresQLDB.Get(rType), startDateUnix, endDateUnix)
	if err != nil {
		return events, echo.NewHTTPError(400)
	}
	for rows.Next() {
		var event models.EventMin
		err := rows.Scan(
			&event.Uuid,
			&event.Name,
			&event.Owner,
			&event.Place,
			&event.Description,
			&event.StartDate,
			&event.EndDate,
			&event.MinParticipants,
			&event.MaxParticipants,
			&event.Rating,
		)
		if err != nil {
			return events, echo.NewHTTPError(500)
		}
		events = append(events, event)
	}

	return events, nil
}

func EventGet(c echo.Context) error {
	type ResponseBody struct {
		Events []models.EventMin `json:"events"`
	}

	var resBody ResponseBody

	filterType := c.Param("filter-type")
	startDate := c.Param("start-date")
	endDate := c.Param("end-date")

	conn, err := db.PostgresQLDB.Open()
	if err != nil {
		return err
	}
	defer conn.Close()

	if filterType == "" {
		resBody.Events, err = getEvents(conn, resBody.Events, "select_event_date", startDate, endDate)
		if err != nil {
			return err
		}
	} else if filterType == "best" {
		resBody.Events, err = getEvents(conn, resBody.Events, "select_event_best", startDate, endDate)
		if err != nil {
			return err
		}
	} else if filterType == "newest" {
		rows, err := conn.Query(db.PostgresQLDB.Get("select_event_newest"))
		if err != nil {
			return echo.NewHTTPError(400)
		}
		for rows.Next() {
			var event models.EventMin
			err := rows.Scan(
				&event.Uuid,
				&event.Name,
				&event.Owner,
				&event.Place,
				&event.Description,
				&event.StartDate,
				&event.EndDate,
				&event.MinParticipants,
				&event.MaxParticipants,
				&event.Rating,
			)
			if err != nil {
				return echo.NewHTTPError(500)
			}
			resBody.Events = append(resBody.Events, event)
		}
	} else if filterType == "popular" {
		resBody.Events, err = getEvents(conn, resBody.Events, "select_event_popular", startDate, endDate)
		if err != nil {
			return err
		}
	} else {
		return echo.NewHTTPError(400)
	}

	resBytes, err := json.Marshal(resBody)
	if err != nil {
		return err
	}

	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Write(resBytes)

	return c.String(200, "")
}