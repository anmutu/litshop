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
	connClientMapping = make(map[string]*sql.DB)
	connections = make([]string, 0)

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

	fmt.Printf("connClientMappings %#v \n", connClientMapping)
}

func Connections() []string {
	return connections
}

func buildDsn(p mysqlConnectParams) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", p.User, p.Password, p.Host, p.Port, p.Name)
	//fmt.Printf("mysql dsn %s \n", dsn)
	return dsn
}

func ClientByConn(conn string) *sql.DB {
	return connClientMapping[conn]
}

func GormClientByConn(conn string) *gorm.DB {
	if conn == "" {
		return nil
	}

	sqlDB := connClientMapping[conn]
	fmt.Printf("sqlDB conn %#v %#v \n", sqlDB, conn)

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
