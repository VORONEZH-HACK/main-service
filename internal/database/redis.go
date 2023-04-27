package database

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	Options *redis.Options
}

func (db *Redis) Open() *redis.Client {
	return redis.NewClient(db.Options)
}

var RedisDB Redis

func (db *Redis) Init(options DatabaseOptions) error {
	db.Options = &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", options.Host, options.Port),
		Password: options.Password,
		DB:       0,
	}

	return nil
}
