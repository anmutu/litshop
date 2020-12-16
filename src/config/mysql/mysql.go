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
	"time"
)

var clients map[string]*sql.DB

type mysqlConnectParams struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func init() {
	clients = make(map[string]*sql.DB, 1)

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

		err = connect(k, buildDsn(param))
		if err != nil {
			log.Fatal(fmt.Sprintf("parse mysql config err: %#v", err))
			return
		}
	}
}

func buildDsn(p mysqlConnectParams) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", p.User, p.Password, p.Host, p.Port, p.Name)
}

func ConnectByName(name string) *sql.DB {
	db, ok := clients[name]
	if !ok {
		panic(errors.New("connect not defined"))
	}

	return db
}

func connect(connectName string, dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	clients[connectName] = db

	runtime.RegisterShutdown(func() {
		_ = db.Close()
	})

	return nil
}
