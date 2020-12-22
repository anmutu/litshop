package redis

import (
	"fmt"
	rv8 "github.com/go-redis/redis/v8"
	"litshop/src/config"
)

var rdb *rv8.Client

func init() {
	dsn := config.GetString("database.redis.default.host") + ":" + config.GetString("database.redis.default.port")
	fmt.Printf("redis dsn %s\n", dsn)

	rdb = rv8.NewClient(&rv8.Options{
		Addr: dsn,
		DB:   0,
	})
}

func Get() *rv8.Client {
	return rdb
}
