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

var API_PREFIX string = "/api/v1"

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
		"accept_invite", "select_events_date",
		"insert_event", "select_events_newest",
		"insert_invite", "select_events_popular",
		"insert_team", "select_invites",
		"select_event", "select_team",
		"select_event_teams", "select_team_users",
		"select_events_best", "select_user", "select_user_x_team",
	}
	for i := 0; i < len(db_requests); i++ {
		err := db.PostgresQLDB.Prepare("internal/sql/", db_requests[i])
		if err != nil {
			log.Print(err.Error())
		}
	}

	e := echo.New()
	e.Use(middleware.Error)
	e.Use(middleware.Logger)
	e.GET(API_PREFIX+"/user/:uuid", endpoint.UserUuidGet)
	e.GET(API_PREFIX+"/team/:uuid", endpoint.TeamUuidGet)
	e.GET(API_PREFIX+"/events/:uuid", endpoint.EventsUuidGet)
	e.POST(API_PREFIX+"/events", endpoint.EventsPost, middleware.Auth)
	e.GET(API_PREFIX+"/events", endpoint.EventsGet)
	e.POST(API_PREFIX+"/team", endpoint.TeamPost)
	e.GET(API_PREFIX+"/invite", endpoint.InviteGet)
	e.POST(API_PREFIX+"/invite", endpoint.InvitePost)
	e.PATCH(API_PREFIX+"/invite", endpoint.InvitePatch)
	log.Print(e.Start(":10001").Error())
}

func main() {
	Start()
}
