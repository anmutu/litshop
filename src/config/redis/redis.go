package redis

import (
	"fmt"
	rv8 "github.com/go-redis/redis/v8"
	"litshop/src/config"
)

var rdb *rv8.Client

func init() {
	dsn := config.GetString("database.redis.default.host") + ":" + config.GetString("database.redis.default.port")
	dbNo := config.GetInt("database.redis.default.database")
	if dbNo < 0 || dbNo > 16 {
		dbNo = 0
	}

	fmt.Printf("redis dsn %s\n", dsn)
	rdb = rv8.NewClient(&rv8.Options{
		Addr: dsn,
		DB:   dbNo,
	})
}

func Get() *rv8.Client {
	return rdb
}
