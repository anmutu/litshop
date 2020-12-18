package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"litshop/src/config"
	"litshop/src/core/runtime"
	"log"
	"sync"
	"time"
)

var connClientMappingPool map[string]sync.Pool

type mysqlConnectParams struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func init() {
	conf := config.GetMap("database.mysql")
	for k, v := range conf {
		jByte, err := json.Marshal(v)
		if err != nil {
			log.Fatal(fmt.Sprintf("parse mysql config err: %#v", err))
			return
		}

		param := mysqlConnectParams{}
		err = json.Unmarshal(jByte, &param)
		if err != nil {
			log.Fatal(fmt.Sprintf("parse mysql config err: %#v", err))
			return
		}

		connClientMappingPool = make(map[string]sync.Pool, len(conf))
		connClientMappingPool[k] = sync.Pool{
			New: func() interface{} {
				return connect(buildDsn(param))
			},
		}
	}
}

func buildDsn(p mysqlConnectParams) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", p.User, p.Password, p.Host, p.Port, p.Name)
}

func ConnectByName(name string) *sql.DB {
	poll, ok := connClientMappingPool[name]
	if !ok {
		panic(errors.New("connect not defined"))
	}

	db := poll.Get()
	return db.(*sql.DB)
}

func connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	runtime.RegisterShutdown(func() {
		_ = db.Close()
	})

	return db
}
