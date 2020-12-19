package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"litshop/src/config"
	"litshop/src/lvm/runtime"
	"log"
	"time"
)

var connections []string
var connClientMapping map[string]*sql.DB

type mysqlConnectParams struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func init() {
	conf := config.GetMap("database.mysql")
	connClientMapping = make(map[string]*sql.DB, len(conf))
	connections = make([]string, len(conf))

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
		connections = append(connections, k)
		connClientMapping[k] = connect(buildDsn(param))
	}
}

func Connections() []string {
	return connections
}

func buildDsn(p mysqlConnectParams) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", p.User, p.Password, p.Host, p.Port, p.Name)
}

func ClientByConn(conn string) *sql.DB {
	return connClientMapping[conn]
}

func GormClientByConn(conn string) *gorm.DB {
	sqlDB := connClientMapping[conn]
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return gormDB
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
