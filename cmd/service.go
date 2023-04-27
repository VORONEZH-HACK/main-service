package main

import (
	"log"
	"math/rand"
	"time"

	db "github.com/VORONEZH-HACK/main-service/internal/database"
	"github.com/VORONEZH-HACK/main-service/internal/endpoint"
	"github.com/VORONEZH-HACK/main-service/internal/middleware"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var API_PRFIX string = "/api/v1"

func Start() {
	rand.Seed(time.Now().UnixNano())

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// err := db.PostgresQLDB.Init(db.DatabaseOptions{
	// 	Host:     os.Getenv("POSTGRES_HOST"),
	// 	Port:     os.Getenv("POSTGRES_PORT"),
	// 	User:     os.Getenv("POSTGRES_USER"),
	// 	Password: os.Getenv("POSTGRES_PASSWORD"),
	// 	Database: os.Getenv("POSTGRES_DB"),
	// })

	err := db.PostgresQLDB.Init(db.DatabaseOptions{
		Host:     "c-c9qt1ll800dqgld2lds5.rw.mdb.yandexcloud.net",
		Port:     "6432",
		User:     "bolat",
		Password: "Voro123456",
		Database: "voronezhack",
	})

	if err != nil {
		log.Print(err.Error())
		return
	}

	// err = db.RedisDB.Init(db.DatabaseOptions{
	// 	Host:     os.Getenv("REDIS_HOST"),
	// 	Port:     os.Getenv("REDIS_PORT"),
	// 	Password: os.Getenv("REDIS_PASSWORD"),
	// })

	// if err != nil {
	// 	log.Print(err.Error())
	// 	return
	// }

	var db_requests = []string{
		"accept_invite", "select_event", "select_invites", "insert_event",
		"select_events_date", "select_team", "insert_invite", "select_events_participants",
		"select_user", "insert_team", "select_events_rating",
	}
	for i := 0; i < len(db_requests); i++ {
		db.PostgresQLDB.Prepare("internal/sql/", db_requests[i])
	}

	e := echo.New()
	e.Use(middleware.Error)
	e.Use(middleware.Logger)
	e.GET("/user/:uuid", endpoint.UserUuidGet)
	log.Print(e.Start(":10000").Error())
}

func main() {
	Start()
}
